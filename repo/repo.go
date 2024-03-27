package repo

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Account struct {
	Username string
	Password string
}

func connect() *pgx.Conn {
	urlExample := "postgres://postgres:postgres@localhost:5432/real_world"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func GetAccount(accountId string) *Account {
	conn := connect()
	defer conn.Close(context.Background())
	var user Account
	err := conn.
		QueryRow(context.Background(), "select username, password from accounts where username=$1", accountId).
		Scan(&user.Username, &user.Password)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil
	}
	fmt.Printf(user.Username)
	return &user
}
