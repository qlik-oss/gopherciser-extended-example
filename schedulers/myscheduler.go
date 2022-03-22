package schedulers

import (
	"context"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/qlik-oss/gopherciser/connection"
	"github.com/qlik-oss/gopherciser/helpers"
	"github.com/qlik-oss/gopherciser/logger"
	"github.com/qlik-oss/gopherciser/randomizer"
	"github.com/qlik-oss/gopherciser/scenario"
	"github.com/qlik-oss/gopherciser/scheduler"
	"github.com/qlik-oss/gopherciser/statistics"
	"github.com/qlik-oss/gopherciser/users"
)

type (
	MySchedulerSettings struct {
		ConcurrentUsers int `json:"concurrentUsers" displayname:"Concurrent users" doc-key:"config.scheduler.settings.concurrentusers"`
	}

	MyScheduler struct {
		scheduler.Scheduler
		Settings MySchedulerSettings `json:"settings" doc-key:"config.myscheduler.settings"`
	}
)

const MySchedulerName = "myscheduler"

func init() {
	_ = scheduler.RegisterScheduler(MySchedulerName, MyScheduler{})
}

// Validate scheduler
func (sched MyScheduler) Validate() ([]string, error) {
	// validate inherited settings
	if err := sched.Scheduler.Validate(); err != nil {
		return nil, err
	}

	// validate MyScheduler settings
	if sched.Settings.ConcurrentUsers < 1 {
		return nil, errors.Errorf("ConcurrentUsers<%d> most be > 0", sched.Settings.ConcurrentUsers)
	}

	return nil, nil
}

// Execute scheduler
func (sched MyScheduler) Execute(
	ctx context.Context,
	log *logger.Log,
	timeout time.Duration,
	scenario []scenario.Action,
	outputsDir string, // outputsDir
	users users.UserGenerator,
	connectionSettings *connection.ConnectionSettings,
	counters *statistics.ExecutionCounters,
) error {
	if counters == nil {
		return errors.New("execution counters are nil")
	}
	sched.Scheduler.ConnectionSettings = connectionSettings

	// Create a seeded randomizer for reproducbile results
	seededRandomizer := randomizer.NewSeededRandomizer(12345)

	var mErr *multierror.Error
	for thread := counters.Threads.Inc(); thread <= uint64(sched.Settings.ConcurrentUsers); thread = counters.Threads.Inc() {
		if thread != 1 { // No wait time for first user
			// wait random time between 10-30s inbetween introducing users
			rampup, err := seededRandomizer.RandDuration(10*time.Second, 30*time.Second)
			if err != nil {
				return err
			}
			helpers.WaitFor(ctx, rampup)
		}
		if helpers.IsContextTriggered(ctx) {
			return nil
		}
		user := users.GetNext(counters)
		err := sched.Scheduler.StartNewUser(ctx, timeout, log, scenario, thread, outputsDir, user, 1, false, counters)
		if err != nil {
			mErr = multierror.Append(mErr, err)
		}
	}

	return errors.WithStack(helpers.FlattenMultiError(mErr))
}

// RequireScenario reports if the scheduler requires a scenario or not for scheduler validation.
func (sched MyScheduler) RequireScenario() bool {
	return true
}

// PopulateHookData populate map with data which can be used by go template in hooks
func (sched MyScheduler) PopulateHookData(data map[string]interface{}) {
	data["ConcurrentUsers"] = sched.Settings.ConcurrentUsers
}
