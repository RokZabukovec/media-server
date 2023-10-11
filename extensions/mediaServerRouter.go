package extensions

import (
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"mediaserver/controllers"
	"net/http"
)

type MediaRouter struct{}

func (c *MediaRouter) ConfigureRoutes(r *chi.Mux) {
	r.Mount("/swagger", httpSwagger.WrapHandler)

	r.Get("/files/stream", func(w http.ResponseWriter, r *http.Request) {
		controllers.Stream(w, r)
	})

	r.Get("/files", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetFiles(w, r)
	})

	r.Post("/files", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateHslSegments(w, r)
	})
}
