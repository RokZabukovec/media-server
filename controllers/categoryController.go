package controllers

import (
	"encoding/json"
	"github.com/charmbracelet/log"
	"mediaserver/models/requests"
	"mediaserver/repositories"
	"net/http"
	"strconv"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req requests.CategoryCreateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createError := repositories.CreateCategory(req.Name, &req.Parent)

	if createError != nil {
		http.Error(w, createError.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	var req requests.CategoryGetAllRequest

	queryValues := r.URL.Query()
	parentStr := queryValues.Get("parent")

	if len(parentStr) == 0 {
		req.Parent = 0
	} else {
		parentInt, err := strconv.ParseInt(parentStr, 10, 64)

		if err != nil {
			log.Error("Cannot parse query parameter", "parent", parentStr)
		}

		req.Parent = uint(parentInt)
	}

	categories, err := repositories.GetCategories(&req.Parent)
	if err != nil {
		http.Error(w, "Error fetching categories", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
