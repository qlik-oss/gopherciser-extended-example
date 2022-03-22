package main

import (
	"github.com/qlik-oss/gopherciser/cmd"

	// Register custom actions
	_ "github.com/qlik-oss/gopherciser-extended-example/customactions"

	// Register custom schedulers
	_ "github.com/qlik-oss/gopherciser-extended-example/schedulers"
)

//go:generate go run ./generatedocs/cmd/compiledocs
//go:generate go run ./generatedocs/cmd/generatemarkdown --output ./generatedocs/generated/settingup.md
func main() {
	cmd.Execute()
}
