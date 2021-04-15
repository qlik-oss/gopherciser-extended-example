# Example of how to extend [gopherciser](https://github.com/qlik-oss/gopherciser) with custom actions

The repo showcases how to extend [gopherciser](https://github.com/qlik-oss/gopherciser) with custom actions to be used during user simulations towards a Qlik Sense environment. All functionality of the the original [gopherciser](https://github.com/qlik-oss/gopherciser) will be included by importing it.

## Included actions

This repo includes two actions registered on `init`.

### `dummy` action

This action is an "empty" action which can be used as a template action to be copied and used to implement new actions.

### `slacknotify` action

This action is an example action on custom action implementation. This example action should be deleted when using the repo as a template repo. The action send a message to slack channel webhook after each *successful* user iteration, accounting for which session and thread was finished as well as how many warnings the iteration had.

An example script using this action can be found [here](./examples/notifydemo.json). After updating it with correct `server` and `webhook` it can be executed as follows:
```bash
# Build the tool
go build -o myowngopherciser .
# Execute script including custom actions
./myowngopherciser execute -c ./examples/notifydemo.json
```

## Generate documentation

This project is set up to extend the documentation of gopherciser. To
(re)generate documentation run `go generate` in the root of this repository.
This will generate the following files:
- generatedocs/generated/documentation.go
- generatedocs/generated/settingup.md

The documentation of documentation generation can be found
[here](https://github.com/qlik-oss/gopherciser/blob/master/generatedocs/README.md#how-extending-existing-documetation).
