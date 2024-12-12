package controllers

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ListUsers []User

var Users = ListUsers{
	{
		Id:       1,
		Fullname: "Budiono Siregar",
		Email:    "budi@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$xYjCiAsMbo/xmH8I/4SkeQ$MvbfLtQCqyryc+p9ghXxwEniSGDZLMF2ckH7+Hzxqq0",
	},
	{
		Id:       2,
		Fullname: "Endra Prasmanan",
		Email:    "endra@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$NRaP9cPxjmhEqCvwF5+4ig$hY+X9oq8wlA3aWKEiwx0z/THUkypKplR4C0j+jv2qtA",
	},
	{
		Id:       3,
		Fullname: "Rama Ajarindong",
		Email:    "rama@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$ksC1dsVheRE4cfO9fR8odw$+hStsdH+e9w68Zxn30RMKWLjEcIFXVhhQ6EIGfbKeoU",
	},
	{
		Id:       4,
		Fullname: "Adiv Bened",
		Email:    "adiv@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$vRI56QMSk9gBXGnbLJ3XiQ$O7hU2yZt7zSTokAueXwcQAEX7lZs6ufAEyddfjc12vk",
	},
	{
		Id:       5,
		Fullname: "Nanda Brew",
		Email:    "nanda@mail.com",
		Password: "$argon2i$v=19$m=65536,t=1,p=2$gCjUU9eQbquuRJcNEoQm9g$bMrkwmI9O4bA23xLEqVEYX26uc1s0k/pKyrPkaC6J1c",
	},
}

func GetAllUsers(ctx *gin.Context) {
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "3"))
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

		for _, userSearch := range users {
			if strings.Contains(strings.ToLower(userSearch.Fullname), search) || strings.Contains(userSearch.Fullname, search) {
				resUsers = userSearch
				ctx.JSON(http.StatusOK, Response{
					Succsess: true,
					Message:  "detail user",
					Results:  resUsers,
				})
			}
		}
		if resUsers == (User{}) {
			ctx.JSON(http.StatusNotFound, Response{
				Succsess: false,
				Message:  "User not found",
			})
		}
	}
}

func GetUserById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	users := Users

	var resUser User

	for _, searchUser := range users {
		if searchUser.Id == id {
			resUser = searchUser
		}
	}
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
					Users[i].Password = form.Password
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
