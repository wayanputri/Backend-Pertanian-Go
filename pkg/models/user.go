package models

type User struct {
	ID              int `gorm:"primaryKey"`
	Name            string
	Email           string `gorm:"uniqueIndex"`
	Role            string
	Password        string
	Address         string
	LandArea        string
	TypeOfLivestock string

	Cctv       []Cctv       `gorm:"foreignKey:UserID"`
	Action     []Actions    `gorm:"foreignKey:UserID"`
	Livestocks []Livestocks `gorm:"foreignKey:UserID"`
	Crops      []Crops      `gorm:"foreignKey:UserID"`
}
