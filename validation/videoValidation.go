package validation

import (
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
	mimeType, err := GetFileMimeType(*file)
	if err != nil {
		isVideo = false
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
