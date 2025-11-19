package models

import (
	"time"

	"gorm.io/gorm"
)

type Workspace struct {
	gorm.Model
	Name        string            `gorm:"not null" json:"name"`
	Description string            `json:"description"`
	OwnerID     uint              `gorm:"not null;index" json:"owner_id"`
	Owner       User              `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
	Members     []User            `gorm:"many2many:workspace_members;" json:"members,omitempty"`
	Files       []File            `gorm:"foreignKey:WorkspaceID" json:"files,omitempty"`
	Memberships []WorkspaceMember `gorm:"foreignKey:WorkspaceID" json:"memberships,omitempty"`
}

type WorkspaceMember struct {
	gorm.Model
	WorkspaceID uint      `gorm:"not null;uniqueIndex:idx_workspace_user" json:"workspace_id"`
	Workspace   Workspace `gorm:"foreignKey:WorkspaceID" json:"workspace,omitempty"`
	UserID      uint      `gorm:"not null;uniqueIndex:idx_workspace_user" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user,omitempty"`

	Role     string    `gorm:"not null;default:'viewer'" json:"role"`
	JoinedAt time.Time `gorm:"autoCreateTime" json:"joined_at"`
}

const (
	RoleOwner  = "owner"
	RoleAdmin  = "admin"
	RoleEditor = "editor"
	RoleViewer = "viewer"
)
