package main

import (
	"github.com/qlik-oss/gopherciser/generatedocs/pkg/extenddocs"

	// import for register actions side effect
	_ "github.com/qlik-oss/gopherciser-extended-example/customactions"
)

func main() {
	extenddocs.ExtendOSSDocs()
}
