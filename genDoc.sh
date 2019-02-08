#!/bin/bash
go run main.go
graphdoc -e http://localhost:6789/graphql/v0.1.0/company=test -o ./doc
graphql-markdown http://localhost:6789/graphql/v0.1.0/company=test > schema.md