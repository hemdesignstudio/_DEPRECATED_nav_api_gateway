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
	viper.AddConfigPath("/etc/nav/")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".") // optionally look for config in the working directory

	viper.SetConfigType("json")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func NAVCredentials() (string, string) {
	username := viper.GetString("USER")
	password := viper.GetString("PWD")
	return username, password
}

var (
	BaseUrl                = viper.GetString("BASEURL")
	SoapBaseUrl            = viper.GetString("SOAPBASEURL")
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
	TestCompany            = viper.GetString("TEST_STORE")
	EUCompany              = viper.GetString("EU_STORE")
	USCompany              = viper.GetString("US_STORE")
	SSL_Cert               = "./cert/localhost.crt"
	SSL_Key                = "./cert/localhost.key"
)

var CA_CERT = "./cert/rootCA.pem"

var Companies = map[string]string{
	"test": TestCompany,
	"us":   USCompany,
	"eu":   EUCompany,
}
