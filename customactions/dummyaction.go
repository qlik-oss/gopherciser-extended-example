package customactions

import (
	"github.com/qlik-oss/gopherciser/action"
	"github.com/qlik-oss/gopherciser/connection"
	"github.com/qlik-oss/gopherciser/session"
)

type (
	// DummySettings is an empty dummy action to serve as a template
	DummySettings struct {
		ID string `json:"id"`
	}
)

// Validate DummySettings action (Implements ActionSettings interface)
func (settings DummySettings) Validate() ([]string, error) {
	// Add validation checks for the action here, Validate will be run before starting a simulation or
	// when the script validate command is issued
	return nil, nil
}

// Execute DummySettings action (Implements ActionSettings interface)
func (settings DummySettings) Execute(sessionState *session.State, actionState *action.State, connection *connection.ConnectionSettings, label string, reset func()) {

	// *** Add actual user simulation here ***

	sessionState.Wait(actionState) // Await all async requests, e.g. those triggered on changed objects
}

// IsContainerAction Implements the ContainerAction interface. This marks the action as a
// container action containing other actions. A container action will not log result as a normal action,
// instead result will be logged as level=info, infotype: containeractionend
// Returns if action is to be considered a container action.
// ContainerAction can't be used in conjunction with StartActionOverrider interface
// see e.g. "iterated" action
// func (settings DummySettings) IsContainerAction() {}

// AppStructureAction Implements AppStructureAction interface. It returns if this action should be included
// when doing an "get app structure" from script, IsAppAction tells the scenario
// to insert a "getappstructure" action after that action using data from
// sessionState.CurrentApp. A list of Sub action to be evaluated can also be included
// AppStructureAction returns if this action should be included when getting app structure
// and any additional sub actions which should also be included
//func (settings *DummySettings) AppStructureAction() (*scenario.AppStructureInfo, []scenario.Action) {
//	return &scenario.AppStructureInfo{
//		IsAppAction: false,
//		Include:     true,
//	}, nil
//}

// AffectsAppObjectsAction implements AffectsAppObjectsAction interface
// Should be implemented by all actions that affect the availability of selectable objects for app structure consumption.
// App structure of the current app is passed as an argument. The return is
// * added *config.AppStructurePopulatedObjects - objects to be added to the selectable list by this action
// * removed []string - ids of objects that are removed (including any children) by this action
// * clearObjects bool - clears all objects except bookmarks and sheets
//func (settings DummySettings) AffectsAppObjectsAction(structure appstructure.AppStructure) ([]*appstructure.AppStructurePopulatedObjects, []string, bool) {
//	selectables, err := structure.GetSelectables(settings.ID)
//	if err != nil {
//		return nil, nil, false
//	}
//	newObjs := appstructure.AppStructurePopulatedObjects{
//		Parent:    settings.ID,
//		Objects:   make([]appstructure.AppStructureObject, 0),
//		Bookmarks: nil,
//	}
//	newObjs.Objects = append(newObjs.Objects, selectables...)
//	return []*appstructure.AppStructurePopulatedObjects{&newObjs}, nil, true
//}
