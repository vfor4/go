package main

import (
	"rworld/controller"
)

type Message struct {
	Name string
	body string
	time int64
}

func main() {
	controller.GetAccount()
	// // fmt.Print(service.GenerateJWT("lambda"))
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTE4OTg4MjYsImlzcyI6ImxvY2FsaG9zdC5jb20iLCJzdWIiOiJsYW1iZGEifQ.g4ccDlaJ-n3sp3PIZR0TPhwOrWRaT63T2Pfp-klwi64"
	// fmt.Println(service.ParseJWT(token))
}
