package repositories

import (
	"errors"
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
	"mediaserver/connectors"
	"mediaserver/models/database"
)

func CreateCategory(name string, parent *uint) error {
	category := database.Category{
		Name:     name,
		ParentID: parent,
	}

	db := connectors.GetDatabase()
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

	db := connectors.GetDatabase()

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

func GetCategory(id uint) (*database.Category, error) {
	var category database.Category

	db := connectors.GetDatabase()

	var result = db.Preload("Videos").Where("id = ? AND deleted_at = NULL", id).Find(&category)

	if result.Error != nil {
		log.Error("Error retrieving category", "error", result.Error)
		return nil, result.Error
	}

	log.Info("Successfully retrieved category ✅", "category", category.Name)

	return &category, nil
}

func DeleteCategory(id uint) error {
	if id == 0 {
		return errors.New("invalid ID")
	}

	db := connectors.GetDatabase()

	result := db.Delete(&database.Category{}, id)
	if result.Error != nil {
		log.Error("Error deleting a category", "error", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no record found to delete")
	}

	log.Info("Successfully deleted a category", "id", id)

	return nil
}
