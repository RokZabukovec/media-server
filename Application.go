package main

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"mediaserver/routes"
	"mediaserver/services"
	"net/http"
	"strconv"
)

type Application struct {
	port            int
	applicationName string
}

func NewApplication(port int, applicationName string) *Application {
	return &Application{
		port:            port,
		applicationName: applicationName,
	}
}

func (app *Application) Run() {
	service := services.NewBroadcastService(app.port, app.applicationName)
	go service.Broadcast()

	router := chi.NewRouter()

	// Before production evaluate these options
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		Debug:          false,
	})

	router.Use(c.Handler)

	mediaRouter := &routes.MediaRouter{}

	mediaRouter.ConfigureRoutes(router)

	http.Handle("/", router)
	address := ":" + strconv.Itoa(app.port)

	log.Fatal(http.ListenAndServe(address, nil))
}
