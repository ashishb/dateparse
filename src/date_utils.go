package dateparse

import (
	"strings"
	"time"
)

// Remember the last successful format and reuse it to same time
var lastSuccessfulFormat = time.RFC822Z

// GetParsedDate can handle date coming from email header as well as
// Calendar's date and time.
// Different email clients seems to be using different date formats.
// TODO: store the last use date pattern and try it first the next time.
func GetParsedDate(date string) (*time.Time, error) {
	// The expected time format is "Fri, 20 Mar 2020 17:05:25 +0000" but some email clients
	// add extraneous information at the end, for example,
	// "Fri, 20 Mar 2020 17:05:25 +0000 (UTC)"
	// Try again by removing that
	index := strings.Index(date, "(")
	if index > 0 {
		date = strings.TrimSpace(date[:index])
	}
	// Occasionally, dates have UT instead of UTC :/
	if strings.HasSuffix(date, "UT") {
		// Change to UT to UTC
		date += "C"
	}

	//Another common case is when an extraneous GMT with timezone is present
	// Mon, 25 Nov 2019 18:18:47 +0100 GMT
	if strings.HasSuffix(date, "GMT") &&
		strings.Contains(date, "+") {
		date = strings.TrimSuffix(date, "GMT")
	}
	date = strings.TrimSpace(date)

	// Remove commas
	date = strings.ReplaceAll(date, ", ", " ")
	date = strings.ReplaceAll(date, ",", " ")

	formats := []string{
		lastSuccessfulFormat,
		"02 Jan 06 15:04 MST", // time.RFC822
		"2 Jan 06 15:04:05 -0700",
		"2 Jan 06 15:04:05 MST",
		"2 Jan 2006 15:04:05 -0700",
		"2 Jan 2006 15:04:05 MST",
		"02 Jan 06 15:04:05 -0700",
		"02 Jan 06 15:04:05 MST",
		"02 Jan 2006 15:04:05 -0700",
		"02 Jan 2006 15:04:05 MST",
		"02 Jan 06 15:04 -0700", // time.RFC822Z,
		// time.RFC1123 with comma removed
		"Mon 02 Jan 2006 15:04:05 MST",
		// time.RFC1123Z with comma removed
		"Mon 02 Jan 2006 15:04:05 -0700",
		"Mon 2 Jan 2006 15:04:05 -0700",
		"Mon 2 Jan 2006 15:04:05 MST",
		// time.RFC3339 with comma removed
		"2006-01-02T15:04:05Z07:00",
		"Jan _2 2006",
		"January _2 2006",
		"Monday, 02-Jan-06 15:04:05 MST", //	time.RFC850
		// time.RFC850 with comma removed
		"Monday 02-Jan-06 15:04:05 MST",
		"Monday 02-Jan-2006 15:04:05 MST",
	}
	var tmpDate time.Time
	var err error
	for _, format := range formats {
		tmpDate, err = time.Parse(format, date)
		if err == nil {
			// For faster parsing in the future
			lastSuccessfulFormat = format
			return &tmpDate, nil
		}
	}
	return nil, err
}
