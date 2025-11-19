package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Filename    string      `gorm:"not null" json:"filename"`
	Filepath    string      `gorm:"not null" json:"filepath"`
	Mimetype    string      `gorm:"not null" json:"mimetype"`
	Size        int64       `gorm:"not null" json:"size"`
	UploadedAt  time.Time   `gorm:"autoCreateTime" json:"uploaded_at"`
	UserID      uint        `gorm:"not null" json:"user_id"`
	User        User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	WorkspaceID *uint       `gorm:"index" json:"workspace_id,omitempty"` // null = personal file
	Workspace   *Workspace  `gorm:"foreignKey:WorkspaceID" json:"workspace,omitempty"`
	Categories  []Category  `gorm:"many2many:file_categories;" json:"categories,omitempty"`
	Shares      []FileShare `gorm:"foreignKey:FileID" json:"shares,omitempty"`
}
