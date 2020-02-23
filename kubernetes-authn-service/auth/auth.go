package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const secretKey = "bryson"

type CustomClaims struct {
	jwt.StandardClaims

	Username string `json:"username"`
	Group string `json:"group"`
}

/** 创建token */
func CreateToken(username string, group string) (string, error) {
	claims := &CustomClaims {
		StandardClaims:	jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
		},
		Username: username,
		Group: group,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}

// 解析token
func ParseToken(tokenSrt string) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	claims = token.Claims
	return
}


