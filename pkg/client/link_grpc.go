package client

import (
	"errors"
	"os"

	"github.com/kevinyobeth/go-boilerplate/pkg/genproto/link"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewLinkClient() (link.LinkServiceClient, error) {
	addr, ok := os.LookupEnv("GRPC_LINK_ADDRESS")

	if !ok {
		return nil, errors.New("empty env GRPC_LINK_ADDRESS")
	}

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return link.NewLinkServiceClient(conn), nil
}
