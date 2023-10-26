package book

import (
	"context"
	"library/domain/author"
	"library/shared"
	model "library/shared/models"
	"time"

	"github.com/google/uuid"
)

func NewBookUseCase(repo Repo, authorUseCase author.UseCase) *UseCase {
	return &UseCase{
		Repo:          repo,
		AuthorUseCase: authorUseCase,
	}
}

func (uc *UseCase) Create(ctx context.Context, payload UpsertBookEntity) (model.Book, error) {
	author, err := uc.GetById(ctx, payload.AuthorId)
	if err != nil {
		return model.Book{}, err
	}

	// TODO: Add author check

	book := model.Book{
		Id:        uuid.New(),
		Title:     payload.Title,
		AuthorId:  author.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return book, nil
}

func (uc *UseCase) List(ctx context.Context, pagination shared.LimitPagination) (GetAllBookReturn, error) {
	books, err := uc.Repo.GetAll(ctx, pagination)
	if err != nil {
		return GetAllBookReturn{}, err
	}

	return books, nil
}

func (uc *UseCase) GetById(ctx context.Context, bookId uuid.UUID) (model.Book, error) {
	book, err := uc.Repo.GetById(ctx, bookId)
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (uc *UseCase) Update(ctx context.Context, bookId uuid.UUID, payload UpsertBookEntity) (model.Book, error) {
	bookFromDB, err := uc.Repo.GetById(ctx, bookId)
	if err != nil {
		return model.Book{}, err
	}

	var authorId uuid.UUID

	if payload.AuthorId != uuid.Nil {
		author, err := uc.AuthorUseCase.GetById(ctx, payload.AuthorId)
		if err != nil {
			return model.Book{}, err
		}

		authorId = author.Id
	} else {
		authorId = bookFromDB.AuthorId
	}

	book := model.Book{
		Id:        bookFromDB.Id,
		Title:     payload.Title,
		AuthorId:  authorId,
		CreatedAt: bookFromDB.CreatedAt,
		UpdatedAt: time.Now(),
	}

	err = uc.Repo.Update(ctx, bookId, book)
	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}

func (uc *UseCase) DeleteById(ctx context.Context, bookId uuid.UUID) error {
	err := uc.Repo.DeleteById(ctx, bookId)
	if err != nil {
		return err
	}

	return nil
}
