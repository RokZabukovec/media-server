package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/charmbracelet/log"
	"io"
	"mediaserver/models/responses"
	"mediaserver/services"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	MaxFileSize = 1 << 30     // 1 gigabyte (GB) in bytes
	SegmentSize = 1024 * 1024 // 5MB segment size
	BufferSize  = 1024        // 1MB buffer size
)

func Stream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonString, err := json.Marshal("Hello")

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

func GetFiles(w http.ResponseWriter, r *http.Request) {
	dirPath, _ := services.GetMediaServerBaseDirectory()

	files := services.GetFiles(dirPath)

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
	const MaxFileSize = 1 << 30 // 1 gigabyte (GB) in bytes
	err := r.ParseMultipartForm(MaxFileSize)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dirPath, _ := services.GetMediaServerBaseDirectory()
	err = os.MkdirAll(dirPath, 0700)
	if err != nil {
		http.Error(w, "Failed to create directory", http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(dirPath, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to create destination file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	buffer := make([]byte, 1024*1024) // 1MB buffer size
	var nbBytes int64
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}
		dst.Write(buffer[:n])
		nbBytes += int64(n)
	}

	nbMB := float64(nbBytes) / (1024 * 1024)
	log.Printf("Successfully uploaded %.2f MB\n", nbMB)

	w.WriteHeader(http.StatusOK)
	response := responses.NewFileUploadedResponse("File uploaded successfully", nbBytes, handler.Filename)
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
}

func CreateHslSegments(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(MaxFileSize)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	segmentsFolderPath, err := CreateHlsFolder(handler.Filename)
	if err != nil {
		http.Error(w, "Failed to create HLS folder", http.StatusInternalServerError)
		return
	}

	buffer := make([]byte, BufferSize)
	var nbBytes int64
	segmentIndex := 0
	segmentFile, err := createNewSegmentFile(segmentsFolderPath, handler.Filename, segmentIndex)
	if err != nil {
		http.Error(w, "Failed to create segment file", http.StatusInternalServerError)
		return
	}
	defer segmentFile.Close()

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}

		nbBytes += int64(n)

		if nbBytes >= int64(SegmentSize) {
			segmentIndex++
			segmentFile.Close()
			segmentFile, err = createNewSegmentFile(segmentsFolderPath, handler.Filename, segmentIndex)
			if err != nil {
				http.Error(w, "Failed to create segment file", http.StatusInternalServerError)
				return
			}

			defer segmentFile.Close()

			nbBytes = 0
		}

		_, writeErr := segmentFile.Write(buffer[:n])

		if writeErr != nil {
			http.Error(w, "Failed to write segment data", http.StatusInternalServerError)
			return
		}
	}

	nbMB := float64(nbBytes) / (1024 * 1024)
	log.Printf("Successfully uploaded %.2f MB\n", nbMB)

	w.WriteHeader(http.StatusOK)
	response := responses.NewFileUploadedResponse("File uploaded successfully", nbBytes, handler.Filename)
	jsonBytes, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonBytes)
}

func createNewSegmentFile(segmentsFolderPath, filename string, segmentIndex int) (*os.File, error) {
	segmentFileName := fmt.Sprintf("%s-%04d.ts", strings.TrimSuffix(filename, filepath.Ext(filename)), segmentIndex)
	segmentFilePath := filepath.Join(segmentsFolderPath, segmentFileName)
	return os.Create(segmentFilePath)
}

func CreateHlsFolder(filename string) (string, error) {

	dirPath, _ := services.GetMediaServerBaseDirectory()
	folder := filepath.Join(dirPath, services.RemoveExtension(filename))

	error := os.MkdirAll(folder, 0700)
	if error != nil {
		return "", error
	}

	return folder, nil
}
