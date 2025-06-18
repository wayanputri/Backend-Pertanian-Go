package models

type Cctv struct {
	ID           int `gorm:"primaryKey"`
	Location     string
	RecordingUrl string
	UserID       int
}
