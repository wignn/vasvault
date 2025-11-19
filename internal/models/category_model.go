package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string `gorm:"not null;uniqueIndex:idx_user_category" json:"name"`
	Color  string `gorm:"default:'#3B82F6'" json:"color"` // hex color
	UserID uint   `gorm:"not null;uniqueIndex:idx_user_category" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Files  []File `gorm:"many2many:file_categories;" json:"files,omitempty"`
}
