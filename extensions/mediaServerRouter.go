package extensions

import (
	"github.com/go-chi/chi"
	"mediaserver/controllers"
	"net/http"
)

type MediaRouter struct{}

func (c *MediaRouter) ConfigureRoutes(r *chi.Mux) {
	r.Get("/thumbnail/{folder}", func(w http.ResponseWriter, r *http.Request) { controllers.GetThumbnail(w, r) })
	r.Get("/stream/{folder}/manifest.m3u8", func(w http.ResponseWriter, r *http.Request) { controllers.Playlist(w, r) })
	r.Get("/stream/{folder}/{segment}", func(w http.ResponseWriter, r *http.Request) { controllers.Stream(w, r) })
	r.Get("/files", func(w http.ResponseWriter, r *http.Request) { controllers.GetFiles(w) })
	r.Post("/files", func(w http.ResponseWriter, r *http.Request) { controllers.UploadFile(w, r) })
}
