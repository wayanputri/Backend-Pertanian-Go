package models

import "time"

type Actions struct {
	ID         int `gorm:"primaryKey"`
	UserID     int
	TargetID   int
	TargetType string
	Action     string
	CreatedAt  time.Time
}
