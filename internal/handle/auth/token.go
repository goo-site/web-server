package login

import (
	"github.com/goo-site/log"
	"time"
	"web-server/internal/common/enum"

	"github.com/dgrijalva/jwt-go"
)

var SigningKey = []byte("github.com/goo-site/web-server")

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

func CreateToken(userId int64) (string, error) {
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "william", // 签名颁发者
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SigningKey)
	if err != nil {
		log.Error("[CreateToken] err:%v", err)
		return "", err
	}
	return token, nil
}

func CheckToken(tokenString string, userId int64) enum.TokenStatus {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return enum.TokenStatus_Expire
			}
		}
		return enum.TokenStatus_Invalid
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.UserId == userId {
			if time.Now().Unix() > claims.IssuedAt+60*60*2 {
				return enum.TokenStatus_Refresh
			}
			return enum.TokenStatus_Pass
		}
	}
	return enum.TokenStatus_Invalid
}
