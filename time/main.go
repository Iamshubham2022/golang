// Current Time: time.Now() gives the current local time.
// Formatting Time: You can format time using the Format method with a reference layout.
// Parsing Time: time.Parse() allows you to parse a string into a Time object.
// UTC Time: UTC() converts time to UTC.
// Adding/Subtracting Durations: Add() and AddDate() are used to add time (e.g., hours, days).
// Time Difference: Sub() calculates the difference between two times as a Duration.
// Unix Timestamp: Unix() converts time into a Unix timestamp (seconds since 1970).
// Timezone Conversion: In() converts time to another timezone.
// Duration Parsing: ParseDuration() parses a duration string (e.g., "2h45m30s").
// Compare Times: After(), Before(), and Equal() methods to compare two Time objects.
// Nanoseconds: UnixNano() returns the time in nanoseconds.
// Custom Formats: You can format the time into a custom layout string.
// Trim Time: You can trim specific parts of time by formatting or using string manipulation.
// Add/Subtract Days: AddDate() allows you to add or subtract years, months, and days.



package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Get Current Time
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	// 2. Format Time to String
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Time:", formattedTime)

	// 3. Parse String to Time
	parsedTime, _ := time.Parse("2006-01-02 15:04:05", "2025-01-18 15:30:00")
	fmt.Println("Parsed Time:", parsedTime)

	// 4. Time in UTC
	utcTime := currentTime.UTC()
	fmt.Println("UTC Time:", utcTime)

	// 5. Add Duration to Time
	oneHourLater := currentTime.Add(1 * time.Hour)
	fmt.Println("One Hour Later:", oneHourLater)

	// 6. Subtract Duration from Time
	oneDayAgo := currentTime.Add(-24 * time.Hour)
	fmt.Println("One Day Ago:", oneDayAgo)

	// 7. Time Difference
	duration := oneHourLater.Sub(currentTime)
	fmt.Println("Duration Between Times:", duration)

	// 8. Convert Time to Unix Timestamp
	unixTimestamp := currentTime.Unix()
	fmt.Println("Unix Timestamp:", unixTimestamp)

	// 9. Convert Unix Timestamp to Time
	timeFromUnix := time.Unix(unixTimestamp, 0)
	fmt.Println("Time from Unix Timestamp:", timeFromUnix)

	// 10. Convert Time to String (ISO 8601 format)
	isoTime := currentTime.Format(time.RFC3339)
	fmt.Println("ISO 8601 Time:", isoTime)

	// 11. Parse String to Time (ISO 8601 format)
	parsedISOTime, _ := time.Parse(time.RFC3339, "2025-01-18T15:04:05Z")
	fmt.Println("Parsed ISO Time:", parsedISOTime)

	// 12. Convert Time to Specific Time Zone
	loc, _ := time.LoadLocation("America/New_York")
	newYorkTime := currentTime.In(loc)
	fmt.Println("Time in New York:", newYorkTime)

	// 13. Time Duration Parsing and Adding
	durationStr := "2h45m30s"
	parsedDuration, _ := time.ParseDuration(durationStr)
	fmt.Println("Parsed Duration:", parsedDuration)
	newTimeAfterDuration := currentTime.Add(parsedDuration)
	fmt.Println("Time After Adding Duration:", newTimeAfterDuration)

	// 14. Duration in Hours, Minutes, and Seconds
	fmt.Println("Duration in Hours:", parsedDuration.Hours())
	fmt.Println("Duration in Minutes:", parsedDuration.Minutes())
	fmt.Println("Duration in Seconds:", parsedDuration.Seconds())

	// 15. Compare Times (After, Before, Equal)
	if oneHourLater.After(currentTime) {
		fmt.Println("One hour later is after the current time")
	}

	if oneDayAgo.Before(currentTime) {
		fmt.Println("One day ago is before the current time")
	}

	if currentTime.Equal(time.Now()) {
		fmt.Println("Times are equal")
	}

	// 16. Trim and Format Time
	trimmedTime := currentTime.Format("2006-01-02")
	fmt.Println("Trimmed Date:", trimmedTime)

	// 17. Convert Between Unix and Nanoseconds
	unixNano := currentTime.UnixNano()
	fmt.Println("Unix Nanoseconds:", unixNano)

	// Convert Nanoseconds to Time
	timeFromNano := time.Unix(0, unixNano)
	fmt.Println("Time from Nanoseconds:", timeFromNano)

	// 18. Formatting Time Using Custom Layout
	customFormat := currentTime.Format("Monday, January 2, 2006, 15:04:05")
	fmt.Println("Custom Format:", customFormat)

	// 19. Time Zone Conversion (Asia/Kolkata)
	kolkataLoc, _ := time.LoadLocation("Asia/Kolkata")
	kolkataTime := currentTime.In(kolkataLoc)
	fmt.Println("Time in Kolkata:", kolkataTime)

	// 20. Time to String (Custom Format)
	customFormattedTime := currentTime.Format("02/01/2006 03:04 PM")
	fmt.Println("Custom Formatted Time:", customFormattedTime)

	// 21. Add Days to Time
	addDays := currentTime.AddDate(0, 0, 5) // Adds 5 days
	fmt.Println("Time After Adding 5 Days:", addDays)

	// 22. Subtract Days from Time
	subtractDays := currentTime.AddDate(0, 0, -3) // Subtracts 3 days
	fmt.Println("Time After Subtracting 3 Days:", subtractDays)
}


