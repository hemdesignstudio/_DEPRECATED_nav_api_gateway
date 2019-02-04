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
	args := Args{}
	args.FilterArgs = SliceOfMaps{{
		"key":   "No",
		"value": "12230",
	},
	}

	args.CreateArgs = SliceOfMaps{{
		"No":                     "12233",
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

	return args
}
