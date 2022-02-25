package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

// StringOrInt is for either string or int type return from the api
type StringOrInt struct {
	String string
	Int    int
}

// StringOrBool is for either string or bool type return from the api
type StringOrBool struct {
	String string
	Bool   bool
}

// UnmarshalJSON assign json value to appropriate field
func (strint *StringOrInt) UnmarshalJSON(p []byte) error {
	var i interface{}
	if err := json.Unmarshal(p, &i); err != nil {
		return err
	}
	switch val := i.(type) {
	case string:
		strint.String = val
	case int:
		strint.Int = val
	case float64:
		var p int = int(val)
		strint.Int = p
	default:
		return fmt.Errorf("invalid type: %T", val)
	}
	return nil
}

// UnmarshalJSON assign json value to appropriate field
func (strint *StringOrBool) UnmarshalJSON(p []byte) error {
	var i interface{}
	if err := json.Unmarshal(p, &i); err != nil {
		return err
	}
	switch val := i.(type) {
	case string:
		strint.String = val
	case bool:
		strint.Bool = val
	default:
		return fmt.Errorf("invalid type: %T", val)
	}
	return nil
}

func (strint *StringOrInt) ToString() string {
	if strint.String == "" {
		return fmt.Sprintf("%d", strint.Int)
	}
	return strint.String
}
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
