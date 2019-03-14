package inventory

import (
	"fmt"
	"github.com/vjeantet/jodaTime"
	"strings"
	"time"
)

func getDate() string {
	date := jodaTime.Format("YYYY-MM-dd", time.Now())
	return date
}

func getPayload(args map[string]interface{}) []byte {

	item := "<tns:itemFilter xsi:nil=\"true\"/>"
	category := "<tns:categoryFilter xsi:nil=\"true\"/>"
	returnValues := "<tns:retValues xsi:nil=\"true\"/>"
	startDate := "<tns:startDate>" + getDate() + "</tns:startDate>"

	if args != nil {
		if args["key"].(string) == "ItemNo" {
			item = fmt.Sprintf("<tns:itemFilter>%v</tns:itemFilter>", args["value"])
		}
	}

	payload := []byte(strings.TrimSpace(`
	<env:Envelope xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xmlns:tns="urn:microsoft-dynamics-schemas/codeunit/GetInventory" 
	xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" 
	xmlns:ins0="urn:microsoft-dynamics-nav/xmlports/x50000">
		<env:Body>
			<tns:GetInventory>` +
		item +
		category +
		returnValues +
		startDate +
		`</tns:GetInventory>
		</env:Body>
	</env:Envelope>`))

	return payload
}
