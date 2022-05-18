# GAT

![Build status](https://gitlab.com/Ipfaze/gat/badges/main/pipeline.svg)
![coverage](https://gitlab.com/Ipfaze/gat/badges/main/coverage.svg)

Gat is a command line program that provides multiple functionalities such as :
- Printing file's content
- Searching for keyword using regex or exact matching
- Count lines
- Count words

## Usage
First you can check the output of the help command like :

### Linux or MacOS
```shell
gat --help
```
### Windows
```shell
gat.exe --help
```

## Build

The go program is required to build this project, if it's not installed on your machine please download it here : [https://go.dev/dl/](https://go.dev/dl/)

To build the project, simply download this repository and execute :
```sh
go build
```

## Tests

### Unit tests

#### Without coverage report
```shell
go test
```

#### With coverage report
```shell
go test -cover
```

#### With coverage report by function
```shell
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

#### With coverage report as html
```shell
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```