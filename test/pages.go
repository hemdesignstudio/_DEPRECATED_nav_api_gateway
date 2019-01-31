package test

const assemblyBomPage = "AssemblyBom"
const customerCardPage = "CustomerCard"

type Page struct {
	assemblybom  string
	customerCard string
}

var pages = Page{
	assemblybom:  assemblyBomPage,
	customerCard: customerCardPage,
}
