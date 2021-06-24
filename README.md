# Cronlog

Cronlog is a command line interface for cron job logging, taking a cron job as an argument and logging out the job in a human-readable format.

## Task

Time box: 3 hours

Write a cli which parses a cron string and expands each field to show the times at which it will run.

## Prerequisite

You will need [Docker](https://www.docker.com/) or [Go](https://golang.org/) installed on your machine

## Usage

### Docker

For Docker, run the following command
```shell
make serve
```

This will open up a shell to run any command for the cli

```shell
cronlog "*/15 0 1,5 * 1-5 /bin"

# output:
#minute        15 30 45 60
#hour          0
#day of month  1 5
#month         1 2 3 4 5 6 7 8 9 10 11 12
#day of week   1 2 3 4 5
#command       /bin
```

For Go, run the following command

```shell
go build
```

This will build the cli in your current directory, where you can run the cli

```shell
./cronlog "*/15 0 1,5 * 1-5 /bin"
```

## Test and coverage

Run the following commands for testing and coverage reports

```shell
make test

make cover
```