package controllers

import (
	"example/postman/lib"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var serial int = len(Users)

func Register(ctx *gin.Context) {
	var form User
	ctx.ShouldBind(&form)
	found := FindUserByEmail(form.Email)
	if found != (User{}) {
		ctx.JSON(http.StatusBadRequest, Response{
			Succsess: false,
			Message:  "email not available",
		})
		return
	}

	if len(form.Email) < 8 || !strings.Contains(form.Email, "@") {
		ctx.JSON(http.StatusBadRequest, Response{
			Succsess: false,
			Message:  "email must be 8 character and contains @",
		})
		return
	}
	if len(form.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, Response{
			Succsess: false,
			Message:  "password length at least 6 chatacter",
		})
		return
	} else {
		hasher := lib.CreateHash(form.Password, form.Password)
		serial++
		form.Id = serial
		form.Password = hasher
		Users = append(Users, form)

	}
	ctx.JSON(http.StatusOK, Response{
		Succsess: true,
		Message:  "register success",
	})
}

func Login(ctx *gin.Context) {
	var form User
	ctx.ShouldBind(&form)

	user := FindUserByEmail(form.Email)
	isValid := lib.HashValidator(form.Password, form.Password, user.Password)
	if isValid {
		token := lib.GenerateToken(struct {
			UserID int `json:"userId"`
		}{
			UserID: user.Id,
		})
		ctx.JSON(http.StatusOK, Response{
			Succsess: true,
			Message:  "login success",
			Results:  token,
		})
		return
	}
	ctx.JSON(http.StatusUnauthorized, Response{
		Succsess: false,
		Message:  "wrong email or password",
	})
}
