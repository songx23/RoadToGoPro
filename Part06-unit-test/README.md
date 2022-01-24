# Road to Go Pro — Unit Test

Please read the tutorial on [Medium](https://medium.com/digio-australia/road-to-go-pro-unit-test-69591a553412).

## Go test 

To run all the unit tests:

```bash
go test ./...
```

To run unit tests for a specific package:

```bash
go test ./pkg/calculator
```

To run all the unit tests and output test coverages in the console:

```bash
go test ./... coverprofile=coverage.out
go tool cover -func=coverage.out
```

To run all the unit tests and output test coverage details in HTML format:

```bash
go test ./... coverprofile=coverage.out
go tool cover -html=coverage.out
```

To run a specific sub test

```bash
# example
go test ./pkg/calculator -run=TestTableDrivenTestDivide/Divided_by_zero -v
```
