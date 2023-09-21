package jwt

import (
	"context"
	"time"

	"github.com/GonzaPiccinini/twitter-go/consts"
	"github.com/GonzaPiccinini/twitter-go/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(ctx context.Context, u models.User) (string, error) {
	jwtSign := ctx.Value(models.Key(consts.JWTSIGN)).(string)
	sign := []byte(jwtSign)

	payload := jwt.MapClaims{
		"_id": u.ID.Hex(),

		"biography": u.Biography,
		"birthdate": u.Birthdate,
		"email":     u.Email,
		"firstname": u.Firstname,
		"lastname":  u.Lastname,
		"ubication": u.Ubication,
		"web":       u.Web,

		"exp": time.Now().Add(time.Hour * 3).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(sign)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
