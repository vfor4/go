package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"rworld/service"
)

func profileHandler(w http.ResponseWriter, r *http.Request) {
	switch m := r.Method; m {
	case "GET":
		body, err := json.Marshal(service.WrapJson("profile", service.GetProfile(r.PathValue("username"))))
		if err != nil {
			log.Printf("Profile handler - cannot build profile body, %v", err)
		}
		w.Write(body)
	default:
		w.WriteHeader(500)
	}
}
