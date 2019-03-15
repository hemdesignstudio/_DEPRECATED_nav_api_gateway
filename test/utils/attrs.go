package utils

type Attr []string

func GetAssemblyBomAttrs() Attr {
	attrs := Attr{
		"No",
		"Parent_Item_No",
		"No",
		"Type",
		"Quantity_per",
		"Unit_of_Measure_Code",
	}
	return attrs
}

func GetInventoryAttrs() Attr {
	attrs := Attr{
		"No",
		"Description",
		"Receipt_Date",
		"Quantity_Available",
	}
	return attrs
}

func GetCustomerCardAttrs() Attr {
	attrs := Attr{
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

func GetItemCardAttrs() Attr {
	attrs := Attr{
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

func GetSalesOrderAttrs() Attr {
	attrs := Attr{
		"No",
		"Sell_to_Customer_No",
		"Sell_to_Customer_Name",
		"Sell_to_Address",
		"Sell_to_Post_Code",
		"Sell_to_City",
		"Sell_to_Contact",
		"No_of_Archived_Versions",
		"Customer_No_Web",
		"Posting_Date",
		"Order_Date",
		"Document_Date",
		"Requested_Delivery_Date",
		"Promised_Delivery_Date",
		"External_Document_No",
		"Salesperson_Code",
		"Assigned_User_ID",
		"Job_Queue_Status",
		"Status",
		"Whs_Shipment_Lines_Exists",
		"Bill_to_Name",
		"Bill_to_Address",
		"Bill_to_Post_Code",
		"Bill_to_City",
		"Shortcut_Dimension_2_Code",
		"Due_Date",
		"Payment_Discount_Percent",
		"Pmt_Discount_Date",
		"Payment_Method_Code",
		"Prices_Including_VAT",
		"VAT_Bus_Posting_Group",
		"Ship_to_Name",
		"Ship_to_Address",
		"Ship_to_Address_2",
		"Ship_to_Post_Code",
		"Ship_to_City",
		"Ship_to_Country_Region_Code",
		"Location_Code",
		"Late_Order_Shipping",
		"Shipment_Date",
		"Shipping_Advice",
		"Currency_Code",
		"EU_3_Party_Trade",
		"Prepayment_Percent",
		"Compress_Prepayment",
		"Prepayment_Due_Date",
		"Prepmt_Payment_Discount_Percent",
		"Prepmt_Pmt_Discount_Date",
	}
	return attrs
}

func GetSalesInvoiceAttrs() Attr {
	attrs := Attr{
		"No",
		"Sell_to_Customer_No",
		"Sell_to_Contact_No",
		"Sell_to_Customer_Name",
		"Sell_to_Address",
		"Sell_to_Address_2",
		"Sell_to_Post_Code",
		"Sell_to_City",
		"Sell_to_Contact",
		"PEB_Note_of_Goods",
		"Posting_Date",
		"Document_Date",
		"Salesperson_Code",
		"Responsibility_Center",
		"Bill_to_Contact_No",
		"Bill_to_Name",
		"Bill_to_Address",
		"Bill_to_Address_2",
		"Bill_to_Post_Code",
		"Bill_to_City",
		"Bill_to_Contact",
		"Shortcut_Dimension_1_Code",
		"Shortcut_Dimension_2_Code",
		"Ship_to_Code",
		"Ship_to_Name",
		"Ship_to_Address",
		"Ship_to_Address_2",
		"Ship_to_Post_Code",
		"Ship_to_City",
		"Ship_to_Contact",
		"Location_Code",
		"Shipment_Method_Code",
		"Shipping_Agent_Code",
		"Package_Tracking_No",
		"Shipment_Date",
		"Your_Reference",
		"Incoming_Document_Entry_No",
		"External_Document_No",
		"Campaign_No",
		"Assigned_User_ID",
		"Job_Queue_Status",
		"Status",
		"Payment_Terms_Code",
		"Due_Date",
		"Payment_Discount_Percent",
		"Pmt_Discount_Date",
		"Payment_Method_Code",
		"Direct_Debit_Mandate_ID",
		"Prices_Including_VAT",
		"VAT_Bus_Posting_Group",
		"Currency_Code",
		"EU_3_Party_Trade",
		"Transaction_Type",
		"Transaction_Specification",
		"Transport_Method",
		"Exit_Point",
		"Area",
	}
	return attrs
}

func GetPostShipAttrs() Attr {
	attrs := Attr{
		"No",
		"Sell_to_Customer_No",
		"Sell_to_Contact_No",
		"Sell_to_Customer_Name",
		"Sell_to_Address",
		"Sell_to_Address_2",
		"Sell_to_Post_Code",
		"Sell_to_City",
		"Sell_to_Contact",
		"PEB_Note_of_Goods",
		"No_Printed",
		"Posting_Date",
		"Document_Date",
		"Requested_Delivery_Date",
		"Promised_Delivery_Date",
		"Quote_No",
		"Order_No",
		"External_Document_No",
		"Salesperson_Code",
		"Responsibility_Center",
		"Bill_to_Customer_No",
		"Bill_to_Contact_No",
		"Bill_to_Name",
		"Bill_to_Address",
		"Bill_to_Address_2",
		"Bill_to_Post_Code",
		"Bill_to_City",
		"Bill_to_Contact",
		"Shortcut_Dimension_1_Code",
		"Shortcut_Dimension_2_Code",
		"Ship_to_Code",
		"Ship_to_Name",
		"Ship_to_Address",
		"Ship_to_Address_2",
		"Ship_to_Post_Code",
		"Ship_to_City",
		"Ship_to_Country_Region_Code",
		"Ship_to_Contact",
		"Location_Code",
		"Outbound_Whse_Handling_Time",
		"Shipping_Time",
		"Shipment_Method_Code",
		"Shipping_Agent_Code",
		"Shipping_Agent_Service_Code",
		"Package_Tracking_No",
		"Shipment_Date",
	}
	return attrs
}

func GetSalesLineAttrs() Attr {
	attrs := Attr{
		"No",
		"Document_Type",
		"Document_No",
		"Line_No",
		"Type",
		"Description",
		"Reserve",
		"Quantity",
		"Reserved_Quantity",
		"Unit_of_Measure_Code",
		"Unit_Price",
		"Line_Amount",
		"Line_Discount_Percent",
		"Line_Discount_Amount",
		"Prepayment_Percent",
		"Prepmt_Line_Amount",
		"Qty_to_Ship",
		"Quantity_Shipped",
		"Qty_to_Invoice",
		"Quantity_Invoiced",
	}
	return attrs
}
