package models

type Livestocks struct {
	ID             int `gorm:"primaryKey"`
	Spesies        string
	Age            string
	HealthStatus   string
	KebutuhanMakan string
	TypeOfFood     string
	Quantity       int
	FarmAreaID     int
	UserID         int
	Reference      string

	Actions []Actions `gorm:"foreignKey:TargetID"`
}
