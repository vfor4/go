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
	log.Printf("Fiter for url: %s", r.URL.EscapedPath())
	if inWhiteList := strings.HasSuffix(r.URL.EscapedPath(), "users/login"); inWhiteList {
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
