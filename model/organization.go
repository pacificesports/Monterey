package model

import "time"

type Organization struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Tag                string    `json:"tag"`
	Bio                string    `json:"bio"`
	Website            string    `json:"website"`
	IconURL            string    `json:"icon_url"`
	BannerURL          string    `json:"banner_url"`
	SocialTwitterURL   string    `json:"social_twitter_url"`
	SocialInstagramURL string    `json:"social_instagram_url"`
	SocialTikTokURL    string    `json:"social_tiktok_url"`
	Verified           bool      `json: "verified"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
}
