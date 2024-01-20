package models

/*
Contains the models for Tiger Sightings data
*/
type TigerSightingData struct {
	SightingID    string  `json:"sighting_id" gorm:"primaryKey"`
	TigerID       string  `json:"tiger_id"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Timestamp     string  `json:"timestamp"`
	SightingImage string  `json:"sighting_image"`
}

func (tigerSightingData *TigerSightingData) TableName() string {
	return "tiger_sighting_data"
}
