# gotrain
Fundamentals training in Golang

Main goal af this repo: Fundamentals training in Golang.
--------------------------------------------------------

## Daily devops tasks automation

1. Parse csv dataset into predefined structs/maps:

```bash
go run parseCsv.go 
```

## Unit testing

0. Go to `unit-test-example`

1. Run unit testing:

```bash
go test -v
```

2. Check testing coverage:

```bash
go test -v -coverprofile=test-coverage.txt
```

3. Prettify coverage profile in html:

```bash
go tool cover -html=test-coverage.txt -o test-coverage.html
```

## Benchmark testing

0. Go to `sorting-example/merge-sort`

1. Run unit testing:

```bash
go test -bench=.
```

