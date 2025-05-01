package http

import (
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/token"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
)

func TransformToHTTPToken(tokenObj *token.Token) Token {
	return Token{
		Token:     tokenObj.Token,
		ExpiredAt: tokenObj.ExpiredAt,
		RefreshToken: RefreshToken{
			Token:     tokenObj.RefreshToken.Token,
			ExpiredAt: tokenObj.RefreshToken.ExpiredAt,
		},
	}
}

func TransformToHTTPUser(userObj *user.User) User {
	return User{
		Id:        userObj.ID,
		Email:     userObj.Email,
		FirstName: userObj.FirstName,
		LastName:  userObj.LastName,
	}
}
