package models

import "time"

type Status string

const (
	READY      Status = "ready"
	UPLOADED   Status = "uploaded"
	QUEUED     Status = "queued"
	PROCESSING Status = "processing"
	FAILED     Status = "failed"
)

type Video struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	OriginalURL  string    `json:"original_url"`
	Video360URL  string    `json:"video_360_url"`
	Video480URL  string    `json:"video_480_url"`
	Video720URL  string    `json:"video_720_url"`
	Video1080URL string    `json:"video_1080_url"`
	Status       Status    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
