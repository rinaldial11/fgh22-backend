package controllers

import (
	"encoding/json"
	"example/postman/lib"
	"example/postman/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllMovies(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	sorting := ctx.DefaultQuery("order", "ASC")
	sortBy := ctx.DefaultQuery("sort_by", "id")

	if sorting != "ASC" {
		sorting = "DESC"
	}

	var movies models.ListMovies
	var count int
	modifyRequestUri := fmt.Sprintf("count+%s", ctx.Request.RequestURI)

	get := lib.GetFromRedis(ctx.Request.RequestURI)
	getCount := lib.GetFromRedis(modifyRequestUri)

	if get.Val() != "" {
		raw := []byte(get.Val())
		json.Unmarshal(raw, &movies)
	} else {
		movies = models.GetAllMovies(page, limit, sortBy, sorting)
		encoded, _ := json.Marshal(movies)
		lib.SetToRedis(ctx.Request.RequestURI, encoded)
	}

	if getCount.Val() != "" {
		raw := []byte(getCount.Val())
		json.Unmarshal(raw, &count)
	} else {
		count = models.CountData(search)
		encoded, _ := json.Marshal(count)
		lib.SetToRedis(modifyRequestUri, encoded)
	}

	foundMovie := models.SearchMovieByTitle(search, page, limit)

	if search != "" {
		if len(foundMovie) == 1 {
			ctx.JSON(http.StatusOK, models.Response{
				Succsess: true,
				Message:  "list all movies",
				PageInfo: models.PageInfo(lib.GetPageInfo(page, limit, count)),
				Results:  foundMovie[0],
			})
			return
		}
		ctx.JSON(http.StatusOK, models.Response{
			Succsess: true,
			Message:  "list all movies",
			PageInfo: models.PageInfo(lib.GetPageInfo(page, limit, count)),
			Results:  foundMovie,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "list all movies",
		PageInfo: models.PageInfo(lib.GetPageInfo(page, limit, count)),
		Results:  movies,
	})
}

func GetMovieById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	foundMovie := models.SelectOneMovie(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  "Wrong movie id format",
		})
		return
	}

	if foundMovie == (models.Movie{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "movie not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "Details movie",
		Results:  foundMovie,
	})
}

func AddMovie(ctx *gin.Context) {
	var formMovie models.Movie
	ctx.ShouldBind(&formMovie)
	file, _ := ctx.FormFile("images")

	if file != nil {
		filename := uuid.New().String()
		splitedfilename := strings.Split(file.Filename, ".")
		ext := splitedfilename[len(splitedfilename)-1]
		if ext != "jpg" && ext != "png" {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Succsess: false,
				Message:  "wrong file format",
			})
			return
		}
		storedFile := fmt.Sprintf("%s.%s", filename, ext)
		ctx.SaveUploadedFile(file, fmt.Sprintf("uploads/movies/%s", storedFile))
		formMovie.Image = storedFile
	}

	newlyAdded := models.AddMovie(formMovie)
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "movie added",
		Results:  newlyAdded,
	})
}

func EditMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	foundMovie := models.SelectOneMovie(id)
	ctx.ShouldBind(&foundMovie)
	file, _ := ctx.FormFile("images")

	if foundMovie == (models.Movie{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "movie not found",
		})
		return
	}
	if file != nil {
		filename := uuid.New().String()
		splitedfilename := strings.Split(file.Filename, ".")
		ext := splitedfilename[len(splitedfilename)-1]
		if ext != "jpg" && ext != "png" {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Succsess: false,
				Message:  "wrong file format",
			})
			return
		}
		storedFile := fmt.Sprintf("%s.%s", filename, ext)
		ctx.SaveUploadedFile(file, fmt.Sprintf("uploads/movies/%s", storedFile))
		foundMovie.Image = storedFile
		fmt.Println(foundMovie)
	}

	updatedMovie := models.UpdateMovie(foundMovie)

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "movie detail has modify",
		Results:  updatedMovie,
	})
}

func DeleteMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	found := models.SelectOneMovie(id)
	if found == (models.Movie{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "movie not found",
		})
		return
	}
	deletedMovie := models.DropMovie(id)

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "movie deleted",
		Results:  deletedMovie,
	})
}
