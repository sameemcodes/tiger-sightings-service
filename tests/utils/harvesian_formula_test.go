package tests

import (
	"fmt"
	"testing"
	"tigerhall-kittens/cmd/utils"
)

func TestHaversine(t *testing.T) {
	lat1, lon1 := 34.0522, -118.2437 // Los Angeles
	lat2, lon2 := 37.7749, -122.4194 // San Francisco
	expectedDistance := 2
	distance := utils.Haversine(lat1, lon1, lat2, lon2)
	fmt.Println("Distance between Los Angeles and San Francisco:", distance, "km")
	// Check if the calculated distance is close to the expected distance
	if distance < 5 {
		t.Errorf("Expected %v but got %v", expectedDistance, distance)
	}
}

func TestHaversine2(t *testing.T) {
	// Example usage
	lat1, lon1 := 43.7128, -74.0060
	lat2, lon2 := 40.7413, -74.0060

	expectedDistance := 4 // Expected distance in kilometers

	distance := utils.Haversine(lat1, lon1, lat2, lon2)
	fmt.Println("Distance between Los Angeles and San Francisco:", distance, "km")

	// Check if the calculated distance is close to the expected distance
	if distance < 5 {
		t.Errorf("Expected %v but got %v", expectedDistance, distance)
	}
}

func TestHaversine3(t *testing.T) {
	// Example usage
	lat1, lon1 := 40.7128, -74.0060 // Los Angeles
	lat2, lon2 := 40.7060, -74.0110 // San Francisco

	expectedDistance := 4 // Expected distance in kilometers

	distance := utils.Haversine(lat1, lon1, lat2, lon2)
	fmt.Println("Distance between Los Angeles and San Francisco:", distance, "km")
	// Check if the calculated distance is close to the expected distance
	if distance > 5 {
		t.Errorf("Expected %v but got %v", expectedDistance, distance)
	}
}
