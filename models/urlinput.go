package models

type UrlInput struct {
	Url string `json:"url" binding:"required"`
}
