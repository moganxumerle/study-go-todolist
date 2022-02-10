# Project Name
> CLI Task Manager TODOs List

## Index
* [Information](#information)
* [Tecnologies](#tecnologias)
* [Instalation](#instalation)
* [Environment](#environment)
* [Starting](#starting)
* [Tests](#tests)

## Information
CLI tool that can be used to manage your TODOs in the terminal.

## Tecnologies
* [GoLang](https://golang.org/) - Go compiler
* [Go Mod](https://github.com/golang/mod) - dependency manager

## Instalation

### Clonando o projeto

```bash
$ cd $PROJECT_HOME
$ git clone git@github.com:moganxumerle/study-go-todolist.git
```

```bash
# Dependencies instalation
$ go get
```

```bash
# Removing unwanted dependencies
$ go mod tidy
```

## Environment
Setting environment variables

| Name | Description | Default | Mandatory |
| -- | -- | -- | -- |
| USER | User name | | :white_check_mark: |


## Starting
Building the project
```bash
# run the command below to build the application and make sure nothing is broken
$ go build
```

Executing the project
```bash
$ go run main.go
```

## Tests
```bash
# To run the automated tests run the command below in the terminal inside the application folder
$ go test -v -cover ./...

# To generate the interface showing all files and lines "Covered", "Not Covered" e "Not Tracked":
$ go test ./... -coverprofile fmtcoverage.html fmt
$ go test ./... -coverprofile cover.out
$ go tool cover -html=cover.out
$ go tool cover -html=cover.out -o cover.html
$ open 'cover.html' file
```

> ###### Current developers:
>
> https://github.com/moganxumerle
