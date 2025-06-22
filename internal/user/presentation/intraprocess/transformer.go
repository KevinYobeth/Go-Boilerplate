package intraprocess

import (
	"github.com/kevinyobeth/go-boilerplate/internal/shared/contract"
	"github.com/kevinyobeth/go-boilerplate/internal/user/domain/user"
)

func TransformToIntraprocessUser(userObj *user.User) *contract.User {
	if userObj == nil {
		return nil
	}

	return &contract.User{
		ID:        userObj.ID,
		FirstName: userObj.FirstName,
		LastName:  userObj.LastName,
		Email:     userObj.Email,
		Password:  userObj.Password,
	}
}
