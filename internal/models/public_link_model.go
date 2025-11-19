package models

import (
	"time"

	"gorm.io/gorm"
)

type PublicLink struct {
	gorm.Model
	FileID uint `gorm:"not null;index" json:"file_id"`
	File   File `gorm:"foreignKey:FileID" json:"file,omitempty"`

	Token      string `gorm:"uniqueIndex;not null" json:"token"`         // random token for URL
	Permission string `gorm:"not null;default:'view'" json:"permission"` // view, download

	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	CreatedBy uint       `gorm:"not null" json:"created_by"`
	Creator   User       `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`

	AccessCount int  `gorm:"default:0" json:"access_count"`
	IsActive    bool `gorm:"default:true" json:"is_active"`
}
