package intraprocess

import (
	intraprocesscontract "github.com/kevinyobeth/go-boilerplate/internal/shared/intraprocess_contract"
	"github.com/kevinyobeth/go-boilerplate/internal/user/domain/user"
)

func TransformToIntraprocessUser(userObj *user.User) *intraprocesscontract.User {
	if userObj == nil {
		return nil
	}

	return &intraprocesscontract.User{
		ID:        userObj.ID,
		FirstName: userObj.FirstName,
		LastName:  userObj.LastName,
		Email:     userObj.Email,
		Password:  userObj.Password,
	}
}
