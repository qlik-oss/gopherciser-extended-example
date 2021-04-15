package main

import (
	"github.com/qlik-oss/gopherciser-extended-example/generatedocs/generated"
	"github.com/qlik-oss/gopherciser/generatedocs/pkg/genmd"

	// import for register actions side effect
	_ "github.com/qlik-oss/gopherciser-extended-example/customactions"
)

func main() {
	genmd.GenerateMarkdown(&genmd.CompiledDocs{
		Actions: generated.Actions,
		Params:  generated.Params,
		Config:  generated.Config,
		Groups:  generated.Groups,
		Extra:   generated.Extra,
	})
}
