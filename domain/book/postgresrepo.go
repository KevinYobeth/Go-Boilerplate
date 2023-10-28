package book

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

func NewPostgresRepo(client *gorm.DB) *PostgresRepo {
	return &PostgresRepo{
		Client: client,
	}
}

func (r *PostgresRepo) GetAll(ctx context.Context, pagination shared.LimitPagination) (GetAllBookReturn, error) {
	var books []model.Book
	var count int64

	var limit = pagination.Limit
	var page = pagination.Page

	result := r.Client.Model(&model.Book{}).
		Limit(limit).
		Offset(helper.CalculateLimitPaginationOffset(limit, page)).
		Preload("Author").
		Find(&books).
		Count(&count)

	if result.Error != nil {
		return GetAllBookReturn{}, fmt.Errorf("failed to get books: %w", result.Error)
	}

	return GetAllBookReturn{
		Books: books,
		Count: count,
	}, nil
}

func (r *PostgresRepo) GetById(ctx context.Context, bookId uuid.UUID) (model.Book, error) {
	var book model.Book

	result := r.Client.First(&book, "id = ?", bookId)

	if result.Error != nil {
		return model.Book{}, fmt.Errorf("failed to get book: %w", result.Error)
	}

	return book, nil
}

func (r *PostgresRepo) Create(ctx context.Context, book model.Book) error {
	result := r.Client.Create(book)

	if result.Error != nil {
		return fmt.Errorf("failed to add book: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) DeleteById(ctx context.Context, bookId uuid.UUID) error {
	result := r.Client.Delete(&model.Book{}, "id = ?", bookId)

	if result.Error != nil {
		return fmt.Errorf("failed to delete book: %w", result.Error)
	}

	return nil
}

func (r *PostgresRepo) UpdateById(ctx context.Context, authorId uuid.UUID, book model.Book) error {
	result := r.Client.Save(&book)

	if result.Error != nil {
		return fmt.Errorf("failed to update book: %w", result.Error)
	}

	return nil
}
