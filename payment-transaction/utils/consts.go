package utils

import "time"

const (
	ISO_FORMAT string = "2006-01-02T15:04:05:07Z000" // similar to time.RFC3339
)

var (
	NULL_DATE     time.Time = time.Time{}
	NULL_DATE_NIL string    = time.Time{}.String()
	NULL_DATE_STR string    = "null"
)
