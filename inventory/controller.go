package inventory

import (
	"github.com/graphql-go/graphql"
	"strings"
)

func getDate() string {
	date := "2019-03-08"
	return date

}

func getPayload() []byte {

	item := "<tns:itemFilter>10005</tns:itemFilter>"
	category := "<tns:categoryFilter xsi:nil=\"true\"/>"
	returnValues := "<tns:retValues xsi:nil=\"true\"/>"
	startDate := "<tns:startDate>" + getDate() + "</tns:startDate>"

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

var _type = createType()
var _args = createArgs()

type Request struct {
	Company string
}

func (*Request) CreateType() *graphql.Object {
	return _type
}

func (*Request) CreateArgs() map[string]*graphql.ArgumentConfig {
	return _args
}

func (r *Request) GetCompany() string {
	return r.Company
}

func (r *Request) SetArgs(args map[string]interface{}) {

}

func (r *Request) SetFields(fields []string) {
}

// GetAll retrieves a List of all ItemCards available Microsoft Navision .
// Function takes a list of fields to be returned by Microsoft Navision.
func (r *Request) GetAll() (interface{}, error) {
	return soapRequest(getPayload())
}

func (r *Request) Filter() (interface{}, error) {
	return nil, nil
}

func (r *Request) Create() (interface{}, error) {
	return nil, nil
}

func (r *Request) Update() (interface{}, error) {
	return nil, nil
}
