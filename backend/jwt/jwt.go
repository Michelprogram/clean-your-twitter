package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my-secret-key")

func GenerateJWT(twitterId string) (*JWT, error) {

	expirationTime := time.Now().Add(72000 * time.Hour)
	claims := &Claims{
		TwitterId: twitterId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	return &JWT{tokenString, expirationTime}, nil
}

func DecodeJwt(tokenString string) (string, error) {

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	for key, val := range claims {
		if key == "twitter_id" {
			return val.(string), nil
		}
	}

	return "", errors.New("incorrect jwt token")

}
