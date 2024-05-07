package client

import (
	"errors"
	"go-boilerplate/pkg/genproto/authors"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthorsClient() (authors.AuthorServiceClient, error) {
	addr, ok := os.LookupEnv("GRPC_AUTHORS_ADDRESS")

	if !ok {
		return nil, errors.New("empty env GRPC_AUTHORS_ADDRESS")
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return authors.NewAuthorServiceClient(conn), nil
}
