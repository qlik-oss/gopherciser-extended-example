# Adding a custom scheduler

## IScheduler interface

The first thing a custom scheduler needs to implement is the `IScheduler`interface

```golang
	// IScheduler interface of scheduler
	IScheduler interface {
		Validate() ([]string, error)
		Execute(
			context.Context,
			*logger.Log,
			time.Duration,
			[]scenario.Action,
			string, // outputsDir
			users.UserGenerator,
			*connection.ConnectionSettings,
			*statistics.ExecutionCounters,
		) error
		RequireScenario() bool
		// PopulateHookData populate map with data which can be used by go template in hooks
		PopulateHookData(data map[string]interface{})
	}
```

## Register a custom scheduler

Once there's an implementation of the Ischeduler interface it can be registered using the `scheduler.RegisterScheduler` or `RegisterSchedulerOverride` methods.

### Examples registering a scheduler

The scheduler also needs to be registered using the `RegisterScheduler` method.

```golang
func init() {
	_ = scheduler.RegisterScheduler(MySchedulerName, MyScheduler{})
}
```

Overriding an existing scheduler with the same name can be done with the `RegisterSchedulerOverride` method.

```golang
func init() {
	_ = scheduler.RegisterSchedulerOverride(MySchedulerName, MyScheduler{})
}
```

## Example scheduler

This package includes [an example scheduler](myscheduler.go) randomizing the period inbetween each user get introduced. An [example](examples/qlikcoresheetchanger.json) script using this can be found in the examples folder.
