package customactions

import (
	"fmt"

	"github.com/qlik-oss/gopherciser/scenario"
)

const (
	DummyAction = "dummy"
)

func init() {
	err := scenario.RegisterActions(
		map[string]scenario.ActionSettings{
			// Remove DummyAction and add list of custom actions here with corresponding Settings
			DummyAction: DummySettings{},
		})
	if err != nil {
		panic(fmt.Sprintf("failed to register custom actions:\n %+v", err))
	}
}
