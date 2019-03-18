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

	doc := http.FileServer(http.Dir("./doc/"))
	router.PathPrefix(path + "/doc/").Handler(http.StripPrefix(path+"/doc/", doc))

	codeDoc := http.FileServer(http.Dir("./codeDoc/"))
	router.PathPrefix(path + "/godoc/").Handler(http.StripPrefix(path+"/godoc/", codeDoc))

	router.HandleFunc(path+"/{company:[a-zA-Z]+}", handler)

	var srv http.Server

	srv.Addr = config.Host
	srv.Handler = router

	http2.ConfigureServer(&srv, nil)

	if config.DEV {
		fmt.Println("ENV = DEV")
		fmt.Println("Server started at https://localhost:6789/graphql/v0.1.0/test")
		log.Fatal(srv.ListenAndServeTLS(config.SSL_Cert, config.SSL_Key))
	} else {
		fmt.Println("Server started at http://localhost:6789/graphql/v0.1.0/test")
		log.Fatal(srv.ListenAndServe())
	}

}
