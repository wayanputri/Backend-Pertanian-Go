package models

type Crops struct {
	ID           int `gorm:"primaryKey"`
	Spesies      string
	GrowStage    string
	HealthStatus string

	FarmAreaID int
	UserID     int
	Actions    []Actions `gorm:"foreignKey:TargetID"`
}
