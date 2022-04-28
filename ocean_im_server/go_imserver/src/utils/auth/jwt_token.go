package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"

	"time"
)

type Token struct {
	Token string `json:"token"`
}

type Claims struct {
	Uid      string `json:"uid"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	NickName string `json:"nickname"`
	Version  string `json:"version"`
	Platform string `json:"platform"`
	jwt.StandardClaims
}

/*
jti：该jwt的唯一标识
iss：该jwt的签发者
iat：该jwt的签发时间
aud：该jwt的接收者
sub：该jwt的面向的用户
nbf：该jwt的生效时间,可不设置,若设置,一定要大于当前Unix UTC,否则token将会延迟生效
exp：该jwt的过期时间
*/

const secret = "ocean"

func CreateToken(claims *Claims) (string, error) {
	claims.Id = "ocean id"
	claims.Issuer = "ocean gzy"
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(signedToken string) (string, error) {
	token, validateErr := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if validateErr != nil {
		return "", validateErr
	}
	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		return claims.Uid, nil
	} else {
		return "", errors.New("token校验失败")
	}

}
