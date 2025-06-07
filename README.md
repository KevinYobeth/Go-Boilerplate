# Go Development Environment Setup
```
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/segmentio/golines@latest
go install github.com/bombsimon/wsl/v4/cmd...@master
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Docker
### Building the docker image for subscriber
To build the Docker image for the subscriber service, use the following command, replacing the `SCRIPT_DIR` argument with the appropriate path to your script directory:
```
docker build -t notification-authentication-subs:latest \
  -f ./docker/subscriber/Dockerfile \
  --build-arg SCRIPT_DIR=./internal/notification/presentation/subscriber/authentication \
  .
```