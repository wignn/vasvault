package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string      `gorm:"not null" json:"username"`
	Email       string      `gorm:"uniqueIndex;not null" json:"email"`
	Password    string      `gorm:"not null" json:"-"`
	Files       []File      `gorm:"foreignKey:UserID" json:"files,omitempty"`
	SharedFiles []FileShare `gorm:"foreignKey:SharedWithUserID" json:"shared_files,omitempty"`
	Workspaces  []Workspace `gorm:"many2many:workspace_members;" json:"workspaces,omitempty"`
}
