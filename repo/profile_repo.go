package repo

import (
	"context"
	"log"
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
