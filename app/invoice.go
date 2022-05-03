/*
Copyright 2022.
This file uses pdfcpu as PDF processor (https://pdfcpu.io/).

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
	http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ledgertech/billing-srv/config"
	"github.com/ledgertech/billing-srv/util/assets"
	"github.com/ledgertech/billing-srv/util/cfile"
	"github.com/ledgertech/billing-srv/util/clog"
	"github.com/ledgertech/billing-srv/util/tpl"
	pdfapi "github.com/pdfcpu/pdfcpu/pkg/api"
)

type elementS struct {
	Text  string
	Value string
}
type elementI struct {
	Text  string
	Value int64
}
type elementF struct {
	Text  string
	Value float64
}
type modelConfig struct {
	FileRootName  string
	FilePath      string
	LogoPath      string
	LogoURL       string
	WatermarkPath string
	CustomColor   string
	CustomText    string
	CustomTextHE  string
	Lang          string
	VATPercent    float64
}
type modelParty struct {
	Head    elementS
	Name    elementS
	Address elementS
	City    elementS
	Country elementS
	Phone   elementS
	Email   elementS
}
type modelInvoice struct {
	Lang               string
	CurrencyISO        string
	CurrencySymbol     string
	CustomText         string
	ExitMessage        string
	Head               elementS
	Number             elementS
	Date               elementS
	DaysToDueDate      elementI
	InvoiceDetails     []modelInvoiceDetail
	InvoiceDetailsHead modelInvoiceDetailHead
	InvoiceSummary     modelInvoiceSummary
}
type modelInvoiceDetailHead struct {
	DescriptionText string
	QuantityText    string
	PriceText       string
	AmountText      string
}
type modelInvoiceDetail struct {
	DescriptionValue string
	QuantityValue    int64
	PriceValue       float64
	AmountValue      float64
}
type modelInvoiceSummary struct {
	DiscountPercent float64
	Discount        float64
	VATPercent      float64
	VAT             float64
	SubTotal        float64
	Total           float64
	SummaryRows     int64
	SummaryDiscount elementS
	SummaryVAT      elementS
	SummarySubTotal elementS
	SummaryTotal    elementS
}
type modelFullInvoice struct {
	ObjectConfig modelConfig
	Insurer      modelParty
	Customer     modelParty
	Invoice      modelInvoice
}

//func compute addition info (summary, labels ...) for the invoice
func (data *modelFullInvoice) prepareInvoiceForTemplate() error {

	var (
		tablerows   int     = 3
		summaryrows int64   = 1
		subtotal    float64 = 0
	)
	//compute summary
	for _, v2 := range data.Invoice.InvoiceDetails {
		subtotal = subtotal + v2.AmountValue
	}

	data.Invoice.InvoiceSummary.Discount = math.Round(data.Invoice.InvoiceSummary.Discount*100) / 100
	if data.Invoice.InvoiceSummary.Discount > 0 {
		summaryrows += 1
	}

	data.Invoice.InvoiceSummary.VAT = math.Round(data.Invoice.InvoiceSummary.VAT*100) / 100
	if data.Invoice.InvoiceSummary.VAT > 0 {
		summaryrows += 1
	}

	data.Invoice.InvoiceSummary.Total = math.Round((subtotal-data.Invoice.InvoiceSummary.Discount+data.Invoice.InvoiceSummary.VAT)*100) / 100

	if data.Invoice.InvoiceSummary.Discount > 0 ||
		data.Invoice.InvoiceSummary.VAT > 0 {
		data.Invoice.InvoiceSummary.SubTotal = math.Round(subtotal*100) / 100
		summaryrows += 1
	}

	data.Invoice.InvoiceSummary.SummaryRows = summaryrows

	//texts - labels
	data.Invoice.CurrencySymbol = getCurrencySymbol(data.Invoice.CurrencyISO)

	lang := data.Invoice.Lang
	if strings.TrimSpace(lang) == "" {
		lang = data.ObjectConfig.Lang
	}
	switch lang {
	case "he-IL":
		{
			//issuer
			data.Insurer.Head.Text = reverseHE("מְבַטֵחַ")
			data.Insurer.Name.Text = reverseHE("שֵׁם:")
			data.Insurer.Name.Value = reverseHE(data.Insurer.Name.Value)
			data.Insurer.Address.Text = reverseHE("כתובת:")
			data.Insurer.Address.Value = reverseHE(data.Insurer.Address.Value)
			data.Insurer.City.Text = reverseHE("עִיר:")
			data.Insurer.City.Value = reverseHE(data.Insurer.City.Value)
			data.Insurer.Country.Text = reverseHE("מדינה:")
			data.Insurer.Country.Value = reverseHE(data.Insurer.Country.Value)
			data.Insurer.Phone.Text = reverseHE("טלפון:")
			data.Insurer.Email.Text = reverseHE("אֶלֶקטרוֹנִי:")

			//customer
			data.Customer.Head.Text = reverseHE("צרכן")
			data.Customer.Name.Text = reverseHE("שֵׁם:")
			data.Customer.Name.Value = reverseHE(data.Customer.Name.Value)
			data.Customer.Address.Text = reverseHE("כתובת:")
			data.Customer.Address.Value = reverseHE(data.Customer.Address.Value)
			data.Customer.City.Text = reverseHE("עִיר:")
			data.Customer.City.Value = reverseHE(data.Customer.City.Value)
			data.Customer.Country.Text = reverseHE("מדינה:")
			data.Customer.Country.Value = reverseHE(data.Customer.Country.Value)
			data.Customer.Phone.Text = reverseHE("טלפון:")
			data.Customer.Email.Text = reverseHE("אֶלֶקטרוֹנִי:")

			//invoice
			strhe := reverseHE(data.ObjectConfig.CustomTextHE)
			switch l := len(strhe); {
			case l > 0:
				data.Invoice.CustomText = splitStringToMultiLinesRTL_v1(strhe, 2, 15)
			case l <= 0:
				data.Invoice.CustomText = " "
			}

			data.Invoice.Head.Text = reverseHE("חשבונית")
			data.Invoice.Number.Text = reverseHE("מִספָּר:")
			data.Invoice.Date.Text = reverseHE("תַאֲרִיך:")

			data.Invoice.ExitMessage = reverseHE("תודה לך על העסק שלך!")
			switch {
			case data.Invoice.DaysToDueDate.Value == 0:
				data.Invoice.DaysToDueDate.Text = reverseHE("התשלום יתבצע היום.")
			case data.Invoice.DaysToDueDate.Value > 0:
				data.Invoice.DaysToDueDate.Text = reverseHE(strings.ReplaceAll("התשלום יתבצע תוך # ימים.", "#", fmt.Sprintf("%d", data.Invoice.DaysToDueDate.Value)))
			case data.Invoice.DaysToDueDate.Value < 0:
				data.Invoice.DaysToDueDate.Text = " "
			}

			data.Invoice.InvoiceDetailsHead.QuantityText = reverseHE("כַּמוּת")
			data.Invoice.InvoiceDetailsHead.DescriptionText = reverseHE("תיאור")
			data.Invoice.InvoiceDetailsHead.PriceText = fmt.Sprintf("(%s) %s", data.Invoice.CurrencySymbol, reverseHE("מחיר"))
			data.Invoice.InvoiceDetailsHead.AmountText = fmt.Sprintf("(%s) %s", data.Invoice.CurrencySymbol, reverseHE("סְכוּם"))

			if len(data.Invoice.InvoiceDetails) > tablerows { // template has a limited number of rows
				return fmt.Errorf("invoice body supports maximum %d rows, but the request asked for %d", tablerows, len(data.Invoice.InvoiceDetails))
			}
			for k, v := range data.Invoice.InvoiceDetails {
				strhev := reverseHE(v.DescriptionValue)
				data.Invoice.InvoiceDetails[k].DescriptionValue = splitStringToMultiLinesRTL_v1(strhev, 6, 50)
			}

			data.Invoice.InvoiceSummary.SummaryDiscount.Text = reverseHE("הנחה") //fmt.Sprintf("(%.2f%s)%s", discountPercentUsed, "%", reverseHE("הנחה"))
			data.Invoice.InvoiceSummary.SummaryDiscount.Value = fmt.Sprintf("(%s) %.2f", data.Invoice.CurrencySymbol, data.Invoice.InvoiceSummary.Discount)
			data.Invoice.InvoiceSummary.SummaryVAT.Text = reverseHE("מס ערך מוסף") //fmt.Sprintf("(%.2f%s)%s", vatPercentUsed, "%", reverseHE("מס ערך מוסף"))
			data.Invoice.InvoiceSummary.SummaryVAT.Value = fmt.Sprintf("(%s) %.2f", data.Invoice.CurrencySymbol, data.Invoice.InvoiceSummary.VAT)
			data.Invoice.InvoiceSummary.SummarySubTotal.Text = reverseHE("סכום ביניים")
			data.Invoice.InvoiceSummary.SummarySubTotal.Value = fmt.Sprintf("(%s) %.2f", data.Invoice.CurrencySymbol, data.Invoice.InvoiceSummary.SubTotal)
			data.Invoice.InvoiceSummary.SummaryTotal.Text = reverseHE("סך הכל")
			data.Invoice.InvoiceSummary.SummaryTotal.Value = fmt.Sprintf("(%s) %.2f", data.Invoice.CurrencySymbol, data.Invoice.InvoiceSummary.Total)
		}
	default:
		{
			//issuer
			data.Insurer.Head.Text = "Insurer"
			data.Insurer.Name.Text = "Name:"
			data.Insurer.Address.Text = "Address:"
			data.Insurer.City.Text = "City:"
			data.Insurer.Country.Text = "Country:"
			data.Insurer.Phone.Text = "Phone:"
			data.Insurer.Email.Text = "Email:"

			//customer
			data.Customer.Head.Text = "Customer"
			data.Customer.Name.Text = "Name:"
			data.Customer.Address.Text = "Address:"
			data.Customer.City.Text = "City:"
			data.Customer.Country.Text = "Country:"
			data.Customer.Phone.Text = "Phone:"
			data.Customer.Email.Text = "Email:"

			//invoice
			switch l := len(data.ObjectConfig.CustomText); {
			case l > 0:
				data.Invoice.CustomText = splitStringToMultiLinesLTR_v1(data.ObjectConfig.CustomText, 2, 15)
			case l <= 0:
				data.Invoice.CustomText = " "
			}

			data.Invoice.Head.Text = "Invoice"
			data.Invoice.Number.Text = "Number:"
			data.Invoice.Date.Text = "Date:"
			data.Invoice.ExitMessage = "Thank you for your business!"

			switch {
			case data.Invoice.DaysToDueDate.Value == 0:
				data.Invoice.DaysToDueDate.Text = "Payment is due today."
			case data.Invoice.DaysToDueDate.Value > 0:
				data.Invoice.DaysToDueDate.Text = fmt.Sprintf("Payment is due within %d days.", data.Invoice.DaysToDueDate.Value)
			case data.Invoice.DaysToDueDate.Value < 0:
				data.Invoice.DaysToDueDate.Text = " "
			}

			data.Invoice.InvoiceDetailsHead.QuantityText = "Qty"
			data.Invoice.InvoiceDetailsHead.DescriptionText = "Description"
			data.Invoice.InvoiceDetailsHead.PriceText = fmt.Sprintf("%s (%s)", "Price", data.Invoice.CurrencySymbol)
			data.Invoice.InvoiceDetailsHead.AmountText = fmt.Sprintf("%s (%s)", "Amount", data.Invoice.CurrencySymbol)

			if len(data.Invoice.InvoiceDetails) > tablerows { // template has a limited number of rows
				return fmt.Errorf("invoice body supports maximum %d rows, but the request asked for %d", tablerows, len(data.Invoice.InvoiceDetails))
			}
			for k, v := range data.Invoice.InvoiceDetails {
				data.Invoice.InvoiceDetails[k].DescriptionValue = splitStringToMultiLinesLTR_v1(v.DescriptionValue, 6, 40)
			}

			data.Invoice.InvoiceSummary.SummaryDiscount.Text = "Discount" //fmt.Sprintf("%s (%.2f%s)", "Disc", discountPercentUsed, "%")
			data.Invoice.InvoiceSummary.SummaryDiscount.Value = fmt.Sprintf("%.2f (%s)", data.Invoice.InvoiceSummary.Discount, data.Invoice.CurrencySymbol)
			data.Invoice.InvoiceSummary.SummaryVAT.Text = "VAT" //fmt.Sprintf("%s (%.2f%s)", "VAT", vatPercentUsed, "%")
			data.Invoice.InvoiceSummary.SummaryVAT.Value = fmt.Sprintf("%.2f (%s)", data.Invoice.InvoiceSummary.VAT, data.Invoice.CurrencySymbol)
			data.Invoice.InvoiceSummary.SummarySubTotal.Text = "SubTotal"
			data.Invoice.InvoiceSummary.SummarySubTotal.Value = fmt.Sprintf("%.2f (%s)", data.Invoice.InvoiceSummary.SubTotal, data.Invoice.CurrencySymbol)
			data.Invoice.InvoiceSummary.SummaryTotal.Text = "Total"
			data.Invoice.InvoiceSummary.SummaryTotal.Value = fmt.Sprintf("%.2f (%s)", data.Invoice.InvoiceSummary.Total, data.Invoice.CurrencySymbol)
		}
	}

	return nil
}

//func create invoice PDF document
func (data *modelFullInvoice) createInvoicePDF() ([]byte, error) {
	var (
		err            error
		fileIdentifier string = fmt.Sprintf("%s-%s", data.ObjectConfig.FileRootName, data.Invoice.Number.Value)
		lang           string
		tmplName       string = "invoice-en-mono"
	)

	clog.Infof("start creating invoice:%s", fileIdentifier)

	lang = data.Invoice.Lang
	if strings.TrimSpace(lang) == "" {
		lang = data.ObjectConfig.Lang
	}
	if lang == "he-IL" {
		tmplName = "invoice-he-mono"
	}
	clog.Infof("template used:'%s'", tmplName)

	tmpl, err := tpl.LoadTemplate(tmplName)
	if err != nil {
		clog.Errorf("fail while loading template:'%s' error: %+v\n", tmplName, err)
		return nil, err
	}

	b := new(bytes.Buffer)
	err = tmpl.Execute(b, data)
	if err != nil {
		clog.Errorf("fail while refreshing template data error: %+v\n", err)
		return nil, err
	}
	//clog.Infof("Json Body:%s", b.String())

	filesName := []string{
		filepath.Join(data.ObjectConfig.FilePath, fmt.Sprintf("%s.%s", fileIdentifier, "json")),
		filepath.Join(data.ObjectConfig.FilePath, fmt.Sprintf("%s-%s.%s", fileIdentifier, "TEMPORARY", "pdf")),
		filepath.Join(data.ObjectConfig.FilePath, fmt.Sprintf("%s.%s", fileIdentifier, "pdf")),
	}
	//clog.Infof("Json FileName:%+v", filesName)

	//create json file
	err = cfile.WriteFile(filesName[0], b.Bytes())
	if err != nil {
		clog.Errorf("error while creating JSON file:%s %+v\n", filesName[0], err)
		return nil, err
	}

	//create intermediate pdf file (json -> pdf1)
	err = pdfapi.CreateFromJSONFile(filesName[0], "", filesName[1], nil)
	if err != nil {
		clog.Errorf("error while creating intermediate PDF file:%s %+v\n", filesName[1], err)
		return nil, err
	}

	err = pdfapi.AddImageWatermarksFile(filesName[1], filesName[2], []string{"1-"}, true, filepath.Join(data.ObjectConfig.WatermarkPath), "sc:0.5 abs, op:0.2, al:c, pos:c, rot:0, mo:0", nil)
	if err != nil {
		clog.Errorf("error while creating intermediate PDF file:%s %+v\n", filesName[1], err)
		return nil, err
	}

	streamPDFbytes, err := ioutil.ReadFile(filesName[2])
	if err != nil {
		clog.Errorf("error while reading final PDF file:%s %+v\n", filesName[2], err)
		return nil, err
	}

	for i := 0; i < len(filesName)-1; i++ {
		if i == 1 {
			err = cfile.RemoveFile(filesName[i])
			if err != nil {
				clog.Errorf("cannot remove file:%s %+v\n", filesName[i], err)
				return nil, err
			}
			clog.Infof("file removed:%s\n", filesName[i])
		}
	}

	return streamPDFbytes, nil
}

//func check values'integrity , values comming from configuration file
func checkSetDefaultValues(config config.AppConfig) (modelConfig, error) {
	var (
		err                        error
		defaultOutputDirectoryName string    = "out"
		defaultCustomColor         string    = "#85C2EB"
		defaultCustomText          string    = "INSURANCE INVOICE"
		defaultCustomTextHE        string    = "חשבונית ביטוח"
		defaultLang                string    = "en-EN"
		t                          time.Time = time.Now()
	)

	out := modelConfig{
		FileRootName: fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day()), //uuid.New().String(),
		LogoURL:      config.WorkStyle.LogoURL,
		VATPercent:   config.WorkStyle.VAT,
	}

	//output folder
	if cfile.CheckFolderExists(config.WorkStyle.OutputPath) {
		mydir, err := os.Getwd()
		if err != nil {
			return modelConfig{}, err
		}
		clog.Infof("PWD:%s\n", mydir)

		out.FilePath = config.WorkStyle.OutputPath
	} else {
		err = cfile.CreateFolder(defaultOutputDirectoryName)
		if err != nil {
			return modelConfig{}, err
		}
		out.FilePath = defaultOutputDirectoryName
		clog.Infof("use default output folder:%s\n", defaultOutputDirectoryName)
	}

	//color
	if len(strings.TrimSpace(config.WorkStyle.CustomColor)) > 0 {
		out.CustomColor = strings.ReplaceAll(strings.ReplaceAll(config.WorkStyle.CustomColor, "/#", "#"), "\\#", "#")
	} else {
		out.CustomColor = defaultCustomColor
		clog.Infof("use default custom color:%s\n", defaultCustomColor)
	}

	//lang
	{
		ln := strings.TrimSpace(config.WorkStyle.Lang)
		switch ln {
		case "he-IL":
			out.Lang = ln
		default:
			out.Lang = defaultLang
		}
	}

	//custom text
	switch out.Lang {
	case "he-IL":
		if len(strings.TrimSpace(config.WorkStyle.CustomText)) > 0 {
			out.CustomTextHE = strings.TrimSpace(config.WorkStyle.CustomText)
		} else {
			out.CustomTextHE = defaultCustomTextHE
		}
		out.CustomText = defaultCustomText
	default:
		if len(strings.TrimSpace(config.WorkStyle.CustomText)) > 0 {
			out.CustomText = strings.TrimSpace(config.WorkStyle.CustomText)
		} else {
			out.CustomText = defaultCustomText
		}
		out.CustomTextHE = defaultCustomTextHE
	}

	imgs, err := assets.LoadEmbeddedImages()
	if err != nil {
		return modelConfig{}, err
	}

	//logo image
	if cfile.CheckFileExists(config.WorkStyle.LogoPath) {
		out.LogoPath = config.WorkStyle.LogoPath
	} else {
		out.LogoPath = imgs["ledgertech_logo1"].FileLongPath
		out.LogoURL = ""
		clog.Infof("use default logo image:%s\n", imgs["ledgertech_logo1"].FileLongPath)
	}
	//watermark image
	if cfile.CheckFileExists(config.WorkStyle.WatermarkPath) {
		out.WatermarkPath = config.WorkStyle.WatermarkPath
	} else {
		out.WatermarkPath = imgs["ledgertech_icon1"].FileLongPath
		clog.Infof("use default watermark image:%s\n", imgs["ledgertech_icon1"].FileLongPath)
	}

	return out, nil
}

func generateInvoice(invoice *modelFullInvoice) ([]byte, error) {
	clog.Info("start generating invoice")

	var err error
	objconf, err := checkSetDefaultValues(config.App)
	if err != nil {
		clog.Errorf("fail to check/set default values - error:%v\n", err)
		return nil, err
	}
	invoice.ObjectConfig = objconf

	err = handleEmbededFonts()
	if err != nil {
		clog.Errorf("fail while handling embedded fonts - error: %+v\n", err)
		return nil, err
	}

	err = invoice.prepareInvoiceForTemplate()
	if err != nil {
		clog.Errorf("fail to prepare invoice for template - error: %+v\n", err)
		return nil, err
	}
	//clog.Infof("invoice after:%s", v)

	streamPDFbytes, err := invoice.createInvoicePDF()
	if err != nil {
		clog.Errorf("fail to create invoice ad PDF - error: %+v\n", err)
		return nil, err
	}

	return streamPDFbytes, nil
}
