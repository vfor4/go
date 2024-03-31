package service

import (
	"rworld/dto"
	"rworld/repo"
)

func GetAccount(username string) *dto.User {
	return repo.GetAccount(username)
}

func Loggedin(loginInfo dto.LoginInfo) bool {
	return repo.UserExists(&loginInfo)
}
