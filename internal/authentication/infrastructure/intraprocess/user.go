package intraprocess

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
	intraprocesscontract "github.com/kevinyobeth/go-boilerplate/internal/shared/intraprocess_contract"
	"github.com/kevinyobeth/go-boilerplate/shared/telemetry"
	"github.com/ztrue/tracerr"
)

type AuthenticationUserIntraprocess struct {
	intraprocess intraprocesscontract.UserInterface
}

func NewAuthenticationUserIntraprocessService(intraprocess intraprocesscontract.UserInterface) UserIntraprocess {
	return &AuthenticationUserIntraprocess{intraprocess: intraprocess}
}

func (i *AuthenticationUserIntraprocess) GetUser(ctx context.Context, id uuid.UUID) (*user.User, error) {
	ctx, span := telemetry.NewIntraprocessSpan(ctx)
	defer span.End()

	user, err := i.intraprocess.GetUser(ctx, id)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return transformIntraprocessUserToDomainUser(user), nil
}

func (i *AuthenticationUserIntraprocess) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	ctx, span := telemetry.NewIntraprocessSpan(ctx)
	defer span.End()

	user, err := i.intraprocess.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return transformIntraprocessUserToDomainUser(user), nil
}

func transformIntraprocessUserToDomainUser(userObj *intraprocesscontract.User) *user.User {
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
