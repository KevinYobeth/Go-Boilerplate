package grpc

import (
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	grpcLink "github.com/kevinyobeth/go-boilerplate/pkg/genproto/link"
)

func TransformToGRPCLink(linkObj *link.Link) *grpcLink.Link {
	return &grpcLink.Link{
		Id:          linkObj.ID.String(),
		Slug:        linkObj.Slug,
		Url:         linkObj.URL,
		Description: linkObj.Description,
		Total:       int32(linkObj.Total),
	}
}
