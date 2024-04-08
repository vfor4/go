package dto

type User struct {
	Id       int    `json:"-"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type Profile struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

type Follower struct {
	UserId     string
	FollowerId string
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
