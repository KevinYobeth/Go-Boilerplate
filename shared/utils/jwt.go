package utils

import (
	"context"
	"go-boilerplate/shared/constants"
	"go-boilerplate/shared/entity"
	"go-boilerplate/shared/errors"
)

func GetClaimsFromContext(ctx context.Context) (*entity.Claims, error) {
	claims := ReadFromCtx(ctx, constants.ContextKeyClaims)

	if claims == nil {
		return nil, errors.NewUnauthenticatedError(nil)
	}

	return claims.(*entity.Claims), nil
}
