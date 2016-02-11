package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/MEDIGO/feature-flag/api"
)

type Client interface {
	FeatureCreate(*api.Feature) (*api.Feature, error)
	FeatureGet(name string) (*api.Feature, error)
	FeatureList() ([]*api.Feature, error)
	FeatureUpdate(name string, in *api.Feature) (*api.Feature, error)

	EnvironmentCreate(*api.Environment) (*api.Environment, error)
	EnvironmentGet(name string) (*api.Environment, error)
	EnvironmentList() ([]*api.Environment, error)
	EnvironmentUpdate(name string, in *api.Environment) (*api.Environment, error)
}

type client struct {
	baseURL   *url.URL
	userAgent string
}

func NewClient(u string) (Client, error) {
	baseURL, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	return &client{baseURL: baseURL, userAgent: "feature-flag-client"}, nil
}

func (c *client) do(method, endpoint string, in interface{}, out interface{}) error {
	rel, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	url := c.baseURL.ResolveReference(rel)
	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/json")

	if in != nil {
		payload, err := json.Marshal(in)
		if err != nil {
			return err
		}
		buf := bytes.NewBuffer(payload)
		req.Body = ioutil.NopCloser(buf)

		req.ContentLength = int64(len(payload))
		req.Header.Set("Content-Length", strconv.Itoa(len(payload)))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if code := res.StatusCode; 200 <= code && code <= 299 {
		if out != nil {
			return json.NewDecoder(res.Body).Decode(out)
		} else {
			return nil
		}
	}

	e := new(api.APIError)
	json.NewDecoder(res.Body).Decode(e)

	return e
}

func (c *client) get(endpoint string, out interface{}) error {
	return c.do("GET", endpoint, nil, out)
}

func (c *client) post(endpoint string, in, out interface{}) error {
	return c.do("POST", endpoint, in, out)
}

func (c *client) patch(endpoint string, in, out interface{}) error {
	return c.do("PATCH", endpoint, in, out)
}
