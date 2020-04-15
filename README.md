# Backend Interview

Clean API for Backend Interview.

## Setup

### Quick start

Deploy it in less than 30 seconds!

### Prerequisite
* [GO](https://golang.org/doc/install) `1.14+`
* [MongoDb](https://docs.mongodb.com/manual)

### Run
#### Running Locally

```shell script
$ go get -u github.com/Pantani/replacer
$ cd $GOPATH/src/github.com/Pantani/replacer
$ make start
```

#### Running inside a Docker

```shell script
docker-compose up -d
```

#### Running in the IDE ( GoLand )

1.  Run
2.  Edit configuration
3.  New Go build configuration
4.  Select `directory` as configuration type
5.  Set `api` as program argument and `-i` as Go tools argument 

#### Tools

-   Setup MongoDb

```shell script
brew install mongodb
``` 

#### Make commands

```
- install       Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
- start         Start API in development mode.
- stop          Stop development mode.
- compile       Compile the binary.
- exec          Run given command. e.g; make exec run="go test ./..."
- clean         Clean build files. Runs `go clean` internally.
- test          Run all unit tests.
- fmt           Run `go fmt` for all go files.
- goreleaser    Release the last tag version with GoReleaser.
- govet         Run go vet.
- golint        Run golint.
- docs          Generate swagger docs.
- script-test   Run tests from `test/test.sh`
```

#### Environment Variables

All environment variables for developing are set inside the .env file.

### Docs

Swagger API docs provided at path `/swagger/index.html`

### Unit Tests

To run the unit tests: `make test`
