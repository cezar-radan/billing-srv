{
    "paper": "A4P",
    "crop": "10",
    "origin": "UpperLeft",
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
        "myCourier": {
            "name": "Courier",
            "size": 10,
            "col": "$Black"
        },
        "myCourierLabel": {
            "name": "Courier-Bold",
            "size": 10,
            "col": "$Black"
        },
        "myCourierTitle": {
            "name": "Courier",
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
            "name": "$myCourier"
        },
        "left": "$logo1",
        "center": "",
        "right": "",
        "height": 30,
        "dx": 10,
        "dy": 5,
        "border": false
    },
    "footer": {
        "font": {
            "name": "$myCourier"
        },
        "left": "Created:\n%t",
        "center": "Page %p of %P",
        "right": "",
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
                        "pos": [60, 60]
                    },
                    {
                        "pos": [200, 200]
                    },
                    {
                        "pos": [360, 360]
                    },
                    {
                        "pos": [400, 400]
                    },
                                        {
                        "pos": [495, 700]
                    }
                ],
                "//bar": [{
                        "hide": false,
                        "y": 75,
                        "width": 0,
                        "col": "$Black",
                        "style": "round"
                    }
                ],
                "box": [{
                        "pos": [0, 110],
                        "width": 495,
                        "height": 24,
                        "align": "left",
                        "fillCol": "$CustomColor"
                    }
                ],
                "text": [{
                        "hide": false,
                        "value": "{{.Invoice.CustomText}}",
                        "pos": [600, 0],
                        "align": "left",
                        "font": {
                            "name": "$myCourierTitle"
                        }
                    },
                    {
                        "hide": false,
                        "value": "{{.Invoice.DaysToDueDate.Text}}",
                        "pos": [0, 655],
                        "align": "left",
                        "font": {
                            "name": "$myCourier"
                        }
                    },
                    {
                        "hide": false,
                        "value": "{{.Invoice.ExitMessage}}",
                        "pos": [0, 670],
                        "align": "left",
                        "font": {
                            "name": "$myCourier"
                        }
                    }
                ],
                "textfield": [{
                        "hide": false,
                        "id": "insHead",
                        "value": "{{.Insurer.Head.Value}}",
                        "pos": [60, 11],
                        "width": 1,
                        "align": "left",
                        "label": {
                            "value": "{{.Insurer.Head.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "insName",
                        "value": "{{.Insurer.Name.Value}}",
                        "pos": [60, 22],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Insurer.Name.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "insAddress",
                        "value": "{{.Insurer.Address.Value}}",
                        "pos": [60, 33],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Insurer.Address.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "insCity",
                        "value": "{{.Insurer.City.Value}}", 
                        "pos": [60, 44],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Insurer.City.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "insCountry",
                        "value": "{{.Insurer.Country.Value}}",
                        "pos": [60, 55],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Insurer.Country.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "insPhone",
                        "value": "{{.Insurer.Phone.Value}}",
                        "pos": [60, 66],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Insurer.Phone.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "insEmail",
                        "value": "{{.Insurer.Email.Value}}",
                        "pos": [60, 77],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Insurer.Email.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "invHead",
                        "value": "{{.Invoice.Head.Value}}",
                        "pos": [60, 99],
                        "width": 1,
                        "align": "left",
                        "label": {
                            "value": "{{.Invoice.Head.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "invNumber",
                        "value": "{{.Invoice.Number.Value}}",
                        "pos": [60, 111],
                        "width": 100,
                        "align": "left",
                        "label": {
                            "value": "{{.Invoice.Number.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "invDate",
                        "value": "{{.Invoice.Date.Value}}",
                        "pos": [400, 111],
                        "width": 100,
                        "align": "left",
                        "label": {
                            "value": "{{.Invoice.Date.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "cusHead",
                        "value": "{{.Customer.Head.Value}}",
                        "pos": [60, 133],
                        "width": 1,
                        "align": "left",
                        "label": {
                            "value": "{{.Customer.Head.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "cusName",
                        "value": "{{.Customer.Name.Value}}",
                        "pos": [60, 144],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Customer.Name.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "cusAddress",
                        "value": "{{.Customer.Address.Value}}",
                        "pos": [60, 155],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Customer.Address.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "cusCity",
                        "value": "{{.Customer.City.Value}}",
                        "pos": [60, 166],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Customer.City.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "cusCountry",
                        "value": "{{.Customer.Country.Value}}",
                        "pos": [60, 177],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Customer.Country.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "cusPhone",
                        "value": "{{.Customer.Phone.Value}}",
                        "pos": [60, 188],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Customer.Phone.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }, {
                        "hide": false,
                        "id": "cusEmail",
                        "value": "{{.Customer.Email.Value}}",
                        "pos": [60, 199],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "{{.Customer.Email.Text}}",
                            "width": 60,
                            "gap": 0,
                            "align": "left",
                            "pos": "left",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "font": {
                            "name": "$myCourier"
                        }
                    }
                ],
                "table": [{
                        "hide": false,
                        "header": {
                            "values": ["{{.Invoice.InvoiceDetailsHead.QuantityText}}", "{{.Invoice.InvoiceDetailsHead.DescriptionText}}", "{{.Invoice.InvoiceDetailsHead.PriceText}}", "{{.Invoice.InvoiceDetailsHead.AmountText}}"],
                            "colAnchors": ["Center", "Center", "Center", "Center"],
                            "bgCol": "$CustomColor",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "values": [
                                    {{range $i, $ci := .Invoice.InvoiceDetails}}
                                    {{if $i}},{{end}}["{{$ci.QuantityValue}}", "{{$ci.DescriptionValue}}", "{{$ci.PriceValue}}", "{{$ci.AmountValue}}"]{{end}}
                                  ],
                        "rows": 3,
                        "lheight": 80,
                        "width": 494,
                        "cols": 4,
                        "colWidths": [0.10, 0.5, 0.20, 0.20],
                        "colAnchors": ["Center", "Left", "Right", "Right"],
                        "grid": true,
                        "pos": [0, 530],
                        "align": "left",
                        "dx": 0,
                        "dy": 0,
                        "bgCol": "$White",
                        "//oddCol": "$LightGrey",
                        "//evenCol": "$Grey",
                        "font": {
                            "name": "$myCourier"
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
                                   {{ if gt .Invoice.InvoiceSummary.SubTotal 0.00}} ["{{.Invoice.InvoiceSummary.SummarySubTotal.Text}}", "{{.Invoice.InvoiceSummary.SummarySubTotal.Value}}"], {{end}}
                                   {{ if gt .Invoice.InvoiceSummary.Discount 0.00}} ["{{.Invoice.InvoiceSummary.SummaryDiscount.Text}}", "{{.Invoice.InvoiceSummary.SummaryDiscount.Value}}"], {{end}}
                                   {{ if gt .Invoice.InvoiceSummary.VAT 0.00}}["{{.Invoice.InvoiceSummary.SummaryVAT.Text}}", "{{.Invoice.InvoiceSummary.SummaryVAT.Value}}"], {{end}} 
                                   ["{{.Invoice.InvoiceSummary.SummaryTotal.Text}}", "{{.Invoice.InvoiceSummary.SummaryTotal.Value}}"]
                        ],
                        "rows": {{.Invoice.InvoiceSummary.SummaryRows}},
                        "width": 198,
                        "cols": 2,
                        "colWidths": [0.50, 0.50],
                        "colAnchors": ["Left", "Right"],
                        "lheight": 25,
                        "grid": true,
                        "//anchor": "right",
                        "pos": [305,
                        {{ if eq .Invoice.InvoiceSummary.SummaryRows 1}} 560 {{end}} 
                        {{ if eq .Invoice.InvoiceSummary.SummaryRows 2}} 585 {{end}} 
                        {{ if eq .Invoice.InvoiceSummary.SummaryRows 3}} 610 {{end}} 
                        {{ if eq .Invoice.InvoiceSummary.SummaryRows 4}} 635 {{end}} 
                        ],
                        "align": "left",
                        "dx": 0,
                        "dy": 0,
                        "bgCol": "$White",
                        "font": {
                            "name": "$myCourierLabel"
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