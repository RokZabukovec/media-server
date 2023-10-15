package main

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"mediaserver/extensions"
	"mediaserver/services"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGTERM)

	app := NewApplication(8080, "_media-server")

	app.Run()
}

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
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5500"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		Debug:          false,
	})

	router.Use(c.Handler)

	mediaRouter := &extensions.MediaRouter{}

	mediaRouter.ConfigureRoutes(router)

	http.Handle("/", router)
	address := ":" + strconv.Itoa(app.port)

	log.Fatal(http.ListenAndServe(address, nil))
}
