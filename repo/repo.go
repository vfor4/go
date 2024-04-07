package repo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"rworld/dto"

	"github.com/jackc/pgx/v5"
)

func connect() *pgx.Conn {
	url := "postgres://postgres:postgres@localhost:5432/real_world"
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func GetAccountByEmail(email string) *dto.User {
	conn := connect()
	defer conn.Close(context.Background())
	var user dto.User
	err := conn.
		QueryRow(context.Background(), "select email, username, bio, image from accounts where email=$1", email).
		Scan(&user.Email, &user.Username, &user.Bio, &user.Image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
	return &user
}

func GetAccountByUsername(username string) *dto.User {
	conn := connect()
	defer conn.Close(context.Background())
	var user dto.User
	err := conn.
		QueryRow(context.Background(), "select email, username, bio, image from accounts where username=$1", username).
		Scan(&user.Email, &user.Username, &user.Bio, &user.Image)
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

func userExistsByEmail(email string) bool {
	conn := connect()
	defer conn.Close(context.Background())
	var exits int
	query := "select 1 from accounts where email=$1"
	err := conn.
		QueryRow(context.Background(), query, email).
		Scan(&exits)
	if err != nil {
		fmt.Fprintf(os.Stderr, "userExistsByEmail - QueryRow failed: %v\n", err)
		return false
	}
	return true
}

func SignUp(user *dto.SignUpUser) error {
	if userExistsByEmail(user.Email) {
		return errors.New("Signup email is already exits")
	}
	pgx.BeginFunc(context.Background(), connect(), func(tx pgx.Tx) error {
		_, err := tx.
			Conn().
			Exec(context.Background(), "insert into accounts(username, password, email) values ($1, $2, $3)", user.Username, user.Password, user.Email)
		return err
	})
	return nil
}

func UpdateUser(user dto.User) error {
	if !userExistsByEmail(user.Email) {
		return errors.New("User's email is not found")
	}
	pgx.BeginFunc(context.Background(), connect(), func(tx pgx.Tx) error {
		_, err := tx.
			Conn().
			Exec(context.Background(), "update accounts set username=$1, password=$2, email=$3, image=$4, bio=$5 where email=$3", user.Username, user.Password, user.Email, user.Image, user.Bio)
		if err != nil {
			log.Printf("Repo - UpdateUser - cannot update, %v", err)
		}
		return err
	})
	return nil
}
