{
    "paper": "A4P",
    "crop": "10",
    "origin": "UpperRight",
    "contentBox": false,
    "guides": false,
    "colors": {
        "White": "#FFFFFF",
        "Black": "#000000",
        "Grey": "#808080",
        "LightGrey": "#D3D3D3",
        "LightBlue": "#85C2EB",
        "CustomColor": "{{.ObjectConfig.CustomColor}}"
    },
    "bgcol": "$White",
    "timestamp": "2006-01-02 15:04:05",
    "files": {
        "logo1": "{{.ObjectConfig.LogoPath}}"
    },
    "fonts": {
        "myOpenSansHebrew": {
            "name": "OpenSansHebrew-Regular",
            "size": 10,
            "col": "$Black"
        },
        "myOpenSansHebrewLabel": {
            "name": "OpenSansHebrew-Bold",
            "size": 10,
            "col": "$Black"
        },
        "myOpenSansHebrewTitle": {
            "name": "OpenSansHebrew-Regular",
            "size": 25,
            "col": "$Black"
        }
    },
    "borders": {
        "myBorder": {
            "width": 5
        }
    },
    "margin": {
        "width": 40
    },
    "images": {
        "logo1": {
            "file": "$logo1",
            "url": "{{.ObjectConfig.LogoURL}}",
            "margin": {
                "width": 5
            }
        }
    },
    "header": {
        "font": {
            "name": "$myOpenSansHebrew"
        },
        "rtl": true,
        "left": "",
        "center": "",
        "right":  "$logo1",
        "height": 30,
        "dx": 10,
        "dy": 5,
        "border": false
    },
    "footer": {
        "font": {
            "name": "$myOpenSansHebrew"
        },
        "rtl": true,
        "left": "",
        "center": "",
        "right": "\n%t",
        "height": 30,
        "dx": 10,
        "dy": 5,
        "border": false
    },
    "pages": {
        "1": {
            "content": {
                "guides": [ 
                                    {
                        "pos": [1, 1]
                    },
                    {
                        "pos": [400, 400]
                    },
                    {
                        "pos": [500, 500]
                    }
                ],
                "box": [{
                        "pos": [0, 110],
                        "width": 495,
                        "height": 24,
                        "align": "right",
                        "fillCol": "$CustomColor"
                    }
                ],
                "text": [{
                        "hide": false,
                        "value": "{{.Invoice.CustomText}}",
                        "pos": [600, 0],
                        "align": "right",
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrewTitle"
                        }
                    },
                    {
                        "hide": false,
                        "value": "{{.Invoice.DaysToDueDate.Text}}",
                        "pos": [0, 655],
                        "align": "left",
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    },
                    {
                        "hide": false,
                        "value": "{{.Invoice.ExitMessage}}",
                        "pos": [0, 670],
                        "align": "left",
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }
                ],
                "textfield": [{
                        "hide": false,
                        "id": "insHead",
                        "value": "{{.Insurer.Head.Value}}",
                        "pos": [60, 11],
                        "width": 1,
                        "align": "right",
                        "label": {
                            "value": "{{.Insurer.Head.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "insName",
                        "value": "{{.Insurer.Name.Value}}",
                        "pos": [360, 22],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Insurer.Name.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "insAddress",
                        "value": "{{.Insurer.Address.Value}}",
                        "pos": [360, 33],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Insurer.Address.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "insCity",
                        "value": "{{.Insurer.City.Value}}",
                        "pos": [360, 44],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Insurer.City.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "insCountry",
                        "value": "{{.Insurer.Country.Value}}",
                        "pos": [360, 55],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Insurer.Country.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "insPhone",
                        "value": "{{.Insurer.Phone.Value}}",
                        "pos": [360, 66],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Insurer.Phone.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "insEmail",
                        "value": "{{.Insurer.Email.Value}}",
                        "pos": [360, 77],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Insurer.Email.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "invHead",
                        "value": "{{.Invoice.Head.Value}}",
                        "pos": [60, 99],
                        "width": 1,
                        "align": "right",
                        "label": {
                            "value": "{{.Invoice.Head.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "invNumber",
                        "value": "{{.Invoice.Number.Value}}",
                        "pos": [160, 111],
                        "width": 100,
                        "align": "right",
                        "label": {
                            "value": "{{.Invoice.Number.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "invDate",
                        "value": "{{.Invoice.Date.Value}}",
                        "pos": [500, 111],
                        "width": 100,
                        "align": "right",
                        "label": {
                            "value": "{{.Invoice.Date.Text}}", 
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "cusHead",
                        "value": "{{.Customer.Head.Value}}",
                        "pos": [60, 133],
                        "width": 1,
                        "align": "right",
                        "label": {
                            "value": "{{.Customer.Head.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "cusName",
                        "value": "{{.Customer.Name.Value}}",
                        "pos": [360, 144],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Customer.Name.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "cusAddress",
                        "value": "{{.Customer.Address.Value}}",
                        "pos": [360, 155],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Customer.Address.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "cusCity",
                        "value": "{{.Customer.City.Value}}",
                        "pos": [360, 166],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Customer.City.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "cusCountry",
                        "value": "{{.Customer.Country.Value}}",
                        "pos": [360, 177],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Customer.Country.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "cusPhone",
                        "value": "{{.Customer.Phone.Value}}",
                        "pos": [360, 188],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Customer.Phone.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }, {
                        "hide": false,
                        "id": "cusEmail",
                        "value": "{{.Customer.Email.Value}}",
                        "pos": [360, 199],
                        "width": 300,
                        "align": "right",
                        "label": {
                            "value": "{{.Customer.Email.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "right",
                            "pos": "right",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        }
                    }
                ],
                "table": [{
                        "hide": false,
                        "header": {
                        "values": ["{{.Invoice.InvoiceDetailsHead.AmountText}}", "{{.Invoice.InvoiceDetailsHead.PriceText}}", "{{.Invoice.InvoiceDetailsHead.DescriptionText}}", "{{.Invoice.InvoiceDetailsHead.QuantityText}}"],
                            "colAnchors": ["Center", "Center", "Center", "Center"],
                            "bgCol": "$CustomColor",
                            "rtl": false,
                            "font": {
                                "name": "$myOpenSansHebrewLabel"
                            }
                        },
                        "values": [
                                    {{range $i, $ci := .Invoice.InvoiceDetails}}
                                    {{if $i}},{{end}}["{{$ci.AmountValue}}", "{{$ci.PriceValue}}", "{{$ci.DescriptionValue}}", "{{$ci.QuantityValue}}"]{{end}}
                                  ],
                        "rows": 3,
                        "lheight": 80,
                        "width": 494,
                        "cols": 4,
                        "colWidths": [0.20, 0.20, 0.5, 0.10],
                        "colAnchors": ["Left", "Left", "Right", "Center"],
                        "grid": true,
                        "//anchor": "right",
                        "pos": [0, 530],
                        "align": "right",
                        "dx": 0,
                        "dy": 0,
                        "bgCol": "$White",
                        "//oddCol": "$LightGrey",
                        "//evenCol": "$Grey",
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrew"
                        },
                        "margin": {
                            "width": 1
                        },
                        "border": {
                            "width": 1
                        },
                        "padding": {
                            "width": 0
                        },
                        "rot": 0
                    }, {
                        "hide": false,
                        "values": [
                                {{ if gt .Invoice.InvoiceSummary.SubTotal 0.00}} ["{{.Invoice.InvoiceSummary.SummarySubTotal.Value}}", "{{.Invoice.InvoiceSummary.SummarySubTotal.Text}}"], {{end}}
                                {{ if gt .Invoice.InvoiceSummary.Discount 0.00}} ["{{.Invoice.InvoiceSummary.SummaryDiscount.Value}}", "{{.Invoice.InvoiceSummary.SummaryDiscount.Text}}"], {{end}}
                                {{ if gt .Invoice.InvoiceSummary.VAT 0.00}}["{{.Invoice.InvoiceSummary.SummaryVAT.Value}}", "{{.Invoice.InvoiceSummary.SummaryVAT.Text}}"], {{end}} 
                                ["{{.Invoice.InvoiceSummary.SummaryTotal.Value}}", "{{.Invoice.InvoiceSummary.SummaryTotal.Text}}"]
                        ],
                        "rows": {{.Invoice.InvoiceSummary.SummaryRows}},
                        "width": 198,
                        "cols": 2,
                        "colWidths": [0.50, 0.50],
                        "colAnchors": ["Left", "Right"],
                        "lheight": 25,
                        "grid": true,
                        "//anchor": "right",
                        "pos": [510,
                        {{ if eq .Invoice.InvoiceSummary.SummaryRows 1}} 560 {{end}} 
                        {{ if eq .Invoice.InvoiceSummary.SummaryRows 2}} 585 {{end}} 
                        {{ if eq .Invoice.InvoiceSummary.SummaryRows 3}} 610 {{end}} 
                        {{ if eq .Invoice.InvoiceSummary.SummaryRows 4}} 635 {{end}} 
                        ],
                        "align": "right",
                        "dx": 0,
                        "dy": 0,
                        "bgCol": "$White",
                        "rtl": false,
                        "font": {
                            "name": "$myOpenSansHebrewLabel"
                        },
                        "margin": {
                            "width": 1
                        },
                        "border": {
                            "width": 1
                        },
                        "padding": {
                            "width": 5
                        },
                        "rot": 0
                    }
                ]
            }
        }
    }
}