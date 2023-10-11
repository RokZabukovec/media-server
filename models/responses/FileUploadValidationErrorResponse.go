package responses

import (
	"encoding/json"
	"github.com/charmbracelet/log"
	"net/http"
)

type FileUploadValidationErrorResponse struct {
	AcceptedTypes []string
	Message       string
	ContentType   string
}

func NewFileUploadValidationErrorResponse(w http.ResponseWriter, acceptedTypes []string, message string, ContentType *string) {
	validationError := &FileUploadValidationErrorResponse{
		AcceptedTypes: acceptedTypes,
		Message:       message,
		ContentType:   *ContentType,
	}
	response, err := json.Marshal(validationError)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	_, err = w.Write(response)

	if err != nil {
		log.Errorf("Could not write to the response: %s", err)
	}

	return
}
