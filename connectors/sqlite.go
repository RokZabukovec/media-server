package connectors

import (
	"github.com/charmbracelet/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mediaserver/services"
	"sync"
	"time"
)

var (
	db     *gorm.DB
	dbLock sync.Mutex
)

func GetDatabase(dbName string) *gorm.DB {
	dbLock.Lock()
	defer dbLock.Unlock()

	if db != nil {
		return db
	}

	var err error
	for i := 0; i < 3; i++ {
		dbFilepath, _ := services.GetDatabaseFilepath(dbName)
		db, err = gorm.Open(sqlite.Open(dbFilepath), &gorm.Config{})
		if err == nil {
			log.Printf("Connected to the database in %d try 💥", i)
			return db
		}
		log.Printf("Failed to connect to database: %v", err)
		time.Sleep(1 * time.Second)
	}

	return nil
}
