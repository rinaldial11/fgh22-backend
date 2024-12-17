package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	connstring := "postgres://rinaldi:1@172.17.0.2:5432/funtastix"
	conn, err := pgx.Connect(context.Background(), connstring)
	if err != err {
		fmt.Println(err)
		os.Exit(1)
	}
	return conn
}
