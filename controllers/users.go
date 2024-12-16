package controllers

import (
	"example/postman/lib"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	users := Users

	sort.Slice(users, func(i, j int) bool {
		return users[i].Fullname < users[j].Fullname
	})

	if search == "" {
		if page*limit > len(users) {
			ctx.JSON(http.StatusOK, Response{
				Succsess: true,
				Message:  "list all users",
				Results:  users[(page-1)*limit : len(Users)],
			})
			return
		}
		ctx.JSON(http.StatusOK, Response{
			Succsess: true,
			Message:  "List all movies",
			Results:  users[(page-1)*limit : page*limit],
		})

	} else {
		var resUsers User
		var listDetails []User
		for _, userSearch := range users {
			if strings.Contains(strings.ToLower(userSearch.Fullname), search) || strings.Contains(userSearch.Fullname, search) {
				resUsers = userSearch
				listDetails = append(listDetails, resUsers)
			}
		}
		if resUsers == (User{}) {
			ctx.JSON(http.StatusNotFound, Response{
				Succsess: false,
				Message:  "User not found",
			})
			return
		}
		ctx.JSON(http.StatusOK, Response{
			Succsess: true,
			Message:  "detail user",
			Results:  listDetails,
		})
	}
}

func GetUserById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var resUser User

	searchUser := FindUserById(id)
	resUser = searchUser
	if resUser == (User{}) {
		ctx.JSON(http.StatusNotFound, Response{
			Succsess: false,
			Message:  "user not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, Response{
		Succsess: true,
		Message:  "Details user",
		Results:  resUser,
	})
}

func EditUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var form User
	ctx.ShouldBind(&form)
	for i, user := range Users {
		if user.Id == id {
			if form.Fullname != "" {
				Users[i].Fullname = form.Fullname
			}
			if form.Email != "" {
				if len(form.Email) < 8 || !strings.Contains(form.Email, "@") {
					ctx.JSON(http.StatusBadRequest, Response{
						Succsess: false,
						Message:  "email must be 8 character and contains @",
					})
					return
				} else {
					Users[i].Email = form.Email
				}
			}
			if form.Password != "" {
				if len(form.Password) < 6 {
					ctx.JSON(http.StatusBadRequest, Response{
						Succsess: false,
						Message:  "password length at least 6 chatacter",
					})
					return
				} else {
					Users[i].Password = lib.CreateHash(form.Password, form.Password)
				}
			}
			ctx.JSON(http.StatusOK, Response{
				Succsess: true,
				Message:  "movie detail has modify",
				Results:  Users[i],
			})
			return
		}
	}
}

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	for i, user := range Users {
		if user.Id == id {
			Users = append(Users[:i], Users[i+1:]...)
			ctx.JSON(http.StatusOK, Response{
				Succsess: true,
				Message:  "user deleted",
				Results:  user,
			})
			return
		}
	}
}
