package main

import (
	// "rworld/controller"
	"fmt"
	"rworld/service"
)

// "example/data-access/controller"
// "example/data-access/service"
// "fmt"

type Message struct {
	Name string
	body string
	time int64
}

func main() {
	// controller.GetAccount()
	fmt.Print(service.GenerateJWT("lambda"))
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTE4OTg5MTMsImlzcyI6ImxvY2FsaG9zdC5jb20iLCJzdWIiOiJsYW1iZGEifQ.jcqRDRB7HyRBHLI5c1Ostpg4qJ64WVqjXKfE9gFUmLI"
	fmt.Println(service.ParseJWT(token))
}
