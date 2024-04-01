package controller

import (
	"fmt"
	"log"
	"net/http"
	"rworld/service"
	"strings"
)

type Filter struct {
	handler http.Handler
}

func (f *Filter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if user := service.ParseJWT(strings.Split(r.Header.Get("Authorization"), " ")[1]); user != "" {
		f.handler.ServeHTTP(w, r)
	} else {
		log.Println("Token is invalid")
		fmt.Fprint(w, "Somehow the token is invalid, try hard next time!")
	}
}

func NewFilter(httpHandler http.Handler) *Filter {
	return &Filter{httpHandler}
}
