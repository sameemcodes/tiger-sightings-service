package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Todo : Generate a unique ID for the tiger sighting
// GenerateTigerSightingID generates a TigerSighting ID with the format "TIG_SIGHT_XXX".
func GenerateTigerSightingID() string {
	// Use the current timestamp to ensure uniqueness
	timestamp := time.Now().UnixNano()

	randomNumber := rand.Intn(10000000000)

	tigerSightingID := fmt.Sprintf("TIG_SIGHT_%d%03d", timestamp, randomNumber)

	return tigerSightingID
}
