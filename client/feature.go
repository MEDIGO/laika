package client

import (
	"fmt"

	"github.com/MEDIGO/feature-flag/api"
)

func (c *client) FeatureGet(name string) (*api.Feature, error) {
	out := new(api.Feature)
	err := c.get(fmt.Sprintf("/api/features/%s", name), out)
	return out, err
}

func (c *client) FeatureCreate(in *api.Feature) (*api.Feature, error) {
	out := new(api.Feature)
	err := c.post("/api/features", in, out)
	return out, err
}

func (c *client) FeatureList() ([]*api.Feature, error) {
	out := []*api.Feature{}
	err := c.get(fmt.Sprintf("/api/features"), &out)
	return out, err
}

func (c *client) FeatureUpdate(name string, in *api.Feature) (*api.Feature, error) {
	out := new(api.Feature)
	err := c.patch(fmt.Sprintf("/api/features/%s", name), in, out)
	return out, err
}
