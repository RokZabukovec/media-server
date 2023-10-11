package responses

import (
	"encoding/json"
	"github.com/charmbracelet/log"
	"net/http"
)

type FileUploadedResponse struct {
	Message       string
	ContentLength int64
	Filename      string
}

func NewFileUploadedResponse(w http.ResponseWriter, message string, contentLength int64, filename string) {
	fileUploadedResponse := &FileUploadedResponse{Message: message, ContentLength: contentLength, Filename: filename}
	response, err := json.Marshal(fileUploadedResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(response)

	if err != nil {
		log.Errorf("Could not write to the response: %s", err)
	}
}
