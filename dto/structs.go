package dto

type User struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type LoginInfo struct {
	Email    string
	Password string
}

type SignUpUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
