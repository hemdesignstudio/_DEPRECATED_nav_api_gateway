package inventory

type Response struct {
	Body struct {
		GetInventoryResult struct {
			RetValues struct {
				InventoryLine InventoryLine `xml:"retValues>InventoryLine"`
			} `xml:"GetInventory_Result>retValues"`
		} `xml:"Body>GetInventory_Result"`
	}
}

type InventoryLine struct {
	ItemNo        string `xml:"ItemNo"`
	Description   string `xml:"Description"`
	ReceiptDate   string `xml:"ReceiptDate"`
	AvailableDate string `xml:"AvailableDate"`
}
