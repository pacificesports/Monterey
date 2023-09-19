package model

import (
	"encoding/json"
	"time"
)

type Organization struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	Tag                string    `json:"tag"`
	Bio                string    `json:"bio"`
	Website            string    `json:"website"`
	IconURL            string    `json:"icon_url"`
	BannerURL          string    `json:"banner_url"`
	SocialTwitterURL   string    `json:"social_twitter_url"`
	SocialInstagramURL string    `json:"social_instagram_url"`
	SocialTikTokURL    string    `json:"social_tiktok_url"`
	Verified           bool      `json:"verified"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Organization) TableName() string {
	return "organization"
}

type OrganizationUser struct {
	OrganizationID string          `json:"organization_id"`
	UserID         string          `json:"user_id"`
	Title          string          `json:"title"`
	Roles          []string        `gorm:"-" json:"roles"`
	User           json.RawMessage `gorm:"-" json:"user"`
	CreatedAt      time.Time       `gorm:"autoCreateTime" json:"created_at"`
}

func (OrganizationUser) TableName() string {
	return "organization_user"
}

type OrganizationUserRole struct {
	OrganizationID string    `json:"organization_id"`
	UserID         string    `json:"user_id"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (OrganizationUserRole) TableName() string {
	return "organization_user_role"
}
