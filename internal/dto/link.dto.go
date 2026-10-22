package dto

import "time"

type CreateShortlinkRequest struct {
	DestinationURL string `json:"destination_url" binding:"required"`
	Slug           string `json:"slug"`
}

type ShortlinkResponse struct {
	ID             int       `json:"id"`
	Slug           string    `json:"slug"`
	DestinationURL string    `json:"destination_url"`
	ClickCount     int       `json:"click_count"`
	CreatedAt      time.Time `json:"created_at"`
}
