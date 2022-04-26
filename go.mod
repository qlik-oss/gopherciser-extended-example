module github.com/qlik-oss/gopherciser-extended-example

go 1.15

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1
	github.com/json-iterator/go v1.1.12
	github.com/pkg/errors v0.9.1
	github.com/qlik-oss/gopherciser v0.15.2
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/eventials/go-tus => github.com/andrewmostello/go-tus v0.0.0-20200314041820-904a9904af9a
