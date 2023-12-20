package main

import (
	"mediaserver/configuration"
	"mediaserver/migrations"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGTERM)

	migrations.Migrate(configuration.AppName)
	app := NewApplication(8888, configuration.AppName)

	app.Run()
}
