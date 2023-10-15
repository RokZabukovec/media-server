package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi"
	"mediaserver/models/responses"
	"mediaserver/services"
	"mediaserver/validation"
	"net/http"
	"os"
	"path"
)

const (
	MaxFileSize = 1 << 30 // 1 gigabyte (GB) in bytes
)

func Stream(w http.ResponseWriter, r *http.Request) {
	folderName := chi.URLParam(r, "folder")
	segmentName := chi.URLParam(r, "segment")

	baseDir, _ := services.GetMediaServerBaseDirectory()
	segmentFile := path.Join(baseDir, folderName, segmentName)

	_, err := os.Stat(segmentFile)
	if os.IsNotExist(err) {
		http.Error(w, "Segment not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")

	segmentContent, err := os.ReadFile(segmentFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading segment file: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusPartialContent)
	w.Write(segmentContent)
}

func Playlist(w http.ResponseWriter, r *http.Request) {
	folderName := chi.URLParam(r, "folder")

	baseDir, _ := services.GetMediaServerBaseDirectory()
	manifestDir := path.Join(baseDir, folderName)
	manifestFile := "manifest.m3u8"
	manifestPath := path.Join(manifestDir, manifestFile)

	manifestContent, err := os.ReadFile(manifestPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading manifest file: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
	w.WriteHeader(http.StatusPartialContent)

	w.Write(manifestContent)
}

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
	if err != nil {
		http.Error(w, "Failed to retrieve file from request", http.StatusInternalServerError)
		return
	}

	if validation.IsVideo(&file) == false {
		mimeType, _ := validation.GetFileMimeType(file)
		responses.NewFileUploadValidationErrorResponse(w, validation.GetAcceptedMIMITypes(), "The file is not in the correct format", &mimeType)
		file.Close()
		return
	}

	basePath, err := services.GetMediaServerBaseDirectory()
	if err != nil {
		http.Error(w, "Failed to get media server directory", http.StatusInternalServerError)
		file.Close()
		return
	}

	videoFolder := path.Join(basePath, services.RemoveExtension(handler.Filename))

	segmentError := services.CreateHLSFilesFromAPIRequest(file, videoFolder, handler.Filename)
	if segmentError != nil {
		log.Error(segmentError)
		http.Error(w, "Internal server error while processing the file", http.StatusInternalServerError)
		file.Close()
		return
	}

	file.Close()
	responses.NewFileUploadedResponse(w, "File uploaded successfully", handler.Filename)
}
