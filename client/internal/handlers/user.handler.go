package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handlers) GetUser(w http.ResponseWriter, r *http.Request) {

	userId := mux.Vars(r)["userId"]

	if userId == "" {
		http.Error(w, "Missing userId", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(userId)

	if err != nil {
		http.Error(w, "Invalid userId", http.StatusInternalServerError)
		return
	}

	user, err := h.services.UserService.GetUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
