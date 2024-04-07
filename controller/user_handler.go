package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rworld/dto"
	"rworld/service"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := extractBody(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		user := dto.SignUpUser{}
		service.MapTo(&user, body)
		err = service.SignUp(user)
		if err != nil {
			fmt.Fprint(w, "Connot register use a")
		} else {
			user := service.GetUserByEmail(user.Email)
			userJson, _ := json.Marshal(service.WrapJson("user", user))
			w.Write(userJson)
		}
	} else {
		w.WriteHeader(405)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case "GET":
		accountJson, _ := json.Marshal(service.WrapJson("user", service.GetUserByUsername(service.GetSubject())))
		w.Write(accountJson)
	case "PUT":
		body, err := extractBody(r)
		if err != nil {
			log.Print("userHander - cannot parse request body")
			w.WriteHeader(500)
		}
		user := dto.User{}
		service.MapTo(&user, body)
		err = service.UpdateUser(user)
		if err != nil {
			log.Print("userHander - cannot update user")
		}
		accountJson, _ := json.Marshal(service.WrapJson("user", service.GetUserByUsername(service.GetSubject())))
		w.Write(accountJson)
	default:
		w.WriteHeader(405)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := extractBody(r)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		loginInfo := dto.LoginInfo{}
		service.MapTo(&loginInfo, body)

		if service.Loggedin(loginInfo) {
			user := service.GetUserByEmail(loginInfo.Email)
			user.Token = service.GenerateJWT(user.Username)
			userJson, _ := json.Marshal(service.WrapJson("user", user))
			w.Write(userJson)
		} else {
			fmt.Fprint(w, "Ops, wrong username or password 🥹")
		}
	} else {
		w.WriteHeader(405)
	}
}
