package api

import (
	"net/http"

	"github.com/cheunn-panaa/eol-checker/pkg/utils"
)

// ProjectCycle is the response of endoflife endpoint for specific project
type ProjectCycle struct {
	Cycle          utils.StringOrInt  `json:"cycle"`
	Release        string             `json:"release"`
	EOL            utils.StringOrBool `json:"eol"`
	Latest         string             `json:"latest"`
	Link           string             `json:"link,omitempty"`
	LTS            bool               `json:"lts,omitempty"`
	Support        utils.StringOrBool `json:"support,omitempty"`
	CycleShortHand utils.StringOrInt  `json:"cycleShortHand"`
	Discontinued   utils.StringOrBool `json:"disconitinued"`
	LatestCycle    utils.StringOrInt
}

//ProjectCycleList list of ProjectCycle
type ProjectCycleList []*ProjectCycle

//errorResponse basic http error type
type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Client is api client
type Client struct {
	baseURL    string
	httpClient *http.Client
}
