package importedclient

//go:generate go run ../../../ --file-path ../oapi.go -p  ClientWithResponsesInterface --package importedclient --metric-help "help to the metric" --metric-name metric_name_total -o gen.go --mode oapi-codegen-client
