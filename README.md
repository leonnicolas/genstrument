# Genstrument

Generate code to instrument an interface.

## Supported Interfaces

### Binary

An interface with methods that return an error as the last return value.
Depending if the error is nil, the corresponding counter vector will be increased.

### Handler

An interface with methods that have no return values, the first argument is an `http.ResponseWriter` and the second argument is an `*http.Request`.

## Install

Install with
```shell
go install github.com/leonnicolas/genstrument@latest
```

or add `github.com/leonnicolas/genstrument` to your `tools.go` file.

## Usage

Run
```shell
genstrument --file-path <file-path-to-interface> -p <interface-name> --metric-help <help text> --metric-name <metric-name> -o -
```
to print the generated code to stdout.

Or add something like
```shell
//go:generate go run github.com/leonnicolas/genstrument --file-path <file-name> -p <interface-name> --metric-help <help text> --metric-name <metric-name> -o gen.go
```
into the go file, that contains the interface you want to instrument.
