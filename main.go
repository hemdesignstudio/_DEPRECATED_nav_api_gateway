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
	router := mux.NewRouter()
	handler := roothandler.RootEndpoint
	router.HandleFunc(config.Endpoint+config.Version+"/{company:[a-zA-Z]+}", handler)

	fmt.Println("Server started at http://localhost:6789/graphql/v0.1.0/test")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(config.Host, nil))

}
