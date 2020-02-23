package test

import (
	"bryson.k8s/kubernetes-authn-service/auth"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func TestConn(t *testing.T) {
	fmt.Println("test")
}

func TestAuth(t *testing.T) {
	tokenString, err := auth.CreateToken("bryson", "manager")
	if err != nil {
		fmt.Println("err: ", err)
	}
	claims, err := auth.ParseToken(tokenString)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(claims.(jwt.MapClaims)["username"].(string))
}