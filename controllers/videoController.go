package controllers

import (
	"encoding/json"
	"github.com/charmbracelet/log"
	"mediaserver/models/requests"
	"mediaserver/repositories"
	"net/http"
	"strconv"
)

func CreateVideo(w http.ResponseWriter, r *http.Request) {
	var req requests.VideoCreateRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createError := repositories.CreateVideo(req.Name, &req.CategoryId, req.Path)

	if createError != nil {
		http.Error(w, createError.Error(), http.StatusInternalServerError)
		return
	}
}

func GetVideo(w http.ResponseWriter, r *http.Request) {
	var req requests.GetVideoRequest

	queryValues := r.URL.Query()
	idStr := queryValues.Get("id")

	if len(idStr) == 0 {
		req.Id = 0
	} else {
		parseInt, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			log.Error("Cannot parse query parameter", "parent", idStr)
		}

		req.Id = uint(parseInt)
	}

	videos, err := repositories.GetVideo(&req.Id)
	if err != nil {
		http.Error(w, "Error fetching categories", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(videos)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func DeleteVideo(w http.ResponseWriter, r *http.Request) {
	var req requests.GetVideoRequest

	queryValues := r.URL.Query()
	idStr := queryValues.Get("id")

	if len(idStr) == 0 {
		req.Id = 0
	} else {
		parseInt, err := strconv.ParseInt(idStr, 10, 64)

		if err != nil {
			log.Error("Cannot parse query parameter", "parent", idStr)
		}

		req.Id = uint(parseInt)
	}

	err := repositories.DeleteVideo(req.Id)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
