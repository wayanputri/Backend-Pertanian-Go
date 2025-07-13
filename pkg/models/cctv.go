package models

import "time"

type Cctv struct {
	ID           int `gorm:"primaryKey"`
	Location     string
	RecordingUrl string
	UserID       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
