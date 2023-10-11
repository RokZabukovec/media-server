package main

import (
	"github.com/charmbracelet/log"
	"github.com/go-chi/chi"
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

	r := chi.NewRouter()

	mediaRouter := &extensions.MediaRouter{}
	mediaRouter.ConfigureRoutes(r)

	http.Handle("/", r)
	address := ":" + strconv.Itoa(app.port)

	log.Fatal(http.ListenAndServe(address, nil))
}
