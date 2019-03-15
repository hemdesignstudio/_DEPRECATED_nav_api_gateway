package utils

type ArgType map[string]interface{}

type Args struct {
	FilterArgs map[string]interface{}
	CreateArgs map[string]interface{}
	UpdateArgs map[string]interface{}
}

func GetAssemblyBomArgs() Args {
	var id = "40001"
	args := Args{}
	args.FilterArgs = ArgType{
		"key":   "No",
		"value": id,
	}

	return args
}

func GetInventoryArgs() Args {
	var id = "10005"
	args := Args{}
	args.FilterArgs = ArgType{
		"key":   "No",
		"value": id,
	}

	return args
}

func GetCustomerCardArgs() Args {
	var id = "12233"
	args := Args{}
	args.FilterArgs = ArgType{
		"key":   "No",
		"value": id,
	}

	args.CreateArgs = ArgType{
		"No":                     id,
		"Name":                   "qraphql Test",
		"Address":                "Endla 80",
		"Address_2":              "Another Address",
		"Post_Code":              "11431",
		"City":                   "Tallinn",
		"Country_Region_Code":    "EE",
		"Contact":                "883737",
		"VAT_Registration_No":    "123",
		"Customer_Posting_Group": "PROJECT",
		"Gen_Bus_Posting_Group":  "DOMESTIC",
		"VAT_Bus_Posting_Group":  "PL-EU",
		"Customer_Price_Group":   "PROJECT",
		"Customer_Disc_Group":    "PROJECT",
		"Payment_Terms_Code":     "CIA",
		"Currency_Code":          "EUR",
		"Language_Code":          "ENU",
		"Web_E_Mail":             "ahmed@test.com",
		"Web_Customer":           true,
	}

	args.UpdateArgs = ArgType{
		"No":                     id,
		"Name":                   "qraphql Test update",
		"Address":                "Endla 80, update",
		"Address_2":              "Another Address in Tallinn",
		"Post_Code":              "11431",
		"City":                   "Tallinn",
		"Country_Region_Code":    "EE",
		"VAT_Registration_No":    "12345",
		"Customer_Posting_Group": "PROJECT",
		"Gen_Bus_Posting_Group":  "DOMESTIC",
		"VAT_Bus_Posting_Group":  "PL-EU",
		"Customer_Price_Group":   "PROJECT",
		"Customer_Disc_Group":    "PROJECT",
		"Payment_Terms_Code":     "CIA",
		"Currency_Code":          "EUR",
		"Language_Code":          "ENU",
		"Web_E_Mail":             "ahmed@test.com",
		"Web_Customer":           true,
		//read only fields
		//"Contact":                "8837372",

	}

	return args
}

func GetItemCardArgs() Args {
	var id = "11233"
	args := Args{}
	args.FilterArgs = ArgType{
		"key":   "No",
		"value": id,
	}

	args.CreateArgs = ArgType{
		"No":                            id,
		"Assembly_BOM":                  false,
		"Assembly_Policy":               "Assemble-to-Stock",
		"Base_Unit_of_Measure":          "PCS",
		"Country_Region_of_Origin_Code": "PL",
		"Description":                   "Hai Chair Mosaic Charcoal Test",
		"Designer":                      "Luca Nichetto",
		"Freight_Type":                  "FREIGHT",
		"Gross_Weight":                  27,
		"Height":                        106.4,
		"Item_Category_Code":            "LOUNGE CH",
		"Length":                        94.4,
		"Net_Weight":                    0,
		"Product_Group_Code":            "LOUNGE CH",
		"Unit_Volume":                   0.948,
		"Web_Status":                    "On",
		"Width":                         94.4,
		//Read Only Fields
		//"Qty_on_Sales_Order":            17,
		//"Qty_on_Purch_Order":            25,
		//"Inventory":                     17,
		//"Last_Date_Modified": "2019-01-16",

	}

	args.UpdateArgs = ArgType{
		"No":                            id,
		"Assembly_Policy":               "Assemble-to-Stock",
		"Base_Unit_of_Measure":          "PCS",
		"Country_Region_of_Origin_Code": "LV",
		"Description":                   "Hai Chair Mosaic Charcoal Test Update",
		"Designer":                      "Luca Nichetto",
		"Freight_Type":                  "PARCEL",
		"Gross_Weight":                  27,
		"Height":                        8.1,
		"Item_Category_Code":            "HANGING MI",
		"Length":                        94.4,
		"Net_Weight":                    9,
		"Product_Group_Code":            "HANGING MI",
		"Unit_Volume":                   0.948,
		"Web_Status":                    "On",
		"Width":                         94.4,
		//Read Only Fields
		//"Assembly_BOM":                  true,
		//"Qty_on_Sales_Order":            17,
		//"Qty_on_Purch_Order":            25,
		//"Inventory":                     17,
		//"Last_Date_Modified": "2019-01-16",
	}

	return args
}

