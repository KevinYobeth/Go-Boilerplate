package http

import "go-boilerplate/src/authentication/domain/user"

func TransformToHTTPUser(userObj *user.User) User {
	return User{
		Id:        userObj.ID,
		Email:     userObj.Email,
		FirstName: userObj.FirstName,
		LastName:  userObj.LastName,
	}
}
