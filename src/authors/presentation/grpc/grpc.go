package grpc

import (
	"context"
	"go-boilerplate/pkg/genproto/authors"

	"go-boilerplate/src/authors/services"
	"go-boilerplate/src/authors/services/query"

	"google.golang.org/grpc"
)

type GRPCTransport struct {
	app *services.Application
}

func NewAuthorsGRPCServer(app *services.Application) GRPCTransport {
	return GRPCTransport{app: app}
}

func (g GRPCTransport) RegisterGRPCRoutes(server *grpc.Server) {
	authors.RegisterAuthorServiceServer(server, g)
}

func (g GRPCTransport) GetAuthors(c context.Context, params *authors.GetAuthorsRequest) (*authors.GetAuthorsResponse, error) {
	authorsObj, err := g.app.Queries.GetAuthors.Handle(c, query.GetAuthorsRequest{Name: &params.Name})
	if err != nil {
		return nil, err
	}

	return &authors.GetAuthorsResponse{
		Data:    TransformToGRPCAuthors(authorsObj),
		Message: "success get authors",
	}, nil
}