func GetSalesOrderArgs() Args {
	var id = "11233"
	args := Args{}
	args.FilterArgs = ArgType{
		"key":   "No",
		"value": id,
	}

	args.CreateArgs = ArgType{
		"No":                              id,
		"Assigned_User_ID":                "",
		"Bill_to_Address":                 "Avenue road bla 3",
		"Bill_to_City":                    "London",
		"Bill_to_Name":                    "Test Name",
		"Bill_to_Post_Code":               "133 37",
		"Compress_Prepayment":             false,
		"Currency_Code":                   "EUR",
		"Customer_No_Web":                 "oscar+test13333337@hem.com",
		"Document_Date":                   "2018-10-26",
		"Due_Date":                        "2018-10-26",
		"EU_3_Party_Trade":                false,
		"External_Document_No":            "1001",
		"Job_Queue_Status":                " ",
		"Location_Code":                   "LONDON",
		"No_of_Archived_Versions":         0,
		"Order_Date":                      "2018-10-26",
		"Payment_Discount_Percent":        0,
		"Payment_Method_Code":             "BRAINTREE",
		"Pmt_Discount_Date":               "2018-10-26",
		"Posting_Date":                    "2018-10-26",
		"Prepayment_Due_Date":             "2018-10-26",
		"Prepayment_Percent":              100,
		"Prepmt_Payment_Discount_Percent": 0,
		"Prepmt_Pmt_Discount_Date":        "2018-10-26",
		"Prices_Including_VAT":            true,
		"Promised_Delivery_Date":          "2018-11-02",
		"Requested_Delivery_Date":         "0001-01-01",
		"Salesperson_Code":                "WEB",
		"Sell_to_Address":                 "",
		"Sell_to_City":                    "",
		"Sell_to_Contact":                 "",
		"Sell_to_Customer_Name":           "Web Customer Currency EUR",
		"Sell_to_Post_Code":               "",
		"Ship_to_Address":                 "Test Road 2",
		"Ship_to_Address_2":               "",
		"Ship_to_City":                    "Test City",
		"Ship_to_Country_Region_Code":     "SE",
		"Ship_to_Name":                    "Oscar Alsingg",
		"Ship_to_Post_Code":               "133 37",
		"Shipment_Date":                   "2018-10-26",
		"Shortcut_Dimension_2_Code":       "310",
		"VAT_Bus_Posting_Group":           "SV-DOM",
		"Whs_Shipment_Lines_Exists":       false,

		//Read Only fields
		//"Status":                    "Released",
		//"Shipping_Advice":                 "Complete",
		//"Sell_to_Customer_No": "WEB-10",
		//"Late_Order_Shipping":             true,

	}

	args.UpdateArgs = ArgType{
		"No":                              id,
		"Assigned_User_ID":                "",
		"Bill_to_Address":                 "Avenue road bla 3",
		"Bill_to_City":                    "London",
		"Bill_to_Name":                    "Test Name 2",
		"Bill_to_Post_Code":               "133 37",
		"Compress_Prepayment":             false,
		"Currency_Code":                   "EUR",
		"Customer_No_Web":                 "test@hem.com",
		"Document_Date":                   "2018-10-26",
		"Due_Date":                        "2018-10-26",
		"EU_3_Party_Trade":                false,
		"External_Document_No":            "1001",
		"Job_Queue_Status":                " ",
		"Location_Code":                   "LONDON",
		"No_of_Archived_Versions":         0,
		"Order_Date":                      "2018-10-26",
		"Payment_Discount_Percent":        0,
		"Payment_Method_Code":             "BRAINTREE",
		"Pmt_Discount_Date":               "2018-10-26",
		"Posting_Date":                    "2018-10-26",
		"Prepayment_Due_Date":             "2018-10-26",
		"Prepayment_Percent":              100,
		"Prepmt_Payment_Discount_Percent": 0,
		"Prepmt_Pmt_Discount_Date":        "2018-10-26",
		"Prices_Including_VAT":            true,
		"Promised_Delivery_Date":          "2018-11-02",
		"Requested_Delivery_Date":         "0001-01-01",
		"Salesperson_Code":                "WEB",
		"Sell_to_Address":                 "",
		"Sell_to_City":                    "",
		"Sell_to_Contact":                 "",
		"Sell_to_Customer_No":             "WEB-10",
		"Sell_to_Customer_Name":           "Web Customer Currency EUR",
		"Sell_to_Post_Code":               "",
		"Ship_to_Address":                 "Test Road 2",
		"Ship_to_Address_2":               "",
		"Ship_to_City":                    "Test City",
		"Ship_to_Country_Region_Code":     "SE",
		"Ship_to_Name":                    "Oscar Alsingg",
		"Ship_to_Post_Code":               "133 37",
		"Shipment_Date":                   "2018-10-26",
		"Shortcut_Dimension_2_Code":       "310",
		"VAT_Bus_Posting_Group":           "SV-DOM",
		"Whs_Shipment_Lines_Exists":       false,

		//Read Only fields
		//"Status":                    "Released",
		//"Shipping_Advice":                 "Complete",
		//"Late_Order_Shipping": true,
	}
	return args
}

