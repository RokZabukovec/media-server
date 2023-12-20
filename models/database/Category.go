package database

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string
	Videos   []Video
	ParentID *uint
	Parent   *Category  `gorm:"foreignKey:ParentID"`
	Children []Category `gorm:"foreignKey:ParentID"`
}
