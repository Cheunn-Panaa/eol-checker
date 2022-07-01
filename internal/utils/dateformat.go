package utils

import "time"

//DateFormat returns a date format from unknown type
func DateFormat(date interface{}, format string) string {
	var returnedDate time.Time
	switch v := date.(type) {
	case string:
		returnedDate, _ = time.Parse("2006-01-02", v)
	case time.Time:
		returnedDate = v
	default:
		panic("invalid type: %T")
	}
	return returnedDate.Format(format)
}