func GetSalesInvoiceArgs() Args {
	args := Args{}
	args.FilterArgs = ArgType{
		"key":   "No",
		"value": "",
	}
	args.CreateArgs = ArgType{

		//Read Only Fields
		//"Bill_to_Address":            "Klippgatan 14 I",
		//"Bill_to_Address_2":          "",
		//"Bill_to_City":               "Stockholm",
		//"Bill_to_Contact":            "Tove Sjöberg",
		//"Bill_to_Contact_No":         "",
		//"Bill_to_Post_Code":          "116 35",
		//"Sell_to_Contact":            "Tove Sjöberg",
		//"Sell_to_Customer_No":       "206901",
		//"Ship_to_Address":  		   "Klippgatan 14 I",
		//"Ship_to_City":     		   "Stockholm",
		//"Ship_to_Code":              "",
		//"Ship_to_Name":              "Studio Feuer AB",
		//"Ship_to_Post_Code":         "116 35",

		"Area":                       "",
		"Assigned_User_ID":           "",
		"Bill_to_Name":               "Studio Feuer AB",
		"Campaign_No":                "",
		"Currency_Code":              "EUR",
		"Direct_Debit_Mandate_ID":    "",
		"Document_Date":              "2016-05-03",
		"Due_Date":                   "2016-05-17",
		"EU_3_Party_Trade":           true,
		"Exit_Point":                 "",
		"External_Document_No":       "1001",
		"Incoming_Document_Entry_No": 0,
		"Job_Queue_Status":           " ",
		"Location_Code":              "MIERZYN",
		"PEB_Note_of_Goods":          "sdwsd",
		"Package_Tracking_No":        "123445",
		"Payment_Discount_Percent":   0,
		"Payment_Method_Code":        "ACCOUNT",
		"Payment_Terms_Code":         "14 DAYS",
		"Pmt_Discount_Date":          "2016-05-03",
		"Posting_Date":               "2016-05-03",
		"Prices_Including_VAT":       true,
		"Responsibility_Center":      "",
		"Salesperson_Code":           "",
		"Sell_to_Address":            "Klippgatan 14 I",
		"Sell_to_Address_2":          "asa",
		"Sell_to_City":               "Stockholm",
		"Sell_to_Contact_No":         "",
		"Sell_to_Customer_Name":      "Studio Feuer AB",
		"Sell_to_Post_Code":          "116 35",
		"Ship_to_Address_2":          "",
		"Ship_to_Contact":            "Tove Sjöberg",
		"Shipment_Date":              "2016-05-03",
		"Shipment_Method_Code":       "DPD",
		"Shipping_Agent_Code":        "",
		"Shortcut_Dimension_1_Code":  "",
		"Shortcut_Dimension_2_Code":  "110",
		"Status":                     "Open",
		"Transaction_Specification":  "",
		"Transaction_Type":           "",
		"Transport_Method":           "",
		"VAT_Bus_Posting_Group":      "SV-DOM",
		"Your_Reference":             "asas",
	}

	args.UpdateArgs = ArgType{
		/* Read Only Fields
		"Bill_to_Address":            "Klippgatan 14 I",
		"Bill_to_Address_2":          "",
		"Bill_to_City":               "Stockholm",
		"Bill_to_Contact":            "Tove Sjöberg",
		"Bill_to_Contact_No":         "",
		"Bill_to_Post_Code":          "116 35",
		"Job_Queue_Status":           " ",
		"Sell_to_Contact":            "Tove Sjöberg",
		"Ship_to_Address_2":          "",
		"Ship_to_City":               "Stockholm",
		"Ship_to_Name":               "Studio Feuer AB",
		"Ship_to_Post_Code":          "116 35",
		*/

		"Area":                       "",
		"Assigned_User_ID":           "",
		"Bill_to_Name":               "Studio Feuer AB",
		"Campaign_No":                "",
		"Currency_Code":              "",
		"Direct_Debit_Mandate_ID":    "",
		"Document_Date":              "2016-05-03",
		"Due_Date":                   "2016-05-17",
		"EU_3_Party_Trade":           false,
		"Exit_Point":                 "",
		"External_Document_No":       "1001",
		"Incoming_Document_Entry_No": 0,
		"Location_Code":              "MIERZYN",
		"PEB_Note_of_Goods":          "1234",
		"Package_Tracking_No":        "123445",
		"Payment_Discount_Percent":   0,
		"Payment_Method_Code":        "ACCOUNT",
		"Payment_Terms_Code":         "14 DAYS",
		"Pmt_Discount_Date":          "2016-05-03",
		"Posting_Date":               "2016-05-03",
		"Prices_Including_VAT":       false,
		"Responsibility_Center":      "",
		"Salesperson_Code":           "",
		"Sell_to_Address":            "Klippgatan 14 I",
		"Sell_to_Address_2":          "aaa",
		"Sell_to_City":               "Stockholm",
		"Sell_to_Contact_No":         "",
		"Sell_to_Customer_Name":      "Studio Feuer AB",
		"Sell_to_Customer_No":        "206901",
		"Sell_to_Post_Code":          "116 35",
		"Ship_to_Address":            "Klippgatan 14 I",
		"Ship_to_Code":               "",
		"Ship_to_Contact":            "Tove Sjöberg",
		"Shipment_Date":              "2016-05-03",
		"Shipment_Method_Code":       "DPD",
		"Shipping_Agent_Code":        "",
		"Shortcut_Dimension_1_Code":  "110",
		"Shortcut_Dimension_2_Code":  "110",
		"Status":                     "Open",
		"Transaction_Specification":  "",
		"Transaction_Type":           "",
		"Transport_Method":           "",
		"VAT_Bus_Posting_Group":      "SV-DOM",
		"Your_Reference":             "ref",
	}

	return args
}

