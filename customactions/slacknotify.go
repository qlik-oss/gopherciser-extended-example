package customactions

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/pkg/errors"
	"github.com/qlik-oss/gopherciser/action"
	"github.com/qlik-oss/gopherciser/connection"
	"github.com/qlik-oss/gopherciser/session"
)

type (
	// SlackNotifySettings is an empty dummy action to serve as a template
	SlackNotifySettings struct {
		WebHook string                 `json:"webhook"`
		Msg     session.SyncedTemplate `json:"msg"`
	}

	// Message slack message structure
	Message struct {
		Text string `json:"text"`
	}

	// LocalData used to feed custom data to variable replacer use in go template as {{.Local.warnings}}
	LocalData struct {
		Warnings uint64
	}
)

// Validate SlackNotifySettings action (Implements ActionSettings interface)
func (settings SlackNotifySettings) Validate() error {
	if settings.WebHook == "" {
		return errors.New("no slack webhook defined")
	}
	if settings.Msg.String() == "" {
		return errors.New("no slack message to send defined")
	}
	return nil
}

// Execute SlackNotifySettings action (Implements ActionSettings interface)
func (settings SlackNotifySettings) Execute(sessionState *session.State, actionState *action.State, connection *connection.ConnectionSettings, label string, reset func()) {
	// setup default rest client if not setup
	if sessionState.Rest.Client == nil {
		client, err := session.DefaultClient(connection.Allowuntrusted, sessionState)
		if err != nil {
			actionState.AddErrors(errors.WithStack(err))
			return
		}
		sessionState.Rest.SetClient(client)
	}

	// Add data for custom variables
	localData := LocalData{Warnings: sessionState.EW.TotWarnings()}

	// Execute go template to add variables to "Msg"
	txt, err := sessionState.ReplaceSessionVariablesWithLocalData(&settings.Msg, localData)
	if err != nil {
		actionState.AddErrors(err)
		return
	}
	msg := Message{Text: txt}

	content, err := jsoniter.Marshal(msg)
	if err != nil {
		actionState.AddErrors(errors.Wrap(err, "failed to marshal slack message"))
		return
	}

	options := session.DefaultReqOptions()
	options.NoVirtualProxy = true // if connection is normally using virtual proxy, disable it for this request
	sessionState.Rest.PostAsync(settings.WebHook, actionState, sessionState.LogEntry, content, options)

	sessionState.Wait(actionState) // Await all async requests and reactions, errors will be added to actionState and reported
}
