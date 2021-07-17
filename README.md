# ntc-goredis
ntc-goredis is module redis golang client.  

## Install
```bash
go get -u github.com/congnghia0609/ntc-goredis
```

## Build from source
```bash
# Install dependencies
#make deps
go mod download

# update go.mod file
go mod tidy

# Build
export GO111MODULE=on
make build

# Clean file build
make clean
```

## Run with environment: development | test | staging | production
### development
```bash
make run
```
### test
```bash
make run-test
```
### staging
```bash
make run-stag
```
### production
```bash
make run-prod
```

## Run test file
```bash
cd ntc-goredis

go run main.go
```


## License
This code is under the [Apache License v2](https://www.apache.org/licenses/LICENSE-2.0).  
