package main

import (
	"github.com/qlik-oss/gopherciser/cmd"

	// make sure to register custom actions
	_ "github.com/qlik-oss/gopherciser-extended-example/customactions"
)

//go:generate go run ./generatedocs/cmd/compiledocs
//go:generate go run ./generatedocs/cmd/generatemarkdown --output ./generatedocs/generated/settingup.md
func main() {
	cmd.Execute()
}
