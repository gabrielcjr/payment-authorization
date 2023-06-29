package dateTime

import (
	"database/sql"
	"time"

	"github.com/renatospaka/payment-transaction/utils"
)

// Convert date/time values to string applying the ISO format
// similar to time.RFC3339
func FormatDateToNull(date time.Time) string {
	if !date.IsZero() {
		return date.Format(utils.ISO_FORMAT)
	} 
	return utils.NULL_DATE_STR
}

// When database/sql (postgres) returns a date/time value, it may be null
// So it is necessary to verify when the value is valid (time is a date/time)
// or invalid (time is zero date/tim) before converting to the proper format
func FormatNullDate(nullDate sql.NullTime) time.Time {
	if nullDate.Valid {
		return nullDate.Time
	} 
	return utils.NULL_DATE
}
