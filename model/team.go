package model

import (
	"encoding/json"
	"time"
)

type Team struct {
	ID                 string    `gorm:"primaryKey" json:"id"`
	Name               string    `json:"name"`
	Tag                string    `json:"tag"`
	Bio                string    `json:"bio"`
	Website            string    `json:"website"`
	IconURL            string    `json:"icon_url"`
	BannerURL          string    `gorm:"-" json:"banner_url"`
	Game               string    `json:"game"`
	AverageRank        int       `json:"average_rank"`
	SocialTwitterURL   string    `json:"social_twitter_url"`
	SocialInstagramURL string    `json:"social_instagram_url"`
	SocialTikTokURL    string    `json:"social_tiktok_url"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Team) TableName() string {
	return "team"
}

type TeamUser struct {
	TeamID    string          `gorm:"primaryKey" json:"team_id"`
	UserID    string          `gorm:"primaryKey" json:"user_id"`
	Title     string          `json:"title"`
	Roles     []string        `gorm:"-" json:"roles"`
	User      json.RawMessage `gorm:"-" json:"user"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
}

func (TeamUser) TableName() string {
	return "team_user"
}

type TeamUserRole struct {
	TeamID    string    `gorm:"primaryKey" json:"team_id"`
	UserID    string    `gorm:"primaryKey" json:"user_id"`
	Role      string    `gorm:"primaryKey" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (TeamUserRole) TableName() string {
	return "team_user_role"
}

type TeamOrganization struct {
	TeamID         string    `gorm:"primaryKey" json:"team_id"`
	OrganizationID string    `json:"organization_id"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (TeamOrganization) TableName() string {
	return "team_organization"
}
