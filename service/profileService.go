package service

import (
	"rworld/dto"
	"rworld/repo"
)

func GetProfile(username string) dto.User {
	user := GetUserByUsername(username)
	if repo.IsFollowing(username, GetSubject()) {
		user.Following = true
	}
	return *user
}
