package models

import (
	"context"
	"example/postman/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Movie struct {
	Id          int    `json:"id"`
	Title       string `json:"title" form:"title"`
	Image       string `json:"image" form:"image"`
	Genre       string `json:"genre" form:"genre"`
	Releasedate string `json:"releaseDate" form:"release_date" db:"release_date"`
	Author      string `json:"author" form:"author"`
	Duration    string `json:"duration" form:"duration"`
	Casts       string `json:"casts" form:"casts"`
	Synopsis    string `json:"synopsis" form:"synopsis"`
	// CreatedAt   time.Time  `json:"createdAt" db:"created_date"`
	// UpdatedAt   *time.Time `json:"updatedAt" db:"updated_date"`
}

type ListMovies []Movie

func GetAllMovies(page int, limit int) ListMovies {
	conn := lib.DB()
	defer conn.Close(context.Background())

	offset := (page - 1) * limit
	rows, err := conn.Query(context.Background(), `
		SELECT id, title, image, genre, release_date, author, duration, casts, synopsis
		FROM movies
		ORDER BY id ASC
		OFFSET $1 
		LIMIT $2
	`, offset, limit)
	if err != nil {
		fmt.Println(err)
	}
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[Movie])
	if err != nil {
		fmt.Println(err)
	}
	return users
}

func SearchMovieByTitle(title string, page int, limit int) ListMovies {
	conn := lib.DB()
	defer conn.Close(context.Background())

	offset := (page - 1) * limit

	titleSubstring := "%" + title + "%"
	rows, err := conn.Query(context.Background(), `
		SELECT id, title, image, genre, release_date, author, duration, casts, synopsis
		FROM movies
		WHERE 
		title ILIKE $1
    ORDER BY title ASC
		OFFSET $2
		LIMIT $3
	`, titleSubstring, offset, limit)
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
    SELECT id, title, image, genre, release_date, author, duration, casts, synopsis
    FROM movies
    WHERE
    id = $1
  `, idMovie).Scan(&movie.Id, &movie.Title, &movie.Image, &movie.Genre, &movie.Releasedate, &movie.Author, &movie.Duration, &movie.Casts, &movie.Synopsis)
	return movie
}

func AddMovie(movieData Movie) Movie {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movie Movie
	err := conn.QueryRow(context.Background(), `
		INSERT INTO movies (title, image, genre, release_date, author, duration, casts, synopsis)
		values
		($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id, title, image, genre, release_date, author, duration, casts, synopsis
	`, movieData.Title, movieData.Image, movieData.Genre, movieData.Releasedate, movieData.Author, movieData.Duration, movieData.Casts, movieData.Synopsis).Scan(&movie.Id, &movie.Title, &movie.Image, &movie.Genre, &movie.Releasedate, &movie.Author, &movie.Duration, &movie.Casts, &movie.Synopsis)
	if err != nil {
		fmt.Println(err)
	}
	return movie
}

func UpdateMovie(movieData Movie) Movie {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var updatedMovie Movie
	conn.QueryRow(context.Background(), `
		UPDATE movies SET title=$1, image=$2, genre=$3,  release_date=$4, author=$5, duration=$6, casts=$7, synopsis=$8 WHERE id=$9
		RETURNING id, title, image, genre, release_date, author, duration, casts, synopsis
	`, movieData.Title, movieData.Image, movieData.Genre, movieData.Author, movieData.Duration, movieData.Casts, movieData.Synopsis, movieData.Id).Scan(&updatedMovie.Id, &updatedMovie.Title, &updatedMovie.Image, &updatedMovie.Genre, &updatedMovie.Releasedate, &updatedMovie.Author, &updatedMovie.Duration, &updatedMovie.Casts, &updatedMovie.Synopsis)
	return updatedMovie
}

func DropMovie(id int) Movie {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var deletedMovie Movie
	conn.QueryRow(context.Background(), `
		DELETE FROM movies
		WHERE id = $1
		RETURNING id, title, image, genre, release_date, author, duration, casts, synopsis
	`, id).Scan(&deletedMovie.Id, &deletedMovie.Title, &deletedMovie.Image, &deletedMovie.Releasedate, &deletedMovie.Genre, &deletedMovie.Author, &deletedMovie.Duration, &deletedMovie.Casts, &deletedMovie.Synopsis)
	return deletedMovie
}

func CountData(search string) int {
	conn := lib.DB()
	defer conn.Close(context.Background())

	titleSubstring := "%" + search + "%"
	var total int
	conn.QueryRow(context.Background(), `
		SELECT COUNT(movies.id) 
		FROM movies
		WHERE title ILIKE $1
	`, titleSubstring).Scan(&total)
	return total
}
