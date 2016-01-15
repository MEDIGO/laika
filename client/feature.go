package client

import (
	"fmt"
	"github.com/MEDIGO/feature-flag/model"
)

func (c *client) FeatureGet(id int64) (*model.Feature, error) {
	out := new(model.Feature)
	err := c.get(fmt.Sprintf("/features/%d", id), out)
	return out, err
}

func (c *client) FeatureCreate(in *model.Feature) (*model.Feature, error) {
	out := new(model.Feature)
	err := c.post("/features", in, out)
	return out, err
}

func (c *client) FeatureList() ([]*model.Feature, error) {
	out := []*model.Feature{}
	err := c.get(fmt.Sprintf("/features"), &out)
	return out, err
}

func (c *client) FeatureUpdate(id int64, in *model.Feature) (*model.Feature, error) {
	out := new(model.Feature)
	err := c.patch(fmt.Sprintf("/features/%d", id), in, out)
	return out, err
}
