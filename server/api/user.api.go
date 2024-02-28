package api

import (
	"encoding/json"
	"github.com/tasnimzotder/tchat/server/utils"
	"log"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *ServerAPI) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// check if request is POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get user from request body
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uuid, _ := utils.GenerateUUID()
	user.ID = uuid

	log.Printf("User created: %v", user)

	//	todo: create a certificate for the user for end-to-end encryption

	//	return user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
