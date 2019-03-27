// Copyright 2019 Hem Design Studio. All rights reserved.
// Use of this source code is governed by a
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	DEV = false
)

func init() {
	viper.SetConfigName("credentials") // name of config file (without extension)
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".") // optionally look for config in the working directory

	viper.SetConfigType("json")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func NAVCredentials() (string, string) {
	username := viper.GetString("username")
	password := viper.GetString("password")
	return username, password
}

var (
	BaseUrl                = "https://odatamyway182.navmoln.se:18202/MyWay182Services/ODataV4"
	SoapBaseUrl            = "https://wsmyway182.navmoln.se:18201/MyWay182Services/WS"
	Endpoint               = "/graphql"
	Version                = "/v0.1.0"
	CompanyEndpoint        = "/Company"
	AssemblyBomEndpoint    = "/Assembly_Bom"
	CustomerCardWSEndpoint = "/CustomerCardWS"
	ItemCardEndpoint       = "/ItemCard"
	SalesOrderEndpoint     = "/SalesOrderWS"
	SalesLineEndpoint      = "/SalesLines"
	SalesInvoiceEndpoint   = "/SalesInvoice"
	PostShipEndpoint       = "/WebSalesShipment"
	InventoryEndpoint      = "/GetInventory"
	Host                   = ":6789"
	TestCompany            = "Hem TEST"
	EUCompany              = "Hem FAL Sweden AB"
	USCompany              = "Hem FAL Inc"
	SSL_Cert               = "./cert/localhost.crt"
	SSL_Key                = "./cert/localhost.key"
)

var CA_CERT = "./cert/rootCA.pem"

var Companies = map[string]string{
	"test": TestCompany,
	"us":   USCompany,
	"eu":   EUCompany,
}
