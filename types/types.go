package types

import "github.com/aws/aws-sdk-go-v2/service/s3"

type R2Service struct {
	R2Client  *s3.Client
	Bucket    string
	AccountID string
}

type VideoJob struct {
	VideoID string `json:"video_id"`
}

type CreateVideoRequest struct {
	VideoID     string `json:"video_id"`
	Title       string `json:"title"`
	OriginalUrl string `json:"original_url"`
}

type CreateVideoResponse struct {
	VideoID      string `json:"video_id"`
	Video360URL  string `json:"video_360_url"`
	Video480URL  string `json:"video_480_url"`
	Video720URL  string `json:"video_720_url"`
	Video1080URL string `json:"video_1080_url"`
	Status       string `json:"status"`
}
