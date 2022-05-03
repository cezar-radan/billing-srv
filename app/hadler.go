package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ledgertech/billing-srv/util/clog"
)

type inputParty struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type inputInvoice struct {
	Lang            string               `json:"langcode"`
	Number          string               `json:"number"`
	Date            int64                `json:"date"`
	DaysToDueDate   int64                `json:"daysToDueDate"`
	Currency        string               `json:"currency"`
	DiscountPercent float64              `json:"discountPercent"`
	Discount        float64              `json:"discount"`
	VATPercent      float64              `json:"vatPercent"`
	VAT             float64              `json:"vat"`
	InvoiceDetails  []inputInvoiceDetail `json:"details"`
}

type inputInvoiceDetail struct {
	Description string  `json:"description"`
	Quantity    int64   `json:"quantity"`
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount"`
}

type inputFullInvoice struct {
	Insurer  inputParty   `json:"insurer"`
	Customer inputParty   `json:"customer"`
	Invoice  inputInvoice `json:"invoice"`
}

func (data *inputFullInvoice) validateInputFullInvoice() (*modelFullInvoice, error) {

	if len(strings.TrimSpace(data.Insurer.Name)) == 0 {
		return nil, fmt.Errorf("bad input insurer name:'%s'", data.Insurer.Name)
	}
	if len(strings.TrimSpace(data.Customer.Name)) == 0 {
		return nil, fmt.Errorf("bad input customer name:'%s'", data.Customer.Name)
	}
	if len(strings.TrimSpace(data.Invoice.Number)) == 0 {
		return nil, fmt.Errorf("bad input invoice number:'%s'", data.Invoice.Number)
	}
	invoiceDate := ""
	if data.Invoice.Date > 0 {
		id := time.UnixMilli(data.Invoice.Date)
		invoiceDate = id.Format("2006-01-02")
	} else {
		return nil, fmt.Errorf("bad input invoice date:'%d'", data.Invoice.Date)
	}
	if len(data.Invoice.InvoiceDetails) == 0 {
		return nil, fmt.Errorf("bad input invoice details. number of items at the invoice level :'%d'", len(data.Invoice.InvoiceDetails))
	}

	out := modelFullInvoice{
		Insurer: modelParty{
			Name:    elementS{Value: data.Insurer.Name},
			Address: elementS{Value: data.Insurer.Address},
			City:    elementS{Value: data.Insurer.City},
			Country: elementS{Value: data.Insurer.Country},
			Phone:   elementS{Value: data.Insurer.Phone},
			Email:   elementS{Value: data.Insurer.Email},
		},
		Customer: modelParty{
			Name:    elementS{Value: data.Customer.Name},
			Address: elementS{Value: data.Customer.Address},
			City:    elementS{Value: data.Customer.City},
			Country: elementS{Value: data.Customer.Country},
			Phone:   elementS{Value: data.Customer.Phone},
			Email:   elementS{Value: data.Customer.Email},
		},
		Invoice: modelInvoice{
			Lang:          data.Invoice.Lang,
			Number:        elementS{Value: data.Invoice.Number},
			Date:          elementS{Value: invoiceDate},
			DaysToDueDate: elementI{Value: data.Invoice.DaysToDueDate},
			CurrencyISO:   data.Invoice.Currency,
			InvoiceSummary: modelInvoiceSummary{
				DiscountPercent: data.Invoice.DiscountPercent,
				Discount:        data.Invoice.Discount,
				VATPercent:      data.Invoice.VATPercent,
				VAT:             data.Invoice.VAT,
			},
			InvoiceDetails: []modelInvoiceDetail{},
		},
	}

	for _, v := range data.Invoice.InvoiceDetails {
		invoicedetail := modelInvoiceDetail{
			DescriptionValue: v.Description,
			QuantityValue:    v.Quantity,
			PriceValue:       v.Price,
			AmountValue:      v.Amount,
		}

		out.Invoice.InvoiceDetails = append(out.Invoice.InvoiceDetails, invoicedetail)
	}

	return &out, nil
}

func generateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("getTestEventsHandler: got following request: Method='%s'  Url='%v' \n", r.Method, r.URL)

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	defer r.Body.Close()

	bRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		outMsg := "error while reading the request"
		log.Printf("getTestEventsHandler: %s error: %+v \n", outMsg, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(outMsg))
		return
	}
	clog.Infof("getTestEventsHandler: got following request body:%s \n", string(bRequest))

	req := &inputFullInvoice{}
	err = json.Unmarshal(bRequest, req)
	if err != nil {
		outMsg := "error while reading the request body"
		clog.Errorf("getTestEventsHandler: %s error: %+v \n", outMsg, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(outMsg))
		return
	}

	mdlinvoice, err := req.validateInputFullInvoice()
	if err != nil {
		outMsg := "error while validating the request body"
		clog.Errorf("getTestEventsHandler: %s error: %+v \n", outMsg, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(outMsg))
		return
	}

	clog.Infof("getTestEventsHandler: extracted data:%+v \n", mdlinvoice)

	streamPDFbytes, err := generateInvoice(mdlinvoice)
	if err != nil {
		outMsg := "error while generating invoice"
		clog.Errorf("getTestEventsHandler: %s error: %+v \n", outMsg, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(outMsg))
		return
	}
	/*
		b := bytes.NewBuffer(streamPDFbytes)
		if _, err := b.WriteTo(w); err != nil {
			outMsg := "error while buffering pdf stream PDF bytes"
			clog.Errorf("getTestEventsHandler: %s error: %+v \n", outMsg, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(outMsg))
			return
		}
	*/

	w.Header().Set("Content-type", "application/pdf")
	w.WriteHeader(http.StatusOK)
	w.Write(streamPDFbytes)
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("NotImplemented - Got following request: Method='%s'  Url='%v'\n", r.Method, r.URL)

	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("NotFound - Got following request: Method='%s'  Url='%v'\n", r.Method, r.URL)

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(http.StatusText(http.StatusNotFound)))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	clog.Infof("Got following request: Method='%s'  Url='%v' \n", r.Method, r.URL)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(" I am alive ! "))
}
