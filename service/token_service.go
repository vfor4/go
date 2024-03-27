package service

import (
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken() string {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte("key")
	t = jwt.New(jwt.SigningMethodHS256)
	s, _ = t.SignedString(key)
	return s
}
