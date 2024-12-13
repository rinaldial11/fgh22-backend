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
	isFound := false
	for _, searchUser := range Users {
		if searchUser.Email == form.Email {
			isFound = true
			ctx.JSON(http.StatusBadRequest, Response{
				Succsess: false,
				Message:  "email not available",
			})
			return
		}
	}

	if !isFound {
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
			// hasher, _ := argon2.CreateHash(form.Password, form.Password, argon2.DefaultParams)
			hasher := lib.CreateHash(form.Password, form.Password)
			serial++
			form.Id = serial
			form.Password = hasher
			Users = append(Users, form)

		}
	}
	ctx.JSON(http.StatusOK, Response{
		Succsess: true,
		Message:  "register success",
		Results:  form,
	})
}

func Login(ctx *gin.Context) {

	var form User
	ctx.ShouldBind(&form)

	temp := false
	for _, user := range Users {
		// isValid, _ := argon2.ComparePasswordAndHash(form.Password, form.Password, user.Password)
		isValid := lib.HashValidator(form.Password, form.Password, user.Password)
		if user.Email == form.Email && isValid {
			ctx.Header("X-Logged-user", "true")
			ctx.JSON(http.StatusOK, Response{
				Succsess: true,
				Message:  "login success",
				Results:  user,
			})
			temp = true
			return
		}
	}
	if !temp {
		ctx.JSON(http.StatusUnauthorized, Response{
			Succsess: false,
			Message:  "wrong email or password",
		})
	}
}
