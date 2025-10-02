package http

import (
	"github.com/kevinyobeth/go-boilerplate/internal/link/domain/link"
	"github.com/kevinyobeth/go-boilerplate/shared/builder/pagination"
)

func TransformToHTTPLink(linkObj *link.Link) Link {
	return Link{
		Id:          linkObj.ID,
		Slug:        linkObj.Slug,
		Url:         linkObj.URL,
		Description: linkObj.Description,
		Total:       linkObj.Total,
		CreatedAt:   linkObj.CreatedAt,
		UpdatedAt:   linkObj.UpdatedAt,
	}
}

func TransformToHTTPLinks(linksObj []link.Link) []Link {
	var links []Link = make([]Link, 0, len(linksObj))
	for _, link := range linksObj {
		links = append(links, TransformToHTTPLink(&link))
	}
	return links
}

func TransformToHTTPMetadata(metadata pagination.Metadata) PaginationMetadata {
	return PaginationMetadata{
		Next:  metadata.Next,
		Prev:  metadata.Prev,
		Total: metadata.Total,
	}
}
