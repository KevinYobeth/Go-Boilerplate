package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services"
	"github.com/kevinyobeth/go-boilerplate/internal/link/services/query"
	"github.com/kevinyobeth/go-boilerplate/pkg/genproto/link"
	"github.com/ztrue/tracerr"
	"google.golang.org/grpc"
)

type GRPCTransport struct {
	app *services.Application
}

func NewLinkGRPCServer(app *services.Application) GRPCTransport {
	return GRPCTransport{app: app}
}

func (g GRPCTransport) RegisterGRPCRoutes(server *grpc.Server) {
	link.RegisterLinkServiceServer(server, g)
}

func (g GRPCTransport) GetLink(c context.Context, params *link.GetLinkRequest) (*link.GetLinkResponse, error) {
	id, err := uuid.Parse(params.Id)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	userId, err := uuid.Parse(params.UserId)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	linksObj, err := g.app.Queries.GetLink.Handle(c, &query.GetLinkRequest{
		ID:     id,
		UserID: userId,
	})
	if err != nil {
		return nil, err
	}

	return &link.GetLinkResponse{
		Data:    TransformToGRPCLink(linksObj),
		Message: "success get links",
	}, nil
}
