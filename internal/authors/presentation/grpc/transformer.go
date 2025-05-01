package grpc

import (
	"github.com/kevinyobeth/go-boilerplate/internal/authors/domain/authors"
	grpcAuthors "github.com/kevinyobeth/go-boilerplate/pkg/genproto/authors"
)

func TransformToGRPCAuthor(authorObj *authors.Author) *grpcAuthors.Author {
	return &grpcAuthors.Author{
		Id:   authorObj.ID.String(),
		Name: authorObj.Name,
	}
}

func TransformToGRPCAuthors(authorsObj []authors.Author) []*grpcAuthors.Author {
	var authors []*grpcAuthors.Author = make([]*grpcAuthors.Author, 0)
	for _, author := range authorsObj {
		authors = append(authors, TransformToGRPCAuthor(&author))
	}
	return authors
}
