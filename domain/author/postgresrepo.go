package author

import (
	"context"
	"fmt"
	helper "library/utils"

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
	return Author{}, nil
}

func (r *PostgresRepo) DeleteById(ctx context.Context, authorId string) error {
	return nil
}

func (r *PostgresRepo) Update(ctx context.Context, authorId string, author Author) error {
	return nil
}

func (r *PostgresRepo) GetAll(ctx context.Context, pagination helper.LimitPagination) (GetAllAuthorReturn, error) {
	fmt.Println("PAGINATION", pagination)

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
