package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cheunn-panaa/eol-checker/configs"
)

// NewHTTPClient generates new http client
func NewHTTPClient(c *configs.Config) *Client {
	return &Client{
		baseURL: c.Default.Url,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

//sendRequest is the basic func for every http request to endoflife.date
func (c *Client) sendRequest(req *http.Request, objStruct interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		var er errorResponse
		if err = json.NewDecoder(res.Body).Decode(&er); err == nil {
			return errors.New(er.Message)
		}
		return fmt.Errorf("error, status code: %d", res.StatusCode)
	}
	if err = json.NewDecoder(res.Body).Decode(&objStruct); err != nil {
		return err
	}
	return nil
}

// GetProjectCycle retrieves specific version of a project
func (c *Client) GetProjectCycle(product *configs.Product) (*ProjectCycle, error) {
	req, err := http.NewRequest("GET", c.generateUrl(product, false), nil)
	if err != nil {
		return nil, err
	}
	res := &ProjectCycle{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetProjectLatestRelease retrieves last version of a project
func (c *Client) GetProjectLatestRelease(product *configs.Product) (*ProjectCycle, error) {
	req, err := http.NewRequest("GET", c.generateUrl(product, true), nil)
	if err != nil {
		return nil, err
	}
	res := ProjectCycleList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return res[0], nil
}

// generateUrl will return specific version endpoint if it's specified
func (c *Client) generateUrl(product *configs.Product, forceNoVersion bool) string {
	url := fmt.Sprintf("%s/api/%s.json",
		c.baseURL, product.Name)
	if product.Version != "" && !forceNoVersion {
		url = fmt.Sprintf("%s/api/%s/%s.json",
			c.baseURL, product.Name, product.Version)

	}
	return url
}
