package repo

import (
	"context"
	"fmt"
	"os"
	"rworld/dto"

	"github.com/jackc/pgx/v5"
)

func connect() *pgx.Conn {
	urlExample := "postgres://postgres:postgres@localhost:5432/real_world"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func GetAccount(email string) *dto.User {
	conn := connect()
	defer conn.Close(context.Background())
	var user dto.User
	err := conn.
		QueryRow(context.Background(), "select email, username from accounts where email=$1", email).
		Scan(&user.Email, &user.Username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	return &user
}

func UserExists(loginInfo *dto.LoginInfo) bool {
	conn := connect()
	defer conn.Close(context.Background())
	var exits int
	query := "select 1 from accounts where email=$1 and password=$2"
	err := conn.
		QueryRow(context.Background(), query, loginInfo.Email, loginInfo.Password).
		Scan(&exits)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return false
	}
	return true
}
