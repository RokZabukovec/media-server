package connectors

import (
	"github.com/charmbracelet/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mediaserver/configuration"
	"mediaserver/services"
	"sync"
	"time"
)

var (
	db     *gorm.DB
	dbLock sync.Mutex
)

func GetDatabase() *gorm.DB {
	dbLock.Lock()
	defer dbLock.Unlock()

	if db != nil {
		return db
	}

	var err error
	for i := 0; i < 3; i++ {
		dbFilepath, _ := services.GetDatabaseFilepath(configuration.AppName)
		db, err = gorm.Open(sqlite.Open(dbFilepath), &gorm.Config{})
		if err == nil {
			log.Info("Connected to the database 💥", "retry", i)
			return db
		}
		log.Error("Failed to connect to database", "error", err)
		time.Sleep(1 * time.Second)
	}

	return nil
}
