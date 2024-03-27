package main

import (
	// "example/data-access/controller"
	"example/data-access/service"
	"fmt"
)

func main() {
	// controller.GetAccount()
	fmt.Print(service.CreateToken())
}
