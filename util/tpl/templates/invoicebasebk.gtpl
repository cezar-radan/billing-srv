{
    "paper": "A4P",
    "crop": "10",
    "origin": "UpperLeft",
    "contentBox": false,
    "guides": false,
    "colors": {
        "Beige": "#F5F5DC",
        "White": "#FFFFFF",
        "LightGrey": "#D3D3D3",
        "Black": "#000000",
        "DarkOrange": "#FF8C00",
        "DarkSeaGreen": "#8FBC8F",
        "LightBlue": "#85C2EB"
    },
    "bgcol": "$White",
    "timestamp": "2006-01-02 15:04:05",
    "dirs": {
        "images": "{{.ObjectConfig.ResourcePath}}"
    },
    "files": {
        "logo1": "$images/{{.ObjectConfig.SmallLogoName}}"
    },
    "fonts": {
        "myCourier": {
            "name": "Courier",
            "size": 10,
            "col": "#000000"
        },
        "myCourierLabel": {
            "name": "Courier-Bold",
            "size": 10,
            "col": "#000000"
        },
        "myCourierTitle": {
            "name": "Courier",
            "size": 25,
            "col": "#000000"
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
            "url": "{{.ObjectConfig.SmallLogoURL}}",
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
        "dx": 5,
        "dy": 10,
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
        "dx": 5,
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
                        "col": "black",
                        "style": "round"
                    }
                ],
                "box": [{
                        "pos": [0, 110],
                        "width": 495,
                        "height": 24,
                        "align": "left",
                        "fillCol": "$LightBlue"
                    }
                ],
                "text": [{
                        "hide": false,
                        "value": "INSURANCE\nINVOICE",
                        "pos": [600, 0],
                        "align": "right",
                        "font": {
                            "name": "$myCourierTitle"
                        }
                    },
                    {
                        "hide": false,
                        "value": "Payment is due within {{.Invoice.DaysToDueDate}} day.",
                        "pos": [0, 655],
                        "align": "left",
                        "font": {
                            "name": "$myCourier"
                        }
                    },
                    {
                        "hide": false,
                        "value": "Thank you for your business!",
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
                        "value": "",
                        "pos": [60, 11],
                        "width": 1,
                        "align": "left",
                        "label": {
                            "value": "Insurer",
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
                        "value": "{{.Insurer.Name}}",
                        "pos": [60, 22],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "Name:",
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
                        "value": "{{.Insurer.Address}}",
                        "pos": [60, 33],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "Address:",
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
                        "value": "{{.Insurer.City}}",
                        "pos": [60, 44],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "City:",
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
                        "value": "{{.Insurer.Country}}",
                        "pos": [60, 55],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "Country:",
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
                        "value": "{{.Insurer.Phone}}",
                        "pos": [60, 66],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "Phone:",
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
                        "value": "{{.Insurer.Email}}",
                        "pos": [60, 77],
                        "width": 300,
                        "align": "left",
                        "label": {
                            "value": "Email:",
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
                        "value": "",
                        "pos": [60, 99],
                        "width": 1,
                        "align": "left",
                        "label": {
                            "value": "Invoice",
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
                        "value": "{{.Invoice.Number}}",
                        "pos": [60, 111],
                        "width": 80,
                        "align": "left",
                        "label": {
                            "value": "Number:",
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
                        "value": "{{.Invoice.Date}}",
                        "pos": [400, 111],
                        "width": 80,
                        "align": "left",
                        "label": {
                            "value": "Date:",
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
                        "value": "",
                        "pos": [60, 133],
                        "width": 1,
                        "align": "left",
                        "label": {
                            "value": "Customer",
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
                        "value": "{{.Customer.Name}}",
                        "pos": [60, 144],
                        "width": 400,
                        "align": "left",
                        "label": {
                            "value": "Name:",
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
                        "value": "{{.Customer.Address}}",
                        "pos": [60, 155],
                        "width": 400,
                        "align": "left",
                        "label": {
                            "value": "Address:",
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
                        "value": "{{.Customer.City}}",
                        "pos": [60, 166],
                        "width": 400,
                        "align": "left",
                        "label": {
                            "value": "City:",
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
                        "value": "{{.Customer.Country}}",
                        "pos": [60, 177],
                        "width": 400,
                        "align": "left",
                        "label": {
                            "value": "Country:",
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
                        "value": "{{.Customer.Phone}}",
                        "pos": [60, 188],
                        "width": 400,
                        "align": "left",
                        "label": {
                            "value": "Phone:",
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
                        "value": "{{.Customer.Email}}",
                        "pos": [60, 199],
                        "width": 400,
                        "align": "left",
                        "label": {
                            "value": "Email:",
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
                            "values": ["Qty", "Description", "Price({{.Invoice.CurrencySymbol}})", "Amount({{.Invoice.CurrencySymbol}})"],
                            "colAnchors": ["Center", "Center", "Center", "Center"],
                            "bgCol": "$LightBlue",
                            "font": {
                                "name": "$myCourierLabel"
                            }
                        },
                        "values": [
                                    {{range $i, $ci := .Invoice.InvoiceDetails}}
                                    {{if $i}},{{end}}["{{$ci.Quantity}}", "{{$ci.Description}}", "{{$ci.Price}}", "{{$ci.Amount}}"]{{end}}
                                  ],
                        "rows": 4,
                        "width": 495,
                        "cols": 4,
                        "colWidths": [0.10, 0.5, 0.20, 0.20],
                        "colAnchors": ["Center", "Left", "Right", "Right"],
                        "lheight": 60,
                        "grid": true,
                        "//anchor": "right",
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
                            "width": 5
                        },
                        "rot": 0
                    }, {
                        "hide": false,
                        "values": [["SubTotal", "{{.Invoice.SubTotal}}({{.Invoice.CurrencySymbol}})"], ["Discount", "{{.Invoice.Discount}}({{.Invoice.CurrencySymbol}})"], ["VAT", "{{.Invoice.VAT}}({{.Invoice.CurrencySymbol}})"], ["Total", "{{.Invoice.Total}}({{.Invoice.CurrencySymbol}})"]],
                        "rows": 4,
                        "width": 198,
                        "cols": 2,
                        "colWidths": [0.50, 0.50],
                        "colAnchors": ["Left", "Right"],
                        "lheight": 25,
                        "grid": true,
                        "//anchor": "right",
                        "pos": [305, 635],
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