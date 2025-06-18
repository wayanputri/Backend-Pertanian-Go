package models

type Livestocks struct {
	ID           int `gorm:"primaryKey"`
	Spesies      string
	Age          string
	HealthStatus string
	Quantity     int
	FarmAreaID   int
	UserID       int

	Actions []Actions `gorm:"foreignKey:TargetID"`
}
