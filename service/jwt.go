package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var email string

func GetEmail() string {
	return email
}

func GenerateJWT(email string) string {
	// key, err := os.ReadFile("/home/vu/coding/go/keys/private_key.pem")
	key := []byte("happy new year")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": email,
			"iss": "localhost.com",
			"exp": time.Now().Add(time.Minute * 30).Unix(),
		})
	s, err := t.SignedString(key)
	if err != nil {
		fmt.Printf("cannot generate jwt %v", err)
	}
	return s
}

func TokenValid(bearerToken string) bool {
	token, err := jwt.Parse(bearerToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		key := []byte("happy new year")
		return key, nil
	}, jwt.WithExpirationRequired(), jwt.WithIssuer("localhost.com"))
	if err != nil {
		fmt.Printf("cannot parse jwt, %v", err)
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		email = claims["sub"].(string)
		return true
	}
	return false
}

// func BytesToPrivateKey(key []byte) *rsa.PrivateKey {
// 	pubInterface, err := x509.ParsePKCS8PrivateKey(key)
// 	if err != nil {
// 		fmt.Printf("ParsePKCS8PrivateKey error, %v", err)
// 	}
// 	priv, ok := pubInterface.(*rsa.PrivateKey)
// 	if !ok {
// 		fmt.Printf("convert to PrivateKey error, %v", err)
// 	}
// 	return priv
// }
