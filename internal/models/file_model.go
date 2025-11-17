package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Filename   string    `gorm:"not null" json:"filename"`
	Filepath   string    `gorm:"not null" json:"filepath"`
	Mimetype   string    `gorm:"not null" json:"mimetype"`
	Size       int64     `gorm:"not null" json:"size"`
	UploadedAt time.Time `gorm:"autoCreateTime" json:"uploaded_at"`
	UserID     uint      `gorm:"not null" json:"user_id"`
	User       User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
