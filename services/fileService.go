package services

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"mediaserver/configuration"
	"mediaserver/models"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetFolders(path string) []models.DirInfo {
	var filesInDir []models.DirInfo

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, folder := range files {
		if folder.IsDir() {
			info, err := folder.Info()

			if err != nil {
				log.Errorf("Could not get folder info for %s", folder.Name())
				continue
			}

			playlist, _ := GetMediaUrl(info.Name())
			thumbnail, _ := GetThumbnailUrl(info.Name())

			dirInfo := models.NewDirInfo(info.Name(), info.Size(), thumbnail, playlist)
			filesInDir = append(filesInDir, *dirInfo)
		}
	}

	return filesInDir
}

func GetThumbnailFilepath(mediaDir string) (string, error) {
	base, err := GetMediaServerBaseDirectory()
	if err != nil {
		return "", err
	}

	thumbnailPath := path.Join(base, mediaDir, configuration.ThumbnailName)
	if _, err := os.Stat(thumbnailPath); os.IsNotExist(err) {
		return "", errors.New("thumbnail file does not exist")
	}

	return thumbnailPath, nil
}

func GetThumbnailUrl(mediaDir string) (string, error) {
	base, err := GetMediaServerBaseDirectory()
	if err != nil {
		return "", err
	}

	thumbnailPath := path.Join(base, mediaDir, configuration.ThumbnailName)
	if _, err := os.Stat(thumbnailPath); os.IsNotExist(err) {
		return "", errors.New("thumbnail file does not exist")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("/thumbnail/%s", mediaDir), nil
}

func GetMediaUrl(mediaDir string) (string, error) {
	base, err := GetMediaServerBaseDirectory()
	if err != nil {
		return "", err
	}

	thumbnailPath := path.Join(base, mediaDir, configuration.PlaylistName)
	if _, err := os.Stat(thumbnailPath); os.IsNotExist(err) {
		return "", errors.New("playlist file does not exist")
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("/stream/%s/%s", mediaDir, configuration.PlaylistName), nil
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
	segmentFileName := fmt.Sprintf("%s%d.ts", strings.TrimSuffix(segmentPrefix, filepath.Ext(segmentPrefix)), segmentIndex)
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
