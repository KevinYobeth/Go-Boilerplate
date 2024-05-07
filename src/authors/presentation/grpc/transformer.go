package grpc

import (
	grpcAuthors "go-boilerplate/pkg/genproto/authors"
	"go-boilerplate/src/authors/domain/authors"
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
