package intraprocess

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
	"github.com/kevinyobeth/go-boilerplate/internal/shared/interfaces"
	"github.com/kevinyobeth/go-boilerplate/shared/telemetry"
	"github.com/ztrue/tracerr"
)

type AuthenticationUserIntraprocessService struct {
	intraprocess interfaces.UserIntraprocess
}

type AuthenticationUserIntraprocess interface {
	GetUser(ctx context.Context, id uuid.UUID) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
}

func NewAuthenticationUserIntraprocessService(intraprocess interfaces.UserIntraprocess) AuthenticationUserIntraprocess {
	return &AuthenticationUserIntraprocessService{intraprocess: intraprocess}
}

func (i *AuthenticationUserIntraprocessService) GetUser(ctx context.Context, id uuid.UUID) (*user.User, error) {
	ctx, span := telemetry.NewIntraprocessSpan(ctx)
	defer span.End()

	user, err := i.intraprocess.GetUser(ctx, id)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return transformIntraprocessUserToDomainUser(user), nil
}

func (i *AuthenticationUserIntraprocessService) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	ctx, span := telemetry.NewIntraprocessSpan(ctx)
	defer span.End()

	user, err := i.intraprocess.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return transformIntraprocessUserToDomainUser(user), nil
}

func transformIntraprocessUserToDomainUser(userObj *interfaces.User) *user.User {
	if userObj == nil {
		return nil
	}

	return &user.User{
		ID:        userObj.ID,
		FirstName: userObj.FirstName,
		LastName:  userObj.LastName,
		Email:     userObj.Email,
		Password:  userObj.Password,
	}
}
