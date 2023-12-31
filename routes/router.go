package routes

import (
	"github.com/go-chi/chi"
	"mediaserver/controllers"
	"net/http"
)

type MediaRouter struct{}

func (c *MediaRouter) ConfigureRoutes(r *chi.Mux) {
	c.configureSpaRoutes(r)
	c.configureApiRoutes(r)
}

func (c *MediaRouter) configureSpaRoutes(r *chi.Mux) {
	r.Handle("/*", controllers.Index("./ui/hub-ui/dist/"))
	r.NotFound(controllers.NotFound)
}

func (c *MediaRouter) configureApiRoutes(r *chi.Mux) {
	r.Get("/api/thumbnail/{folder}", func(w http.ResponseWriter, r *http.Request) { controllers.GetThumbnail(w, r) })
	r.Get("/api/stream/{folder}/manifest.m3u8", func(w http.ResponseWriter, r *http.Request) { controllers.Playlist(w, r) })
	r.Get("/api/stream/{folder}/{segment}", func(w http.ResponseWriter, r *http.Request) { controllers.Stream(w, r) })
	r.Get("/api/files", func(w http.ResponseWriter, r *http.Request) { controllers.GetFiles(w) })
	r.Post("/api/files", func(w http.ResponseWriter, r *http.Request) { controllers.UploadFile(w, r) })

	r.Get("/api/video", func(w http.ResponseWriter, r *http.Request) { controllers.GetVideo(w, r) })
	r.Get("/api/videos", func(w http.ResponseWriter, r *http.Request) { controllers.GetVideo(w, r) })
	r.Post("/api/video", func(w http.ResponseWriter, r *http.Request) { controllers.CreateVideo(w, r) })
	r.Delete("/api/video", func(w http.ResponseWriter, r *http.Request) { controllers.DeleteVideo(w, r) })

	r.Get("/api/category", func(w http.ResponseWriter, r *http.Request) { controllers.GetCategory(w, r) })
	r.Get("/api/categories", func(w http.ResponseWriter, r *http.Request) { controllers.GetAllCategories(w, r) })
	r.Post("/api/categories", func(w http.ResponseWriter, r *http.Request) { controllers.CreateCategory(w, r) })
	r.Delete("/api/category", func(w http.ResponseWriter, r *http.Request) { controllers.DeleteCategory(w, r) })
}
