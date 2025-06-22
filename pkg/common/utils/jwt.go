package utils

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/pkg/common/constants"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/entity"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/errors"
)

func GetClaimsFromContext(ctx context.Context) (*entity.Claims, error) {
	claims := ReadFromCtx(ctx, constants.ContextKeyClaims)

	if claims == nil {
		return nil, errors.NewUnauthenticatedError(nil)
	}

	return claims.(*entity.Claims), nil
}
