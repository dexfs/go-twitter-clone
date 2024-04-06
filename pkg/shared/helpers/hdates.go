package helpers

import "time"

func IsToday(createdAt time.Time) bool {
	// Get the current date
	today := time.Now().UTC().Truncate(24 * time.Hour)

	// Truncate the createdAt time to ignore hours, minutes, and seconds
	createdAtDate := createdAt.UTC().Truncate(24 * time.Hour)

	// Compare the createdAt date with today's date
	return createdAtDate.Equal(today)
}
