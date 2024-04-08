package repo

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func IsFollowing(username, follower string) bool {
	conn := connect()
	defer conn.Close(context.Background())
	var t bool
	err := conn.
		QueryRow(context.Background(), "select true from user_follower where username=$1 and follower=$2", username, follower).
		Scan(&t)
	if err != nil {
		log.Printf("IsFollowing - repo error, %v", err)
	}
	return t
}

func Follow(username, follower string) error {
	conn := connect()
	defer conn.Close(context.Background())
	err := pgx.BeginFunc(context.Background(), connect(), func(tx pgx.Tx) error {
		_, err := tx.
			Conn().
			Exec(context.Background(), "insert into user_follower values ($1, $2)", username, follower)
		return err
	})
	if err != nil {
		log.Printf("Follow - err, %v", err)
		return err
	}
	return nil
}

func Unfollow(username, follower string) error {
	conn := connect()
	defer conn.Close(context.Background())
	log.Printf("delete from user_follower where username=%s and follower=%s", username, follower)
	err := pgx.BeginFunc(context.Background(), connect(), func(tx pgx.Tx) error {
		_, err := tx.
			Conn().
			Exec(context.Background(), "delete from user_follower where username=$1 and follower=$2", username, follower)
		if err != nil {
			log.Printf("Unfollow - err, %v", err)
		}
		return err
	})
	if err != nil {
		log.Printf("Follow - err, %v", err)
		return err
	}
	return nil
}
