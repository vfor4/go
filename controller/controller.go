package controller

import (
	"encoding/json"
	"example/data-access/service"
	"fmt"
	"net/http"
)

func GetAccount() {
	const API_URL = "api"
	http.HandleFunc(fmt.Sprintf("/%s/user", API_URL), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		accountJson, _ := json.Marshal(service.GetAccount(r.URL.Query().Get("id")))
		w.Write(accountJson)
	})
	http.ListenAndServe(":8080", nil)
}
