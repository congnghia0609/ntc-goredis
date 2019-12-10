# ntc-goredis
ntc-goredis is module redis golang client.  

## Install
```bash
go get -u github.com/congnghia0609/ntc-goredis
```

## Install dependencies
```bash
make deps
```

## Build
```bash
make build
```

## Clean file build
```bash
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
or
```bash
go install
ntc-goredis
```
