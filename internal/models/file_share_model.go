package models

import (
	"time"

	"gorm.io/gorm"
)

type FileShare struct {
	gorm.Model
	FileID uint `gorm:"not null;index" json:"file_id"`
	File   File `gorm:"foreignKey:FileID" json:"file,omitempty"`

	SharedByUserID uint `gorm:"not null;index" json:"shared_by_user_id"`
	SharedBy       User `gorm:"foreignKey:SharedByUserID" json:"shared_by,omitempty"`

	SharedWithUserID uint `gorm:"not null;index" json:"shared_with_user_id"`
	SharedWith       User `gorm:"foreignKey:SharedWithUserID" json:"shared_with,omitempty"`

	Permission string     `gorm:"not null;default:'view'" json:"permission"`
	SharedAt   time.Time  `gorm:"autoCreateTime" json:"shared_at"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
}

// Permission constants
const (
	PermissionView     = "view"
	PermissionEdit     = "edit"
	PermissionDownload = "download"
)
