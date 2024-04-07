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

var whiteList = [2]string{"/api/users/login", "/api/users"}

func inWhiteList(url string) bool {
	for _, v := range whiteList {
		if url == v {
			return true
		}
	}
	return false
}

func (f *Filter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Fiter for url: %s - %s", r.Method, r.URL.EscapedPath())
	if inWhiteList(r.URL.EscapedPath()) {
		f.handler.ServeHTTP(w, r)
	} else {
		if service.TokenValid(strings.Split(r.Header.Get("Authorization"), " ")[1]) {
			f.handler.ServeHTTP(w, r)
		} else {
			fmt.Fprint(w, "Somehow the token is invalid, try harder next time!")
		}
	}
}

func NewFilter(httpHandler http.Handler) *Filter {
	return &Filter{httpHandler}
}
