package service

import (
	"rworld/dto"
	"rworld/repo"
)

func GetUserByUsername(username string) *dto.User {
	return repo.GetAccountByUsername(username)
}

func GetUserByEmail(email string) *dto.User {
	return repo.GetAccountByEmail(email)
}

func Loggedin(loginInfo dto.LoginInfo) bool {
	return repo.UserExists(&loginInfo)
}

func SignUp(user dto.SignUpUser) error {
	return repo.SignUp(&user)
}

func UpdateUser(user dto.User) error {
	return repo.UpdateUser(user)
}

func WrapJson(name string, item interface{}) map[string]interface{} {
	return map[string]interface{}{
		name: item,
	}
}
