package client

import (
	"fmt"

	"github.com/MEDIGO/feature-flag/store"
)

func (c *client) EnvironmentGet(id int64) (*store.Environment, error) {
	out := new(store.Environment)
	err := c.get(fmt.Sprintf("/api/environments/%d", id), out)
	return out, err
}

func (c *client) EnvironmentCreate(in *store.Environment) (*store.Environment, error) {
	out := new(store.Environment)
	err := c.post("/api/environments", in, out)
	return out, err
}

func (c *client) EnvironmentList() ([]*store.Environment, error) {
	out := []*store.Environment{}
	err := c.get(fmt.Sprintf("/api/environments"), &out)
	return out, err
}

func (c *client) EnvironmentUpdate(id int64, in *store.Environment) (*store.Environment, error) {
	out := new(store.Environment)
	err := c.patch(fmt.Sprintf("/api/environments/%d", id), in, out)
	return out, err
}
