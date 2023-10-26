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

func (r *PostgresRepo) Insert(ctx context.Context, author model.Author) error {
	result := r.Client.Create(author)

	if result.Error != nil {
		return fmt.Errorf("failed to add to database: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) GetById(ctx context.Context, authorId uuid.UUID) (model.Author, error) {
	var author model.Author

	result := r.Client.First(&author, "id = ?", authorId)

	if result.Error != nil {
		return model.Author{}, fmt.Errorf("failed to add to database: %w", result.Error)
	}

	return author, nil
}

func (r *PostgresRepo) DeleteById(ctx context.Context, authorId uuid.UUID) error {
	result := r.Client.Delete(&model.Author{}, "id = ?", authorId)

	if result.Error != nil {
		return fmt.Errorf("failed to delete: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) Update(ctx context.Context, authorId uuid.UUID, author model.Author) error {
	result := r.Client.Save(&author)

	if result.Error != nil {
		return fmt.Errorf("failed to update: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllAuthorReturn, error) {
	var authors []model.Author

	var limit = pagination.Limit
	var page = pagination.Page

	result := r.Client.
		Limit(limit).
		Offset(helper.CalculateLimitPaginationOffset(limit, page)).
		Preload("Books").
		Find(&authors)

	if result.Error != nil {
		return GetAllAuthorReturn{}, fmt.Errorf("failed to add to database: %w", result.Error)
	}

	return GetAllAuthorReturn{
		Authors: authors,
	}, nil
}
