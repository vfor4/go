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
		body, err := toResponse(r.PathValue("username"))
		if err != nil {
			log.Printf("Profile handler - cannot build profile body, %v", err)
		}
		w.Write(body)
	default:
		w.WriteHeader(500)
	}
}

func followerHandler(w http.ResponseWriter, r *http.Request) {
	u := r.PathValue("username")
	switch m := r.Method; m {
	case "POST":
		err := service.Follow(u)
		if err != nil {
			log.Printf("Profile handler - cannot build profile body, %v", err)
			w.WriteHeader(500)
			return
		}
		body, err := toResponse(u)
		if err != nil {
			log.Printf("Profile handler - cannot build profile body, %v", err)
			w.WriteHeader(500)
			return
		}
		w.Write(body)
	case "DELETE":
		err := service.Unfollow(u)
		if err != nil {
			log.Printf("Profile handler - cannot build profile body, %v", err)
			w.WriteHeader(500)
			return
		}
		body, err := toResponse(u)
		if err != nil {
			log.Printf("Profile handler - cannot build profile body, %v", err)
			w.WriteHeader(500)
			return
		}
		w.Write(body)
	default:
		w.WriteHeader(405)
	}
}

func toResponse(u string) ([]byte, error) {
	body, err := json.Marshal(service.WrapJson("profile", service.GetProfile(u)))
	if err != nil {
		return nil, err
	}
	return body, nil
}
