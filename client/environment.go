package client

import (
	"fmt"

	"github.com/MEDIGO/laika/api"
)

func (c *client) EnvironmentGet(name string) (*api.Environment, error) {
	out := new(api.Environment)
	err := c.get(fmt.Sprintf("/api/environments/%s", name), out)
	return out, err
}

func (c *client) EnvironmentCreate(in *api.Environment) (*api.Environment, error) {
	out := new(api.Environment)
	err := c.post("/api/environments", in, out)
	return out, err
}

func (c *client) EnvironmentList() ([]*api.Environment, error) {
	out := []*api.Environment{}
	err := c.get(fmt.Sprintf("/api/environments"), &out)
	return out, err
}

func (c *client) EnvironmentUpdate(name string, in *api.Environment) (*api.Environment, error) {
	out := new(api.Environment)
	err := c.patch(fmt.Sprintf("/api/environments/%s", name), in, out)
	return out, err
}
