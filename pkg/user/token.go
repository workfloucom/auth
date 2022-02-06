package user

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	Auth    string
	Refresh string
}

func NewToken(u User, authSecret, refreshSecret string) (*Token, error) {
	at := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub":   u.ID,
			"exp":   time.Now().UTC().Add(15 * time.Minute).Unix(),
			"scope": "",
		},
	)

	authToken, err := at.SignedString([]byte(authSecret))

	if err != nil {
		return nil, err
	}

	rt := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": u.ID,
			"exp": time.Now().UTC().Add(24 * time.Hour * 30).Unix(),
		},
	)

	refreshToken, err := rt.SignedString([]byte(refreshSecret))

	if err != nil {
		return nil, err
	}

	return &Token{
		Auth:    authToken,
		Refresh: refreshToken,
	}, nil
}
