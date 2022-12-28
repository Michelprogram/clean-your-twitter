package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	TwitterId string `json:"twitter_id"`
	jwt.StandardClaims
}

type JWT struct {
	Token string
	Time  time.Time
}
