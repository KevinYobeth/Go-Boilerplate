package author

import (
	"context"
	"fmt"
	"library/shared"
	model "library/shared/models"
	helper "library/shared/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	Client *gorm.DB
}

func NewAuthorPostgresRepo(client *gorm.DB) *PostgresRepo {
	return &PostgresRepo{
		Client: client,
	}
}

func (r *PostgresRepo) GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllAuthorReturn, error) {
	var authors []model.Author
	var count int64

	var limit = pagination.Limit
	var page = pagination.Page

	result := r.Client.Model(&model.Author{}).
		Limit(limit).
		Offset(helper.CalculateLimitPaginationOffset(limit, page)).
		Preload("Books").
		Find(&authors).
		Count(&count)

	if result.Error != nil {
		return GetAllAuthorReturn{}, fmt.Errorf("failed to get items: %w", result.Error)
	}

	return GetAllAuthorReturn{
		Authors: authors,
		Count:   count,
	}, nil
}

func (r *PostgresRepo) GetById(ctx context.Context, authorId uuid.UUID) (model.Author, error) {
	var author model.Author

	result := r.Client.First(&author, "id = ?", authorId)

	if result.Error != nil {
		return model.Author{}, fmt.Errorf("failed to get item: %w", result.Error)
	}

	return author, nil
}

func (r *PostgresRepo) Create(ctx context.Context, author model.Author) error {
	result := r.Client.Create(author)

	if result.Error != nil {
		return fmt.Errorf("failed to add: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) UpdateById(ctx context.Context, authorId uuid.UUID, author model.Author) error {
	result := r.Client.Save(&author)

	if result.Error != nil {
		return fmt.Errorf("failed to update: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) DeleteById(ctx context.Context, authorId uuid.UUID) error {
	result := r.Client.Delete(&model.Author{}, "id = ?", authorId)

	if result.Error != nil {
		return fmt.Errorf("failed to delete: %w", result.Error)
	}

	return nil
}
