package client

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

// Config is used to parameterize a client
type Config struct {
	Addr            string
	Username        string
	Password        string
	Environment     string
	PollingInterval time.Duration
}

// Client is a threadsafe client for the Laika API.
type Client interface {
	IsEnabled(string, bool) bool
}

type client struct {
	features *FeatureCache

	url             *url.URL
	username        string
	password        string
	environment     string
	pollingInterval time.Duration
}

// NewClient creates a new Client.
func NewClient(conf Config) (Client, error) {
	if conf.Addr == "" {
		return nil, errors.New("missing address")
	}

	if conf.Username == "" {
		return nil, errors.New("missing username")
	}

	if conf.Password == "" {
		return nil, errors.New("missing password")
	}

	if conf.Environment == "" {
		return nil, errors.New("missing environment")
	}

	if conf.PollingInterval == 0 {
		conf.PollingInterval = 10 * time.Second
	}

	addr, err := url.Parse(conf.Addr)
	if err != nil {
		return nil, err
	}

	endpoint, err := url.Parse("/api/features")
	if err != nil {
		return nil, err
	}

	cl := &client{
		features:        NewFeatureCache(),
		url:             addr.ResolveReference(endpoint),
		username:        conf.Username,
		password:        conf.Password,
		environment:     conf.Environment,
		pollingInterval: conf.PollingInterval,
	}

	if err := cl.poll(); err != nil {
		return nil, err
	}

	go func(cl *client) {
		for {
			cl.poll()
			time.Sleep(cl.pollingInterval)
		}
	}(cl)

	return cl, nil
}

// Poll polls the Laika API for the latest Feature statuses, storing the results
// in the internal cache.
func (c *client) poll() error {
	req, err := http.NewRequest("GET", c.url.String(), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		e := new(Error)
		if err := json.NewDecoder(res.Body).Decode(e); err != nil {
			return err
		}

		return errors.New(e.Message)
	}

	features := []*Feature{}
	if err := json.NewDecoder(res.Body).Decode(&features); err != nil {
		return err
	}

	c.features.AddAll(features)

	return nil
}

// IsEnabled returns the status of the Feature. If the Feature is unknown,
// it will return the default value provided.
func (c *client) IsEnabled(name string, defval bool) bool {
	feature := c.features.Get(name)
	if feature == nil {
		return defval
	}

	status, ok := feature.Status[c.environment]
	if !ok {
		return defval
	}

	return status
}
