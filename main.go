package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/roothandler"
	"log"
	"net/http"
)

func main() {
	path := config.Endpoint + config.Version
	router := mux.NewRouter().StrictSlash(true)
	handler := roothandler.RootEndpoint
	fs := http.FileServer(http.Dir("./doc/"))
	router.PathPrefix(path + "/doc/").Handler(http.StripPrefix(path+"/doc/", fs))
	router.HandleFunc(path+"/{company:[a-zA-Z]+}", handler)
	fmt.Println("Server started at http://localhost:6789/graphql/v0.1.0/test")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(config.Host, nil))
}
