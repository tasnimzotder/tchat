package api

import (
	"encoding/json"
	"github.com/tasnimzotder/tchat/server/models"
	"github.com/tasnimzotder/tchat/server/utils"
	"net/http"
)

func (s *ServerAPI) createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.ErrorResponse(w, "method not allowed", nil)

		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ErrorResponse(w, "", err)

		return
	}

	uuid, _ := utils.GenerateUUID()
	user.ID = uuid

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
