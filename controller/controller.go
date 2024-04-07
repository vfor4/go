package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/login", loginHandler)
	mux.HandleFunc("/api/users", signupHandler)
	mux.HandleFunc("/api/user", userHandler)
	mux.HandleFunc("/api/profiles/{username}", profileHandler)

	filter := NewFilter(mux)

	http.ListenAndServe(":8080", filter)
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
