# Gin Template App

## Getting this Repo on your machine
```
mkdir -p $GOPATH/src/github.com/yiziz/ && cd $_ && git clone https://github.com/yiziz/gin-template
```

## Getting dependencies
Run `glide install` inside the project folder

## Adding dependencies
Run `glide get github.com/username/reponame` inside the project folder

## Environment Variables
Please add these environment variables to your `.bashrc`
```
# this should point to your app path
export GIN_TEMPLATE_PATH=$GOPATH/src/github.com/yiziz/gin-template
```

## Running migrations
Run `go run main.go db:migrate` in your project folder

## Running the server
Run `go run main.go` in your project folder

## Testing

### Via [GoConvey](https://github.com/smartystreets/goconvey)
Execute `go get github.com/smartystreets/goconvey` to get the `goconvey` tool.  Run `goconvey -host '0.0.0.0.'` in the directory that holds tests you want to run.  All tests in child folders will also be run.  You should be able to view the tests in a nice UI at `http://127.0.0.1:8080/`.

GOTCHAS: [Execution Order](https://github.com/smartystreets/goconvey/wiki/Execution-order).

### Via `go test`
You can also run tests by running `go test` in the directory that holds tests you want to run.
