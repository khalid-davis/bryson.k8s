package controllers

import (
	"bryson.k8s/kubernetes-authn-service/auth"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	authenticationv1 "k8s.io/api/authentication/v1"
	"log"
	"net/http"
)

type AuthController struct {
	beego.Controller
}

func (controller *AuthController) Get() {
	log.Println("test method")
	controller.Ctx.Output.SetStatus(http.StatusOK)
	controller.Data["json"] = "temp"
	controller.ServeJSON()
	return
}

//默认只生成username: bryson, group: manager

func (controller *AuthController) GetToken() {
	tokenString, err := auth.CreateToken("bryson", "manager")
	if err != nil {
		controller.Ctx.Output.SetStatus(400)
		controller.Data["json"] = err.Error()
		controller.ServeJSON()
		return
	}
	controller.Ctx.Output.SetStatus(200)
	controller.Data["json"] = tokenString
	controller.ServeJSON()
	return
}

func (controller *AuthController)Authenticate() {
	var tokenReview authenticationv1.TokenReview
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &tokenReview)
	log.Println("tokenReview: ", tokenReview)
	if err != nil {
		log.Println("error: ", err.Error())
		controller.Ctx.Output.SetStatus(http.StatusBadRequest)
		controller.Data["json"] = map[string]interface{}{
			"apiVersion": "authentication.k8s.io/v1beta1",
			"kind": "TokenReview",
			"status": authenticationv1.TokenReviewStatus{
				Authenticated: false,
			},
		}
		controller.ServeJSON()
		return
	}
	tokenString := tokenReview.Spec.Token
	log.Println("tokenString: ", tokenString)
	claim, err := auth.ParseToken(tokenString)
	if err != nil {
		log.Println("Error", err.Error())
		controller.Ctx.Output.SetStatus(http.StatusUnauthorized)
		controller.Data["json"] = map[string]interface{}{
			"apiVersion": "authentication.k8s.io/v1beta1",
			"kind": "TokenReview",
			"status": authenticationv1.TokenReviewStatus{
				Authenticated: false,
			},
		}
		controller.ServeJSON()
		return
	}

	mapClaims := claim.(jwt.MapClaims)
	log.Println("[Success] login as ", mapClaims["username"])
	controller.Ctx.Output.SetStatus(http.StatusOK)
	trs := authenticationv1.TokenReviewStatus{
		Authenticated: true,
		User:          authenticationv1.UserInfo{
			Username: mapClaims["username"].(string),
			Groups: []string{mapClaims["group"].(string)},
		},
	}
	controller.Data["json"] = map[string]interface{}{
		"apiVersion" : "authentication.k8s.io./v1beta1",
		"kind": "TokenReview",
		"status": trs,
	}
	controller.ServeJSON()
}

