package services

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetFiles(path string) []string {
	filesInDir := []string{}

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filesInDir = append(filesInDir, file.Name())
		}
	}

	return filesInDir
}

func GetMediaServerBaseDirectory() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	directoryName := "MediaServer"

	return filepath.Join(homeDir, directoryName), nil
}

func RemoveExtension(fileName string) string {
	baseName := filepath.Base(fileName)
	fileExtension := filepath.Ext(baseName)
	fileNameWithoutExtension := strings.TrimSuffix(baseName, fileExtension)

	return fileNameWithoutExtension
}
