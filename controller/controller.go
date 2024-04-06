package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rworld/dto"
	"rworld/service"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/login", loginHandler)
	mux.HandleFunc("/api/users", signup)
	mux.HandleFunc("/api/user", getUser)

	filter := NewFilter(mux)

	http.ListenAndServe(":8080", filter)
}

func signup(w http.ResponseWriter, r *http.Request) {
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
			user := service.GetAccount(user.Email)
			userJson, _ := json.Marshal(user)
			w.Write(userJson)
		}
	} else {
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
			user := service.GetAccount(loginInfo.Email)
			user.Token = service.GenerateJWT(loginInfo.Email)
			userJson, _ := json.Marshal(user)
			w.Write(userJson)
		} else {
			fmt.Fprint(w, "Ops, wrong username or password 🥹")
		}
	} else {
		w.WriteHeader(405)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-type", "application/json")
		accountJson, _ := json.Marshal(service.GetAccount(service.GetEmail()))
		w.Write(accountJson)
	} else {
		w.WriteHeader(405)
	}
}

func extractBody(r *http.Request) (map[string]interface{}, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Body reading error: %v", err)
		return nil, err
	}
	defer r.Body.Close()

	var i interface{}
	err2 := json.Unmarshal(bodyBytes, &i)
	if err2 != nil {
		fmt.Printf("cannot parse request body: %v", err2)
	}
	body := (i.(map[string]interface{})["user"]).(map[string]interface{})
	return body, nil
}
