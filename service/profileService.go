package service

import (
	"rworld/dto"
	"rworld/repo"
)

func GetProfile(username string) dto.Profile {
	p := toProfileDto(*GetUserByUsername(username))
	if repo.IsFollowing(username, GetSubject()) {
		p.Following = true
	}
	return p
}

func Follow(username string) error {
	return repo.Follow(username, GetSubject())
}

func Unfollow(username string) error {
	return repo.Unfollow(username, GetSubject())
}

func toProfileDto(user dto.User) dto.Profile {
	return dto.Profile{Username: user.Username, Bio: user.Bio, Image: user.Image}
}
