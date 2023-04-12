package oapicodegenclient

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config oapi-codegen.yaml petstore.yaml

//go:generate go run github.com/leonnicolas/genstrument --file-path oapi.go -p  ClientWithResponsesInterface --metric-help "help to the metric" --metric-name metric_name_total -o gen.go --mode oapi-codegen-client
