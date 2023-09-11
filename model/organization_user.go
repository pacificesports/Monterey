package model

import "time"

type OrganizationUser struct {
	OrganizationID string    `json:"organization_id"`
	UserID         string    `json:"user_id"`
	Title          string    `json:"title"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (OrganizationUser) TableName() string {
	return "organization_user"
}
