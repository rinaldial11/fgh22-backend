package controllers

import (
	"example/postman/lib"
	"example/postman/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var formUser models.User
	ctx.ShouldBind(&formUser)
	found := models.FindUserByEmail(strings.ToLower(formUser.Email))
	if found != (models.User{}) {
		ctx.JSON(http.StatusBadRequest, Response{
			Succsess: false,
			Message:  "email not available",
		})
		return
	}
	if len(formUser.Email) < 8 || !strings.Contains(formUser.Email, "@") {
		ctx.JSON(http.StatusBadRequest, Response{
			Succsess: false,
			Message:  "email must be 8 character and contains @",
		})
		return
	}
	if len(formUser.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, Response{
			Succsess: false,
			Message:  "password length at least 6 chatacter",
		})
		return
	}
	hasher := lib.CreateHash(formUser.Password)
	formUser.Email = strings.ToLower(formUser.Email)
	formUser.Password = hasher
	models.AddUser(formUser)

	ctx.JSON(http.StatusOK, Response{
		Succsess: true,
		Message:  "register success",
	})
}

func Login(ctx *gin.Context) {
	var form User
	ctx.ShouldBind(&form)

	// user := FindUserByEmail(form.Email)
	user := models.FindUserByEmail(form.Email)
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
