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
	const API_URL = "api"
	http.HandleFunc(fmt.Sprintf("/%s/user", API_URL), handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("GET")
		w.Header().Set("Content-type", "application/json")
		accountJson, _ := json.Marshal(service.GetAccount(r.URL.Query().Get("id")))
		w.Write(accountJson)
	case "POST":
		fmt.Println("POST")
		loginInfo, _ := postMethodHandler(r)
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

func postMethodHandler(r *http.Request) (*dto.LoginInfo, error) {
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
