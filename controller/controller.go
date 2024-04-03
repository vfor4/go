package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rworld/dto"
	"rworld/service"
)

func GetAccount() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/login", loginHandler)
	mux.HandleFunc("/api/user", handler)

	filter := NewFilter(mux)

	http.ListenAndServe(":8080", filter)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		loginInfo, _ := extractBody(r)
		if service.Loggedin(*loginInfo) {
			user := service.GetAccount(loginInfo.Email)
			user.Token = service.GenerateJWT(loginInfo.Email)
			userJson, _ := json.Marshal(user)
			w.Write(userJson)
		} else {
			fmt.Fprint(w, "Ops, maybe next time")
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-type", "application/json")
		accountJson, _ := json.Marshal(service.GetAccount(service.GetUser()))
		w.Write(accountJson)
	} else {
		w.WriteHeader(405)
	}
}

func extractBody(r *http.Request) (*dto.LoginInfo, error) {
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
	user := (i.(map[string]interface{})["user"]).(map[string]interface{})
	var loginInfo dto.LoginInfo
	loginInfo.Email = user["email"].(string)
	loginInfo.Password = user["password"].(string)
	return &loginInfo, nil
}
