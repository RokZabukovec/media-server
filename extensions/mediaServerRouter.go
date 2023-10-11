package extensions

import (
	"github.com/go-chi/chi"
	"mediaserver/controllers"
	"net/http"
)

type MediaRouter struct{}

func (c *MediaRouter) ConfigureRoutes(r *chi.Mux) {
	r.Get("/files", func(w http.ResponseWriter, r *http.Request) { controllers.GetFiles(w) })
	r.Post("/files", func(w http.ResponseWriter, r *http.Request) { controllers.UploadFile(w, r) })
}
