package http

import "github.com/kevinyobeth/go-boilerplate/internal/user/domain/user"

func TransformToHTTPUser(userObj *user.User) User {
	return User{
		Id:        userObj.ID,
		Email:     userObj.Email,
		FirstName: userObj.FirstName,
		LastName:  userObj.LastName,
	}
}
