package controllers

import (
	"example/postman/lib"
	"example/postman/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//	func GetAllUsers(ctx *gin.Context) {
//		search := ctx.Query("search")
//		page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
//		limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
//		order := ctx.DefaultQuery("order", "asc")
//		orderBy := ctx.DefaultQuery("sort", "id")
//		users := Users
//		if order == "asc" {
//			if orderBy == "id" {
//				sort.Slice(users, func(i, j int) bool {
//					return users[i].Id < users[j].Id
//				})
//			}
//			if orderBy == "fullname" {
//				sort.Slice(users, func(i, j int) bool {
//					return users[i].Fullname < users[j].Fullname
//				})
//			}
//		}
//		if order == "desc" {
//			if orderBy == "id" {
//				sort.Slice(users, func(i, j int) bool {
//					return users[i].Id > users[j].Id
//				})
//			}
//			if orderBy == "fullname" {
//				sort.Slice(users, func(i, j int) bool {
//					return users[i].Fullname > users[j].Fullname
//				})
//			}
//		}
//		if search == "" {
//			if page*limit > len(users) {
//				ctx.JSON(http.StatusOK, Response{
//					Succsess: true,
//					Message:  "list all users",
//					Results:  users[(page-1)*limit : len(Users)],
//				})
//				return
//			}
//			ctx.JSON(http.StatusOK, Response{
//				Succsess: true,
//				Message:  "List all movies",
//				Results:  users[(page-1)*limit : page*limit],
//			})
//			} else {
//				var resUsers User
//				var listDetails []User
//				for _, userSearch := range users {
//					if strings.Contains(strings.ToLower(userSearch.Fullname), search) || strings.Contains(userSearch.Fullname, search) {
//						resUsers = userSearch
//						listDetails = append(listDetails, resUsers)
//					}
//				}
//				if resUsers == (User{}) {
//					ctx.JSON(http.StatusNotFound, Response{
//						Succsess: false,
//						Message:  "User not found",
//					})
//					return
//				}
//				if len(listDetails) == 1 {
//					ctx.JSON(http.StatusOK, Response{
//						Succsess: true,
//						Message:  "detail user",
//						Results:  listDetails[0],
//					})
//					return
//				}
//				ctx.JSON(http.StatusOK, Response{
//					Succsess: true,
//					Message:  "detail user",
//					Results:  listDetails,
//				})
//			}
//		}
func GetAllUsers(ctx *gin.Context) {
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	allUsers := models.GetAllUsers(page, limit)

	foundUser := models.SearchUserByEmail(search)
	if search != "" {
		if len(foundUser) == 1 {
			ctx.JSON(http.StatusOK, models.Response{
				Succsess: true,
				Message:  "list all users",
				Results:  foundUser[0],
			})
			return
		}
		ctx.JSON(http.StatusOK, models.Response{
			Succsess: true,
			Message:  "list all users",
			Results:  foundUser,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "list all users",
		Results:  allUsers,
	})
}

func GetUserById(ctx *gin.Context) {
	idUser, _ := strconv.Atoi(ctx.Param("id"))
	foundUser := models.SelectOneUsers(idUser)

	if foundUser == (models.User{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "user not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "Details user",
		Results:  foundUser,
	})
}

func EditUser(ctx *gin.Context) {
	idUser, _ := strconv.Atoi(ctx.Param("id"))
	foundUser := models.SelectOneUsers(idUser)
	if foundUser == (models.User{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "user not found",
		})
		return
	}

	ctx.ShouldBind(&foundUser)
	if len(foundUser.Email) < 8 || !strings.Contains(foundUser.Email, "@") {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  "email must be 8 character and contains @",
		})
		return
	}
	if !strings.Contains(foundUser.Password, "$argon2i$v=19$m=65536,t=1,p=2$OIIAw9F7QeTBo4nWAfKgLQ$UEZ3jiaGXUw1oZ6TFm/PXN8a6G9RsYKGbbUxYdXZc54") {
		if foundUser.Password != "" {
			if len(foundUser.Password) < 6 {
				ctx.JSON(http.StatusBadRequest, models.Response{
					Succsess: false,
					Message:  "password length at least 6 chatacter",
				})
				return
			}
			foundUser.Password = lib.CreateHash(foundUser.Password)
		}
	}
	updatedUser := models.UpdateUser(foundUser)
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "movie detail has modify",
		Results:  updatedUser,
	})
}

func DeleteUser(ctx *gin.Context) {
	idUser, _ := strconv.Atoi(ctx.Param("id"))
	user := models.SelectOneUsers(idUser)

	if user == (models.User{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "user not found",
		})
		return
	}
	deletedUser := models.DropUser(idUser)
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "user deleted successfully",
		Results:  deletedUser,
	})
}
