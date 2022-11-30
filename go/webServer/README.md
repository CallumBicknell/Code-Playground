# Go Gin Boilerplate
> A starter project with Golang, Gin and DynamoDB

### Boilerplate structure

```
.
├── Makefile
├── Procfile
├── README.md
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controllers
│   └── user.go
├── db
│   └── db.go
├── forms
│   └── user.go
├── header.jpg
├── main.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
└── server
    ├── router.go
    └── server.go
```

## Installation

```sh
make deps
```

## Usage example

`curl http://localhost:8888/health`

## Development setup

Running DynamoDB on Docker Image:

check this project: [vsouza/docker-dynamoDB-local](https://github.com/vsouza/docker-dynamoDB-local)
