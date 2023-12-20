package repositories

import (
	"log"
	"mediaserver/configuration"
	"mediaserver/connectors"
	"mediaserver/models/database"
)

func CreateCategory(name string, parent *uint) error {
	category := database.Category{
		Name:     name,
		ParentID: parent,
	}

	db := connectors.GetDatabase(configuration.AppName)
	result := db.Create(&category)
	if result.Error != nil {
		log.Printf("Error creating category: %v\n ❌", result.Error)

		return result.Error
	}

	log.Printf("Category created successfully:%s ✅", name)

	return nil
}
