package services

import (
	"github.com/charmbracelet/log"
	"github.com/hashicorp/mdns"
	"os"
	"os/signal"
	"syscall"
)

type BroadcastService struct {
	port int
	name string
}

func NewBroadcastService(port int, name string) *BroadcastService {
	return &BroadcastService{
		port: port,
		name: name,
	}
}

func (s *BroadcastService) Broadcast() {

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGTERM)

	info := []string{"Media Server"}
	service, err := mdns.NewMDNSService("Media Server", "_http._tcp.", "", "", 8080, nil, info)

	log.Info("Started service instance Media Server of type _http._tcp on port 8080")
	if err != nil {
		log.Fatalf("Failed to create mDNS service: %v", err)
	}

	server, err := mdns.NewServer(&mdns.Config{Zone: service})

	if err != nil {
		log.Fatalf("Failed to create mDNS server: %v", err)
	}

	log.Info("Service broadcasting started")

	<-terminationChan

	log.Info("Service broadcasting stopping")

	server.Shutdown()

	log.Info("Service broadcasting stopped")

	os.Exit(0)
}
