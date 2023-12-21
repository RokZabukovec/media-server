package repositories

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
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
		log.Error("Error creating category: %v\n ❌", result.Error)

		return result.Error
	}

	log.Info("Category created successfully:%s ✅", name)

	return nil
}

func GetCategories(parentId *uint) ([]database.Category, error) {
	var categories []database.Category

	db := connectors.GetDatabase(configuration.AppName)

	var result *gorm.DB
	if parentId == nil {
		result = db.Where("parent_id = 0").Find(&categories)
	} else {
		result = db.Where("parent_id = ?", *parentId).Find(&categories)
	}

	if result.Error != nil {
		log.Error("Error retrieving categories", "error", result.Error)
		return nil, result.Error
	}

	log.Info("Successfully retrieved categories ✅", "categories", len(categories))

	return categories, nil
}
