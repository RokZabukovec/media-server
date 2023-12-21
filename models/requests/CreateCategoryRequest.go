package requests

type CategoryCreateRequest struct {
	Name   string `json:"name"`
	Parent uint   `json:"parent"`
}

type CategoryGetAllRequest struct {
	Parent uint `json:"parent"`
}
