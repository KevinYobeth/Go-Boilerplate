package http

import (
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/token"
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