func GetPostShipArgs() Args {
	args := Args{}
	args.FilterArgs = ArgType{
		"key":   "No",
		"value": "102005",
	}
	return args
}

func GetSalesLineArgs() Args {
	id := "12971"
	args := Args{}
	args.FilterArgs = ArgType{
		"key":   "No",
		"value": id,
	}

	args.CreateArgs = ArgType{
		//Read Only fields
		//"Unit_of_Measure_Code"
		"Description":           "Alle Table Coffee Large Natural Oak",
		"Document_No":           "2152P",
		"Document_Type":         "Order",
		"Line_Amount":           478.4,
		"Line_Discount_Amount":  119.6,
		"Line_Discount_Percent": 20,
		"Line_No":               40000,
		"No":                    id,
		"Prepayment_Percent":    100,
		"Prepmt_Line_Amount":    478.4,
		"Qty_to_Invoice":        0,
		"Qty_to_Ship":           0,
		"Quantity":              2,
		"Quantity_Invoiced":     0,
		"Quantity_Shipped":      0,
		"Reserve":               "Optional",
		"Reserved_Quantity":     0,
		"Type":                  "Item",
		"Unit_Price":            299,
	}

	args.UpdateArgs = ArgType{
		//Read Only fields
		//"Unit_of_Measure_Code"

		//"No":                    id,
		//"Quantity_Invoiced":     0,
		//"Quantity_Shipped":      0,
		//"Reserve":               "Optional",
		//"Reserved_Quantity":     0,

		"Document_No":           "2152P",
		"Document_Type":         "Order",
		"Line_No":               40000,
		"Description":           "Alle Table Coffee Large Natural Oak",
		"Line_Amount":           400.5,
		"Line_Discount_Amount":  121.6,
		"Line_Discount_Percent": 20,
		"Prepayment_Percent":    90,
		"Prepmt_Line_Amount":    400.5,
		"Qty_to_Invoice":        0,
		"Qty_to_Ship":           0,
		"Quantity":              2,
		"Type":                  "Item",
		"Unit_Price":            299,
	}

	return args
}
