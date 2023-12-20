package requests

type CategoryRequest struct {
	Name   string `json:"name"`
	Parent uint   `json:"parent"`
}
