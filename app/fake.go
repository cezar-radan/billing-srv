package app

import (
	"github.com/ledgertech/billing-srv/config"
	"github.com/ledgertech/billing-srv/util/clog"
)

func generateFakeData(objconfig modelConfig) []*modelFullInvoice {

	allData := []*modelFullInvoice{}

	data1 := modelFullInvoice{
		ObjectConfig: objconfig,
		Insurer: modelParty{
			Name:    elementS{Value: "Ledgertech Insurer"},
			Address: elementS{Value: "Dreikönigstrsse 47"},
			City:    elementS{Value: "?"},
			Country: elementS{Value: "Switzerland"},
			Phone:   elementS{Value: "+00000000000"},
			Email:   elementS{Value: "office@ledgertech.com"},
		},
		Customer: modelParty{
			Name:    elementS{Value: "Buyer1 Name"},
			Address: elementS{Value: "Buyer1 Street Name and Number"},
			City:    elementS{Value: "Buyer1 City"},
			Country: elementS{Value: "Buyer1 Country"},
			Phone:   elementS{Value: "+01111111111111"},
			Email:   elementS{Value: "buyer1@domain.com"},
		},
		Invoice: modelInvoice{
			Lang:          "en-EN",
			CurrencyISO:   "USD",
			Number:        elementS{Value: "LT-1000001"},
			Date:          elementS{Value: "2022-05-01"},
			DaysToDueDate: elementI{Value: -1},
			InvoiceDetails: []modelInvoiceDetail{
				{
					DescriptionValue: "Insurance for primary house1 and hollyyyy complementary insurance for garage1. Insurance for primary house2 and complementary insurance for garage2. Insurance for primary house3 and complementary insurance for garage3.",
					QuantityValue:    1,
					PriceValue:       1277.00,
					AmountValue:      1277.00,
				},
				{
					DescriptionValue: "Home Insurance for parent's house.",
					QuantityValue:    3,
					PriceValue:       210.00,
					AmountValue:      630.00,
				},
				{
					DescriptionValue: "Travel insurance for extended family - wife, children, pets.",
					QuantityValue:    1,
					PriceValue:       1000.00,
					AmountValue:      1000.00,
				},
			},
			InvoiceSummary: modelInvoiceSummary{
				DiscountPercent: 10.10,
				Discount:        10.11,
				VATPercent:      0.00,
				VAT:             0.00,
			},
		},
	}
	if objconfig.Lang == "en-EN" {
		allData = append(allData, &data1)
	}

	data2 := modelFullInvoice{
		ObjectConfig: objconfig,
		Insurer: modelParty{
			Name:    elementS{Value: "Ledgertech Insurer"},
			Address: elementS{Value: "Dreikönigstrsse 47"},
			City:    elementS{Value: "Zürich"},
			Country: elementS{Value: "Switzerland"},
			Phone:   elementS{Value: "+00000000000"},
			Email:   elementS{Value: "office@ledgertech.com"},
		},
		Customer: modelParty{
			Name:    elementS{Value: "Buyer2 Name"},
			Address: elementS{Value: "Buyer2 Street Name and Number"},
			City:    elementS{Value: "Buyer2 City"},
			Country: elementS{Value: "Buyer2 Country"},
			Phone:   elementS{Value: "+02222222222222"},
			Email:   elementS{Value: "buyer2@domain.com"},
		},
		Invoice: modelInvoice{
			Lang:          "en-EN",
			CurrencyISO:   "EUR",
			Number:        elementS{Value: "LT-1000002"},
			Date:          elementS{Value: "2022-05-02"},
			DaysToDueDate: elementI{Value: 11},
			InvoiceDetails: []modelInvoiceDetail{
				{
					DescriptionValue: "Home Insurance for primary house.",
					QuantityValue:    1,
					PriceValue:       610.00,
					AmountValue:      610.00,
				},
				{
					DescriptionValue: "Home Insurance for secondary house Home Insurance for third house.",
					QuantityValue:    3,
					PriceValue:       210.00,
					AmountValue:      630.00,
				},
			},
			InvoiceSummary: modelInvoiceSummary{
				DiscountPercent: 0,
				Discount:        0,
				VATPercent:      0.00,
				VAT:             0.00,
			},
		},
	}

	if objconfig.Lang == "en-EN" {
		allData = append(allData, &data2)
	}

	data3 := modelFullInvoice{
		ObjectConfig: objconfig,
		Insurer: modelParty{
			Name:    elementS{Value: "Ledgertech Insurer"},
			Address: elementS{Value: "Dreikönigstrsse 47"},
			City:    elementS{Value: "Zürich"},
			Country: elementS{Value: "Switzerland"},
			Phone:   elementS{Value: "+00000000000"},
			Email:   elementS{Value: "office@ledgertech.com"},
		},
		Customer: modelParty{
			Name:    elementS{Value: "אופירה עזריה"},
			Address: elementS{Value: "סמטאת רוטשילד"},
			City:    elementS{Value: "מסילות"},
			Country: elementS{Value: "ישראל"},
			Phone:   elementS{Value: "972-5-13346566"},
			Email:   elementS{Value: "customerHE-email@domain.com"},
		},
		Invoice: modelInvoice{
			Lang:          "he-IL",
			CurrencyISO:   "ILS",
			Number:        elementS{Value: "LT-1000003"},
			Date:          elementS{Value: "2022-05-03"},
			DaysToDueDate: elementI{Value: 0},
			InvoiceDetails: []modelInvoiceDetail{
				{
					DescriptionValue: "תן לתלמידים שלך תשובה. הם מעלים שאלה.",
					QuantityValue:    1,
					PriceValue:       610.00,
					AmountValue:      610.00,
				},
				{
					DescriptionValue: "ובדים היטב שכן השתמשתי בכל אחד מהם במשך 35 שנים",
					QuantityValue:    3,
					PriceValue:       210.00,
					AmountValue:      630.00,
				},
				{
					DescriptionValue: `כל אחד זכאי לכל הזכויות והחירויות המופיעות בהצהרה זו, ללא הבחנה מכל סוג שהוא,
					כגון גזע, צבע, מין, שפה, דת, דעה פוליטית או אחרת, מקור לאומי או חברתי, רכוש,
					לידה או מעמד אחר. יתר על כן, לא תיעשה הבחנה על בסיס הפוליטי, השיפוט או
					מעמד בינלאומי של המדינה או השטח שאדם משתייך אליו, בין אם זה עצמאי, אמון,
					לא ממשל עצמי או תחת כל מגבלה אחרת של ריבונות.`,
					QuantityValue: 3,
					PriceValue:    10.11,
					AmountValue:   30.33,
				},
			},
			InvoiceSummary: modelInvoiceSummary{
				DiscountPercent: 10.10,
				Discount:        10.00,
				VATPercent:      10.00,
				VAT:             10.00,
			},
		},
	}
	if objconfig.Lang == "he-IL" {
		allData = append(allData, &data3)
	}

	return allData
}

func GenerateFakeInvoices() {
	clog.Info("start generating test invoices")

	var err error
	objconf, err := checkSetDefaultValues(config.App)
	if err != nil {
		clog.Errorf("fail to check/set default values - error:%v\n", err)
		return
	}

	err = handleEmbededFonts()
	if err != nil {
		clog.Errorf("fail while handling embedded fonts - error: %+v\n", err)
		return
	}

	// get invoice raw data
	invoices := generateFakeData(objconf)

	for _, v := range invoices {
		//clog.Infof("invoice before:%s", v)

		err = v.prepareInvoiceForTemplate()
		if err != nil {
			clog.Errorf("fail to prepare invoice for template - error: %+v\n", err)
			return
		}
		//clog.Infof("invoice after:%s", v)

		_, err = v.createInvoicePDF()
		if err != nil {
			clog.Errorf("fail to create invoice - error: %+v\n", err)
			return
		}
	}
}
