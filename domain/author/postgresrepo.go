package author

import (
	"context"
	"fmt"
	"library/shared"

	"gorm.io/gorm"
)

type PostgresRepo struct {
	Client *gorm.DB
}

func (r *PostgresRepo) Insert(ctx context.Context, author Author) error {
	result := r.Client.Create(author)

	if result.Error != nil {
		return fmt.Errorf("failed to add to database: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) GetById(ctx context.Context, authorId string) (Author, error) {
	var author Author

	result := r.Client.First(&author, "id = ?", authorId)

	if result.Error != nil {
		return Author{}, fmt.Errorf("failed to add to database: %w", result.Error)
	}

	return author, nil
}

func (r *PostgresRepo) DeleteById(ctx context.Context, authorId string) error {
	result := r.Client.Delete(&Author{}, "id = ?", authorId)

	if result.Error != nil {
		return fmt.Errorf("failed to delete: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) Update(ctx context.Context, authorId string, author Author) error {
	result := r.Client.Save(&author)

	if result.Error != nil {
		return fmt.Errorf("failed to update: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllAuthorReturn, error) {
	var authors []Author

	result := r.Client.Limit(int(pagination.Limit)).Offset(int(pagination.Limit)*int(pagination.Page) - int(pagination.Limit)).Find(&authors)

	if result.Error != nil {
		return GetAllAuthorReturn{}, fmt.Errorf("failed to add to database: %w", result.Error)
	}

	return GetAllAuthorReturn{
		Authors: authors,
		Cursor:  0,
	}, nil
}
