package schedulers

import (
	"context"
	"time"

	"github.com/qlik-oss/gopherciser/connection"
	"github.com/qlik-oss/gopherciser/logger"
	"github.com/qlik-oss/gopherciser/scenario"
	"github.com/qlik-oss/gopherciser/statistics"
	"github.com/qlik-oss/gopherciser/users"
)

type MyScheduler struct{}

func (sched MyScheduler) Validate() ([]string, error) {
	return nil, nil
}
func (sched MyScheduler) Execute(
	context.Context,
	*logger.Log,
	time.Duration,
	[]scenario.Action,
	string, // outputsDir
	users.UserGenerator,
	*connection.ConnectionSettings,
	*statistics.ExecutionCounters,
) error {
	return nil
}
func (sched MyScheduler) RequireScenario() bool {
	return false
}

// PopulateHookData populate map with data which can be used by go template in hooks
func (sched MyScheduler) PopulateHookData(data map[string]interface{}) {

}
