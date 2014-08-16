package gographite

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

// Data as returned by graphite
type GraphiteResult struct {
	Target     string          `json:"target"`
	Datapoints [][]interface{} `json:"datapoints"`
}

// A datapoint, as we'd like to see it
type Datapoint struct {
	X int     `json:"x"`
	Y float64 `json:"y"`
}

// A proper result
type Result struct {
	Target     string      `json:"target"`
	Datapoints []Datapoint `json:"datapoints"`
}

// A graphite client
type Client struct {
	baseURL *url.URL
}

// Helper to make creating a new graphite client a little easier
func NewClient(baseURL string) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{BaseURL: u}, nil
}

// Generate a proper request URL based on the inputs
func (c *Client) RequestURL(targets []string, from string) string {
	u := c.BaseURL
	u.Path = "/render"

	q := url.Values{}
	q.Set("format", "json")
	q.Set("from", from)
	for _, t := range targets {
		q.Add("target", t)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

// Hit the graphite render endpoint with the given settings and return a slice of Results
func (c *Client) Render(targets []string, from string) ([]Result, error) {
	res, err := http.Get(c.RequestURL(targets, from))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("non-200 status code returned")
	}

	if res.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("non-json response returned")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var graphiteResults []GraphiteResult
	err = json.Unmarshal(body, &graphiteResults)
	if err != nil {
		return nil, err
	}

	results := make([]Result, 0, len(graphiteResults))
	for _, gr := range graphiteResults {
		result := Result{Target: gr.Target}

		for _, gdp := range gr.Datapoints {
			dp := Datapoint{}
			dp.X = int(gdp[1].(float64))

			// if we have a Float64, grab that value
			// otherwise, we have null, so just move on and use the default of 0
			val := reflect.ValueOf(gdp[0])
			switch val.Kind() {
			case reflect.Float64:
				dp.Y = gdp[0].(float64)
			}
			result.Datapoints = append(result.Datapoints, dp)
		}

		results = append(results, result)
	}
	return results, nil
}
