# Server Golang

## Structure overview
- cmd: contains go file to run project 
- config: environment config
- dal: data access layer
    - model: includes entity map with your database (one to many, etc ...)
    - query: file code gen using to query helping CRUD friendly and safely
    - psql.go: set up connect DB (In this project, i'm using psql)
- db: contains migration files which are used for init database for project.
- grpc: 
  - repository: directly interact with database.
  - server: set up for grpc server
  - service: set up function for grpc service
- utils: helper to generate as well as test query

## Install
### Set Global GOPATH
```shell
    export $GOPATH=$HOME/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
### Get library

```shell
sudo apt install protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/spf13/cobra-cli@latest
```
#### Download pre-built migrate-cli binary (Windows, MacOS, or Linux)
```shell
 https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
```

### Init Database
- Create database with name grpc (psql)

## Run
```shell
  go install
  make migrate-up
  go run main.go (or server-corba) 
```
    
## References
1. https://github.com/go-gorm/gen
2. https://github.com/spf13/cobra
3. https://pkg.go.dev/go.uber.org/fx
4. https://pkg.go.dev/github.com/spf13/cobra
