package controllers

import (
	"encoding/json"
	"mediaserver/models/requests"
	"mediaserver/repositories"
	"net/http"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req requests.CategoryRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createError := repositories.CreateCategory(req.Name, &req.Parent)

	if createError != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
