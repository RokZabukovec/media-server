package responses

type FileUploadedResponse struct {
	Message       string
	ContentLength int64
	Filename      string
}

func NewFileUploadedResponse(message string, contentLength int64, filename string) *FileUploadedResponse {
	return &FileUploadedResponse{Message: message, ContentLength: contentLength, Filename: filename}
}
