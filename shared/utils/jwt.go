package utils

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/shared/constants"
	"github.com/kevinyobeth/go-boilerplate/shared/entity"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
)

func GetClaimsFromContext(ctx context.Context) (*entity.Claims, error) {
	claims := ReadFromCtx(ctx, constants.ContextKeyClaims)

	if claims == nil {
		return nil, errors.NewUnauthenticatedError(nil)
	}

	return claims.(*entity.Claims), nil
}
