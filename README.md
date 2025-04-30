# About this Template

# Architecture
This project use [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) combined with [DDD (Domain Driven Design)](https://martinfowler.com/bliki/DomainDrivenDesign.html) and [CQRS (Command Query Responsibility Segregation)](https://threedots.tech/post/basic-cqrs-in-go/) pattern.

# Project Structure

# How to setup 

## Envfile

## Docker Compose
We provide a docker-compose 

## Migrations
### Generate Migration Binary

### Database Migrations

### Database Seedings

## Go Development Environment Setup
<!-- Explain scripts folder for generating oapi and proto -->

```
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/segmentio/golines@latest
go install github.com/bombsimon/wsl/v4/cmd...@master
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Starting the server
```
make run http|grpc|scheduler
```