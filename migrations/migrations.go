package migrations

import (
	"fmt"
	"github.com/charmbracelet/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mediaserver/models/database"
	"mediaserver/services"
)

func Migrate(dbName string) {
	dbFilepath, _ := services.GetDatabaseFilepath(dbName)
	db, err := gorm.Open(sqlite.Open(dbFilepath), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Print(fmt.Sprintf("Created %s database ✅", dbFilepath))

	migrateError := db.AutoMigrate(&database.Category{}, &database.Video{})
	if migrateError != nil {
		panic("Database migration failed")
	}

	log.Print(fmt.Println("Migration succeeded ✅"))
}
