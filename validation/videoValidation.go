package validation

import (
	"io"
	"mime/multipart"
	"net/http"
)

const (
	VideoMP4       = "video/mp4"
	VideoMPEG      = "video/mpeg"
	VideoQuicktime = "video/quicktime"
	VideoMSVideo   = "video/x-msvideo"
	VideoFLV       = "video/x-flv"
	VideoWebM      = "video/webm"
)

func GetAcceptedMIMITypes() []string {
	return []string{VideoMP4, VideoMPEG, VideoQuicktime, VideoMSVideo, VideoFLV, VideoWebM}
}

func IsVideo(file *multipart.File) bool {
	isVideo := false

	// Store the current position of the file
	currentPosition, _ := (*file).Seek(0, io.SeekCurrent)

	mimeType, err := GetFileMimeType(*file)
	if err != nil {
		return false // directly return false if there's an error
	}

	// Reset the file's position back to where it was
	_, err = (*file).Seek(currentPosition, io.SeekStart)
	if err != nil {
		return false // handle error in seeking back, though this shouldn't generally happen
	}

	var AcceptedMIME = []string{VideoMP4, VideoMPEG, VideoQuicktime, VideoMSVideo, VideoFLV, VideoWebM}
	for _, format := range AcceptedMIME {
		if mimeType == format {
			return true
		}
	}

	return isVideo
}

func GetFileMimeType(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	mimeType := http.DetectContentType(buffer)
	return mimeType, nil
}
