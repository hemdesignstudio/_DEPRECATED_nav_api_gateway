package main

import (
	"fmt"
	"github.com/nav-api-gateway/config"
	"github.com/nav-api-gateway/roothandler"
	"log"
	"net/http"
)

func main() {
	http.Handle(config.Endpoint+config.Version, roothandler.RootEndpoint())
	fmt.Println("Server started at http://localhost:6789/graphql/v0.1.0")
	log.Fatal(http.ListenAndServe(config.Host, nil))
}
