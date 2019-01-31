package test

func getAssemblyBomAttrs() []string {
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

func getAssemblyBomArgs() []map[string]interface{} {
	args := []map[string]interface{}{
		{
			"key":   "No",
			"value": "40001",
		},
	}

	return args
}

func getCustomerCardAttrs() []string {
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
		"Global_Dimension_2_Code",
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

func getCustomerCardArgs() []map[string]interface{} {
	args := []map[string]interface{}{
		{
			"key":   "No",
			"value": "12345",
		},
	}

	return args
}
