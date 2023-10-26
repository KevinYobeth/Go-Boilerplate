package author

import (
	"context"
	"library/shared"
	model "library/shared/models"
	"time"

	"github.com/google/uuid"
)

func NewAuthorUseCase(repo Repo) *UseCase {
	return &UseCase{
		Repo: repo,
	}
}

func (uc *UseCase) Create(ctx context.Context, payload UpsertAuthorEntity) (model.Author, error) {
	author := model.Author{
		Id:        uuid.New(),
		Name:      payload.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := uc.Repo.Create(ctx, author)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (uc *UseCase) List(ctx context.Context, pagination shared.LimitPagination) (GetAllAuthorReturn, error) {
	authors, err := uc.Repo.GetAll(ctx, pagination)
	if err != nil {
		return GetAllAuthorReturn{}, err
	}

	return authors, nil
}

func (uc *UseCase) GetById(ctx context.Context, authorId uuid.UUID) (model.Author, error) {
	author, err := uc.Repo.GetById(ctx, authorId)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (uc *UseCase) Update(ctx context.Context, authorId uuid.UUID, payload UpsertAuthorEntity) (model.Author, error) {
	authorFromDb, err := uc.GetById(ctx, authorId)
	if err != nil {
		return model.Author{}, err
	}

	author := model.Author{
		Id:        authorFromDb.Id,
		Name:      payload.Name,
		CreatedAt: authorFromDb.CreatedAt,
		UpdatedAt: time.Now(),
	}

	err = uc.Repo.UpdateById(ctx, authorId, author)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (uc *UseCase) DeleteById(ctx context.Context, authorId uuid.UUID) error {
	err := uc.Repo.DeleteById(ctx, authorId)
	if err != nil {
		return err
	}

	return nil
}
