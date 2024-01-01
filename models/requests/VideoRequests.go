package requests

type VideoCreateRequest struct {
	Name       string `json:"name"`
	CategoryId uint   `json:"category_id"`
	Path       string `json:"path"`
}

type GetVideoRequest struct {
	Id uint `json:"Id"`
}
