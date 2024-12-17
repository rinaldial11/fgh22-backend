package models

import (
	"context"
	"example/postman/lib"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ListUsers []User

func SelectOneUsers(idUser int) User {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var user User

	conn.QueryRow(context.Background(), `
    SELECT id, email, password
    FROM users
    WHERE
    id = $1
  `, idUser).Scan(&user.Id, &user.Email, &user.Password)
	return user
}

func GetAllUsers() ListUsers {
	conn := lib.DB()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `
		SELECT id, '' as fullname, email, password
		FROM users
	`)
	if err != nil {
		fmt.Println(err)
	}
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	if err != nil {
		fmt.Println(err)
	}
	return users
}

func FindUserByEmail(email string) User {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var user User
	conn.QueryRow(context.Background(), `
		SELECT id, email, password
		FROM users
		WHERE
		email = $1
	`, email).Scan(&user.Id, &user.Email, &user.Password)
	return user
}

func AddUser(userData User) User {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var user User
	conn.QueryRow(context.Background(), `
		INSERT INTO users (email, password)
		values
		($1, $2)
	`, userData.Email, userData.Password).Scan(&user.Email, &user.Password)
	return user
}

func UpdateUser(userData User) User {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var updatedUser User
	conn.QueryRow(context.Background(), `
		UPDATE users SET email=$1, password=$2 WHERE id=$3
		RETURNING id, email, password
	`, userData.Email, userData.Password, userData.Id).Scan(&updatedUser.Id, &updatedUser.Email, &updatedUser.Password)
	return updatedUser
}
