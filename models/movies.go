package models

import (
	"context"
	"example/postman/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Movie struct {
	Id    int    `json:"id"`
	Title string `json:"title" form:"title"`
	Image string `json:"image" form:"image"`
	Genre string `json:"genre" form:"genre"`
	// Releasedate string `json:"release_date" form:"release_date"`
	Author   string `json:"author" form:"author"`
	Duration string `json:"duration" form:"duration"`
	Casts    string `json:"casts" form:"casts"`
	Synopsis string `json:"synopsis" form:"synopsis"`
}

type ListMovies []Movie

func GetAllMovies(page int, limit int) ListMovies {
	conn := lib.DB()
	defer conn.Close(context.Background())

	offset := (page - 1) * limit
	resLimit := page * limit
	rows, err := conn.Query(context.Background(), `
		SELECT id, title, image, genre, author, duration, casts, synopsis
		FROM movies
		ORDER BY id ASC
		OFFSET $1 
		LIMIT $2
	`, offset, resLimit)
	if err != nil {
		fmt.Println(err)
	}
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[Movie])
	if err != nil {
		fmt.Println(err)
	}
	return users
}

func SearchMovieByTitle(title string) ListMovies {
	conn := lib.DB()
	defer conn.Close(context.Background())

	titleSubstring := "%" + title + "%"
	rows, err := conn.Query(context.Background(), `
		SELECT id, title, image, genre, author, duration, casts, synopsis
		FROM movies
		WHERE 
		title ILIKE $1
    ORDER BY title ASC
	`, titleSubstring)
	if err != nil {
		fmt.Println(err)
	}
	movies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Movie])
	if err != nil {
		fmt.Println(err)
	}
	return movies
}

func SelectOneMovie(idMovie int) Movie {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var movie Movie

	conn.QueryRow(context.Background(), `
    SELECT id, title, image, genre, author, duration, casts, synopsis
    FROM movies
    WHERE
    id = $1
  `, idMovie).Scan(&movie.Id, &movie.Title, &movie.Image, &movie.Genre, &movie.Author, &movie.Duration, &movie.Casts, &movie.Synopsis)
	return movie
}

func AddMovie(movieData Movie) Movie {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movie Movie
	conn.QueryRow(context.Background(), `
		INSERT INTO movies (title, image, genre, author, duration, casts, synopsis)
		values
		($1, $2, $3, $4, $5, $6, $7)
    RETURNING id, title, image, genre, author, duration, casts, synopsis
	`, movieData.Title, movieData.Image, movieData.Genre, movieData.Author, movieData.Duration, movieData.Casts, movieData.Synopsis).Scan(&movie.Id, &movie.Title, &movie.Image, &movie.Genre, &movie.Author, &movie.Duration, &movie.Casts, &movie.Synopsis)
	return movie
}

func UpdateMovie(movieData Movie) Movie {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var updatedMovie Movie
	conn.QueryRow(context.Background(), `
		UPDATE movies SET title=$1, image=$2, genre=$3, author=$4, duration=$5, casts=$6, synopsis=$7 WHERE id=$8
		RETURNING id, title, image, genre, author, duration, casts, synopsis
	`, movieData.Title, movieData.Image, movieData.Genre, movieData.Author, movieData.Duration, movieData.Casts, movieData.Synopsis, movieData.Id).Scan(&updatedMovie.Id, &updatedMovie.Title, &updatedMovie.Image, &updatedMovie.Genre, &updatedMovie.Author, &updatedMovie.Duration, &updatedMovie.Casts, &updatedMovie.Synopsis)
	return updatedMovie
}

func DropMovie(id int) Movie {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var deletedMovie Movie
	conn.QueryRow(context.Background(), `
		DELETE FROM movies
		WHERE id = $1
		RETURNING id, title, image, genre, author, duration, casts, synopsis
	`, id).Scan(&deletedMovie.Id, &deletedMovie.Title, &deletedMovie.Image, &deletedMovie.Genre, &deletedMovie.Author, &deletedMovie.Duration, &deletedMovie.Casts, &deletedMovie.Synopsis)
	return deletedMovie
}
