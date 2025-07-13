package models

type FarmArea struct {
	ID          int `gorm:"primaryKey"`
	Nama        string
	Type        string
	Wide        string
	Deskription string

	Crop      []Crops      `gorm:"foreignKey:FarmAreaID"`
	Livestock []Livestocks `gorm:"foreignKey:FarmAreaID"`
	Sensor    []Sensor     `gorm:"foreignKey:FarmAreaID"`
}
