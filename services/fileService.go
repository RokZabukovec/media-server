package services

import (
	"fmt"
	"github.com/charmbracelet/log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DirInfo struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

func GetFolders(path string) []DirInfo {
	var filesInDir []DirInfo

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			info, err := file.Info()

			if err != nil {
				log.Errorf("Could not get file info for %s", file.Name())
				continue
			}

			dirInfo := DirInfo{
				Name:    info.Name(),
				Size:    info.Size(),
				Mode:    info.Mode(),
				ModTime: info.ModTime(),
				IsDir:   info.IsDir(),
			}

			filesInDir = append(filesInDir, dirInfo)
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

func CreateNewSegmentFile(segmentsFolderPath, segmentPrefix string, segmentIndex int) (*os.File, error) {
	segmentFileName := fmt.Sprintf("%s-%04d.ts", strings.TrimSuffix(segmentPrefix, filepath.Ext(segmentPrefix)), segmentIndex)
	segmentFilePath := filepath.Join(segmentsFolderPath, segmentFileName)

	return os.Create(segmentFilePath)
}

func CreateHlsFolder(filename string) (string, error) {
	dirPath, _ := GetMediaServerBaseDirectory()
	folder := filepath.Join(dirPath, RemoveExtension(filename))

	err := os.MkdirAll(folder, 0700)

	if err != nil {
		return "", err
	}

	return folder, nil
}
