package book

import (
	"context"
	"fmt"
	"library/shared"
	helper "library/shared/utils"

	"gorm.io/gorm"
)

type PostgresRepo struct {
	Client *gorm.DB
}

func (r *PostgresRepo) Insert(ctx context.Context, book Book) error {
	result := r.Client.Create(book)

	if result.Error != nil {
		return fmt.Errorf("failed to add to database: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) GetById(ctx context.Context, bookId string) (Book, error) {
	var book Book

	result := r.Client.First(&book, "id = ?", bookId)

	if result.Error != nil {
		return Book{}, fmt.Errorf("failed to add to database: %w", result.Error)
	}

	return book, nil
}

func (r *PostgresRepo) DeleteById(ctx context.Context, bookId string) error {
	result := r.Client.Delete(&Book{}, "id = ?", bookId)

	if result.Error != nil {
		return fmt.Errorf("failed to delete: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) Update(ctx context.Context, authorId string, book Book) error {
	result := r.Client.Save(&book)

	if result.Error != nil {
		return fmt.Errorf("failed to update: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllBookReturn, error) {
	var books []Book

	var limit = pagination.Limit
	var page = pagination.Page

	result := r.Client.Model(&Book{}).
		Limit(limit).
		Offset(helper.CalculateLimitPaginationOffset(limit, page)).
		Preload("Author").
		Find(&books)

	if result.Error != nil {
		return GetAllBookReturn{}, fmt.Errorf("failed to add to database: %w", result.Error)
	}

	return GetAllBookReturn{
		Books: books,
	}, nil
}
