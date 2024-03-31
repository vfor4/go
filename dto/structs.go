package dto

type User struct {
	Email    string
	Token    string
	Username string
	Bio      string
	image    string
}

type LoginInfo struct {
	Email    string
	Password string
}
