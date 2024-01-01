package repositories

import (
	"errors"
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
	"mediaserver/connectors"
	"mediaserver/models/database"
)

func CreateVideo(name string, category *uint, path string) error {
	video := database.Video{
		Name:       name,
		CategoryID: *category,
		Path:       path,
	}

	db := connectors.GetDatabase()
	result := db.Create(&video)
	if result.Error != nil {
		log.Error("Error creating video: %v\n ❌", result.Error)

		return result.Error
	}

	log.Info("Video created successfully ✅", "video", name)

	return nil
}

func GetVideo(id *uint) (*database.Video, error) {
	var video database.Video

	db := connectors.GetDatabase()

	var result *gorm.DB
	if id == nil {
		result = db.Where("id = 0").First(&video)
	} else {
		result = db.Where("id = ? AND deleted_at = NULL", *id).First(&video)
	}

	if result.Error != nil {
		log.Error("Error retrieving videos", "error", result.Error)

		return nil, result.Error
	}

	log.Info("Successfully retrieved videos ✅", "video", video.Name)

	return &video, nil
}

func DeleteVideo(id uint) error {
	if id == 0 {
		return errors.New("invalid ID")
	}

	db := connectors.GetDatabase()

	result := db.Delete(&database.Video{}, id)
	if result.Error != nil {
		log.Error("Error deleting a video", "error", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no record found")
	}

	log.Info("Successfully deleted a video", "id", id)

	return nil
}
