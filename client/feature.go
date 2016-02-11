package client

import (
	"fmt"

	"github.com/MEDIGO/feature-flag/store"
)

func (c *client) FeatureGet(id int64) (*store.Feature, error) {
	out := new(store.Feature)
	err := c.get(fmt.Sprintf("/api/features/%d", id), out)
	return out, err
}

func (c *client) FeatureCreate(in *store.Feature) (*store.Feature, error) {
	out := new(store.Feature)
	err := c.post("/api/features", in, out)
	return out, err
}

func (c *client) FeatureList() ([]*store.Feature, error) {
	out := []*store.Feature{}
	err := c.get(fmt.Sprintf("/api/features"), &out)
	return out, err
}

func (c *client) FeatureUpdate(id int64, in *store.Feature) (*store.Feature, error) {
	out := new(store.Feature)
	err := c.patch(fmt.Sprintf("/api/features/%d", id), in, out)
	return out, err
}
