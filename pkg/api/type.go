package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// StringOrInt is for either string or int type return from the api
type StringOrInt struct {
	String *string
	Int    *int
}

// StringOrBool is for either string or bool type return from the api
type StringOrBool struct {
	String *string
	Bool   *bool
}

// ProjectCycle is the response of endoflife endpoint for specific project
type ProjectCycle struct {
	Cycle          StringOrInt  `json:"cycle"`
	Release        *string      `json:"release"`
	EOL            StringOrBool `json:"eol"`
	Latest         *string      `json:"latest"`
	Link           *string      `json:"link,omitempty"`
	LTS            *bool        `json:"lts,omitempty"`
	Support        StringOrBool `json:"support,omitempty"`
	CycleShortHand StringOrInt  `json:"cycleShortHand"`
	Discontinued   StringOrBool `json:"disconitinued"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Client is api client
type Client struct {
	baseURL    string
	httpClient *http.Client
}

func (p *ProjectCycle) IsSupportedAtDate(alertDate time.Time) bool {
	if p.Support.String != nil {
		supportDate, _ := time.Parse("2006-01-02", *p.Support.String)
		return supportDate.Before(alertDate)
	}
	return *p.Support.Bool
}

// UnmarshalJSON assign json value to appropriate field
func (strint *StringOrInt) UnmarshalJSON(p []byte) error {
	var i interface{}
	if err := json.Unmarshal(p, &i); err != nil {
		return err
	}
	switch val := i.(type) {
	case string:
		strint.String = &val
	case int:
		strint.Int = &val
	case float64:
		var p int = int(val)
		strint.Int = &p
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
		strint.String = &val
	case bool:
		strint.Bool = &val
	default:
		return fmt.Errorf("invalid type: %T", val)
	}
	return nil
}
