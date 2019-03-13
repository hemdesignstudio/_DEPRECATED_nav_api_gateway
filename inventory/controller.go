package inventory

import (
	"github.com/hem-nav-gateway/request"
	"strings"
)

const (
	SOAP_ACTION  = "urn:GetInventory"
	ENDPOINT_URI = "https://wsmyway182.navmoln.se:18201/MyWay182Services/WS/Hem%20TEST/Codeunit/GetInventory"
)

var date = "2019-03-08"

var soapPayload = []byte(strings.TrimSpace(`
	<env:Envelope xmlns:xsd="http://www.w3.org/2001/XMLSchema" 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xmlns:tns="urn:microsoft-dynamics-schemas/codeunit/GetInventory" 
	xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" 
	xmlns:ins0="urn:microsoft-dynamics-nav/xmlports/x50000">
		<env:Body>
			<tns:GetInventory>
				<tns:itemFilter xsi:nil="true"/>
				<tns:categoryFilter xsi:nil="true"/>
				<tns:retValues xsi:nil="true"/>
				<tns:startDate>` + date + `</tns:startDate>
			</tns:GetInventory>
		</env:Body>
	</env:Envelope>
`,
))

func GetAll() {
	req := request.Request{
		Payload:    soapPayload,
		SoapAction: SOAP_ACTION,
		Uri:        ENDPOINT_URI,
		Response:   Response{},
	}

	req.SoapRequest()

}
