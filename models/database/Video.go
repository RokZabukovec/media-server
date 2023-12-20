package database

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	CategoryID uint
	Category   Category
	Name       string
	Path       string
}
