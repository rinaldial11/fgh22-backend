package controllers

import (
	"example/postman/lib"
	"example/postman/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMovies(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	allMovies := models.GetAllMovies(page, limit)
	count := models.CountData(search)

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
		Results:  allMovies,
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
	if foundMovie == (models.Movie{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "movie not found",
		})
		return
	}
	ctx.ShouldBind(&foundMovie)
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
