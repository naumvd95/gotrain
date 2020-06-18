# gotrain
Fundamentals training in Golang

Main goal af this repo: Fundamentals training in Golang.
--------------------------------------------------------

## Daily devops tasks automation

1. Parse csv dataset into predefined structs/maps:

```bash
go run parseCsv.go 
```

2. Parse nginx access log from K8S pod as unformatted csv:

```bash
cd parsing-plain-text-example && go run parseNginx.go
```

topics:

- get log as file from pod logs stdout
- get rid of infra nginx logs and provide only clear access log using `os` for file ops
and `bufio` Scanner to get rid of non-access log lines (using IP regex match)
- parse as csv with whitespace delimeter
- map it into struct with time.Time and int formats for response codes/timelogs

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

## Sorting

Go to `sorting-example/merge-sort`, topics:

- merge sort for customized struct objects
- single thread sorting
- multithread sorting (goroutine and waitgroups)
- threads management (semaphore pattern as buffered channel for driving amount of goroutines and switch to single thread in case of limit exceeded)

## Map Reduce operations

0. Go to `map-merge-reduce-example`
1. Run app:

```bash
go run mmr.go
```

topics:

- get CSV data re: Covid cases from open sources
- Define desired date to monitor cases in all countries (MAP training)
- Define amount of critical deaths border (REDUCE training)
- Sort values using merge sort alghoritm
- MAP + multithread = creating struct objects from csv plain text
- REDUCE + multithread = filter them by critical deaths border
- Sort + multithread = sort by amount of death
- Pretty print
-
