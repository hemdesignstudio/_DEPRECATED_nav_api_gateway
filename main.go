package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/roothandler"
	"golang.org/x/net/http2"
	"log"
	"net/http"
)

func main() {
	path := config.Endpoint + config.Version
	router := mux.NewRouter().StrictSlash(true)
	handler := roothandler.RootEndpoint

	router.HandleFunc(path+"/{company:[a-zA-Z]+}", handler)

	var srv http.Server

	srv.Addr = config.Host
	srv.Handler = router

	err := http2.ConfigureServer(&srv, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	if config.DEV {
		fmt.Println("ENV = DEV")
		fmt.Println("Server started at https://localhost:6789/graphql/v0.1.0/test")
		log.Fatal(srv.ListenAndServeTLS(config.SSL_Cert, config.SSL_Key))
	} else {
		fmt.Println("Server started at http://localhost:6789/graphql/v0.1.0/test")
		log.Fatal(srv.ListenAndServe())
	}

}
