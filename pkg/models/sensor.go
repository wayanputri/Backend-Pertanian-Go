package models

type Sensor struct {
	ID                int `gorm:"primaryKey"`
	Type              string
	Location          string
	FarmAreaID        int
	ValueAverageToday float64
	SensorReading     []SensorReading `gorm:"foreignKey:SensorID"`
}
