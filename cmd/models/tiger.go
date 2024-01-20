package models

/*
Contains the models for the Tiger Table
*/
// Tiger model
type Tiger struct {
	TigerID                string  `json:"tiger_id"`
	Name                   string  `json:"name"`
	DateOfBirth            string  `json:"date_of_birth"`
	LastSeenTimestamp      string  `json:"last_seen_timestamp"`
	LastSeenCoordinatesLat float64 `json:"last_seen_coordinates_lat"`
	LastSeenCoordinatesLon float64 `json:"last_seen_coordinates_lon"`
}

func (tiger *Tiger) TableName() string {
	return "tiger"
}
