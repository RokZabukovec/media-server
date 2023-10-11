package controllers

import (
	"encoding/json"
	"github.com/charmbracelet/log"
	"mediaserver/models/responses"
	"mediaserver/services"
	"mediaserver/validation"
	"net/http"
)

const (
	MaxFileSize = 1 << 30 // 1 gigabyte (GB) in bytes
)

func GetFiles(w http.ResponseWriter) {
	dirPath, _ := services.GetMediaServerBaseDirectory()

	files := services.GetFolders(dirPath)

	jsonString, err := json.Marshal(files)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonString)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(MaxFileSize)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")

	if validation.IsVideo(&file) == false {
		mimeType, _ := validation.GetFileMimeType(file)
		responses.NewFileUploadValidationErrorResponse(w, validation.GetAcceptedMIMITypes(), "The file is not in the correct format", &mimeType)
		return
	}

	bytesUploaded, segmentError := services.CreateSegmentedFiles(handler.Filename, file)

	nbMB := float64(bytesUploaded) / (1024 * 1024)
	log.Printf("Successfully uploaded %.2f MB\n", nbMB)

	if segmentError != nil || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responses.NewFileUploadedResponse(w, "File uploaded successfully", bytesUploaded, handler.Filename)
}
