package models

import "time"

type SensorReading struct {
	ID        int `gorm:"primaryKey"`
	Value     float64
	CreatedAt time.Time
	SensorID  int
}
