package model

import "time"

type OrganizationUserRole struct {
	OrganizationID string    `json:"organization_id"`
	UserID         string    `json:"user_id"`
	Role          string    `json:"role"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (OrganizationUserRole) TableName() string {
	return "organization_user_role"
}