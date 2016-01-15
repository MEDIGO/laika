package client

import (
	"fmt"

	"github.com/MEDIGO/feature-flag/model"
)

func (c *client) EnvironmentGet(featureName string, environmentName string) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.get(fmt.Sprintf("/features/%s/environments/%s", featureName, environmentName), out)
	return out, err
}

func (c *client) EnvironmentCreate(in *model.Environment) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.post("/environments", in, out)
	return out, err
}

func (c *client) EnvironmentList() ([]*model.Environment, error) {
	out := []*model.Environment{}
	err := c.get(fmt.Sprintf("/environments"), &out)
	return out, err
}

func (c *client) EnvironmentUpdate(id int64, in *model.Environment) (*model.Environment, error) {
	out := new(model.Environment)
	err := c.patch(fmt.Sprintf("/environments/%d", id), in, out)
	return out, err
}
