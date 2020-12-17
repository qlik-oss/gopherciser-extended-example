package main

import (
	"github.com/qlik-oss/gopherciser/cmd"

	// make sure to register custom actions
	_ "github.com/qlik-oss/gopherciser-extended-example/customactions"

)

func main() {
	cmd.Execute()
}
