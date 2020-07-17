package customactions

import (
	"fmt"

	"github.com/qlik-oss/gopherciser/scenario"
)

const (
	DummyAction = "dummy"
	SlackNotify = "slacknotify"
)

func init() {
	err := scenario.RegisterActions(
		map[string]scenario.ActionSettings{
			// Remove DummyAction and SlackNotify and add list of custom actions here with corresponding Settings
			DummyAction: DummySettings{}, // "empty" dummy action which can be used as a template
			SlackNotify: SlackNotifySettings{}, // example action with a custom implementation
		})
	if err != nil {
		panic(fmt.Sprintf("failed to register custom actions:\n %+v", err))
	}
}
