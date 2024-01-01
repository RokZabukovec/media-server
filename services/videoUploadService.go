package services

import (
	"errors"
	"mediaserver/validation"
	"mime/multipart"
	"path"
)

type VideoUploadService struct {
	// Add any dependencies here, like configuration or repositories
}

func NewVideoUploadService() *VideoUploadService {
	return &VideoUploadService{
		// Initialize dependencies
	}
}

func (s *VideoUploadService) UploadVideo(file multipart.File, handler *multipart.FileHeader) (string, error) {
	if validation.IsVideo(&file) == false {
		mimeType, _ := validation.GetFileMimeType(file)
		return "", errors.New("invalid file format: " + mimeType)
	}

	basePath, err := GetMediaServerBaseDirectory()
	if err != nil {
		return "", err
	}

	videoFolder := path.Join(basePath, RemoveExtension(handler.Filename))

	err = CreateHLSFilesFromAPIRequest(file, videoFolder, handler.Filename)
	if err != nil {
		return "", err
	}

	return handler.Filename, nil
}
