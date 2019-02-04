package utils

type Args struct {
	FilterArgs SliceOfMaps
	CreateArgs SliceOfMaps
	UpdateArgs SliceOfMaps
}

func GetAssemblyBomAttrs() []string {
	attrs := []string{
		"No",
		"Parent_Item_No",
		"No",
		"Type",
		"Quantity_per",
		"Unit_of_Measure_Code",
	}
	return attrs
}

func GetAssemblyBomArgs() SliceOfMaps {
	args := SliceOfMaps{
		{
			"key":   "No",
			"value": "40001",
		},
	}

	return args
}

func GetCustomerCardAttrs() []string {
	attrs := []string{
		"No",
		"Name",
		"Address",
		"Address_2",
		"Post_Code",
		"City",
		"Country_Region_Code",
		"Phone_No",
		"Contact",
		"VAT_Registration_No",
		"Customer_Posting_Group",
		"Gen_Bus_Posting_Group",
		"VAT_Bus_Posting_Group",
		"Customer_Price_Group",
		"Customer_Disc_Group",
		"Payment_Terms_Code",
		"Currency_Code",
		"Language_Code",
		"Web_E_Mail",
		"Web_Customer",
	}
	return attrs
}

func GetCustomerCardArgs() Args {
	var id = "12233"
	args := Args{}
	args.FilterArgs = SliceOfMaps{{
		"key":   "No",
		"value": id,
	},
	}

	args.CreateArgs = SliceOfMaps{{
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
	}}

	args.UpdateArgs = SliceOfMaps{{
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

	}}

	return args
}

func GetItemCardAttrs() []string {
	attrs := []string{
		"No",
		"Description",
		"Base_Unit_of_Measure",
		"Assembly_BOM",
		"Item_Category_Code",
		"Product_Group_Code",
		"Inventory",
		"Qty_on_Purch_Order",
		"Qty_on_Sales_Order",
		"Last_Date_Modified",
		"Freight_Type",
		"Assembly_Policy",
		"Country_Region_of_Origin_Code",
		"Net_Weight",
		"Gross_Weight",
		"Unit_Volume",
		"Length",
		"Width",
		"Height",
		"Designer",
		"Web_Status",
	}
	return attrs

}

func GetItemCardArgs() Args {
	var id = "11233"
	args := Args{}
	args.FilterArgs = SliceOfMaps{{
		"key":   "No",
		"value": id,
	}}

	args.CreateArgs = SliceOfMaps{{
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

	}}

	args.UpdateArgs = SliceOfMaps{{
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
	}}

	return args
}
