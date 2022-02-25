package api

import (
	"net/http"
	"time"

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
type ProjectCycleList []*ProjectCycle

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
	if p.Support.String != "" {
		supportDate, _ := time.Parse("2006-01-02", p.Support.String)
		return supportDate.Before(alertDate)
	}
	return p.Support.Bool
}
