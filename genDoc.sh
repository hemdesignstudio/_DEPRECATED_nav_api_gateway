#!/bin/bash
go run main.go
graphdoc -e http://localhost:6789/graphql/v0.1.0/test -o ./doc
graphql-markdown http://localhost:6789/graphql/v0.1.0/test > SCHEMA.md

godoc -http=localhost:6060
wget -r -np -N -E -p -k http://localhost:6060/pkg/github.com/hem-nav-gateway/