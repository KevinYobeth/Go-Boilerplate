package intraprocess

import (
	"github.com/kevinyobeth/go-boilerplate/internal/shared/interfaces"
	"github.com/kevinyobeth/go-boilerplate/internal/user/domain/user"
)

func TransformToIntraprocessUser(userObj *user.User) *interfaces.User {
	if userObj == nil {
		return nil
	}

	return &interfaces.User{
		ID:        userObj.ID,
		FirstName: userObj.FirstName,
		LastName:  userObj.LastName,
		Email:     userObj.Email,
		Password:  userObj.Password,
	}
}
