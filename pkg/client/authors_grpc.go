package client

import (
	"errors"
	"go-boilerplate/pkg/genproto/authors"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAccountClient() (authors.AuthorServiceClient, error) {
	addr, ok := os.LookupEnv("GRPC_AEGIS_ADDR")

	if !ok {
		return nil, errors.New("empty env GRPC_AEGIS_ADDR")
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return authors.NewAuthorServiceClient(conn), nil
}
