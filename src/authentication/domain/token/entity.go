package token

import "time"

type Token struct {
	Token        string       `json:"token"`
	ExpiredAt    time.Time    `json:"expired_at"`
	RefreshToken RefreshToken `json:"refresh_token"`
}

type RefreshToken struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}
