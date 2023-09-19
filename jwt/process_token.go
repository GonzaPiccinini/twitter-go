package jwt

import (
	"errors"
	"strings"

	"github.com/GonzaPiccinini/twitter-go/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var UserID string

func ProcessToken(token string, JWTSign string) (*models.Claim, bool, string, error) {
	key := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, "", errors.New("invalid format token")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {

	}

	if !tkn.Valid {
		return &claims, false, "", errors.New("invalid token")
	}

	return &claims, false, "", err
}
