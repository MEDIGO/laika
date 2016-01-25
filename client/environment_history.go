package client

import (
	"fmt"
	"github.com/MEDIGO/feature-flag/model"
)

func (c *client) EnvironmentHistoryGet(id int64) (*model.EnvironmentHistory, error) {
	out := new(model.EnvironmentHistory)
	err := c.get(fmt.Sprintf("/environment_history/%d", id), out)
	return out, err
}

func (c *client) EnvironmentHistoryCreate(in *model.EnvironmentHistory) (*model.EnvironmentHistory, error) {
	out := new(model.EnvironmentHistory)
	err := c.post("/environment_history", in, out)
	return out, err
}

func (c *client) EnvironmentHistoryList() ([]*model.EnvironmentHistory, error) {
	out := []*model.EnvironmentHistory{}
	err := c.get(fmt.Sprintf("/environment_history"), &out)
	return out, err
}

func (c *client) EnvironmentHistoryUpdate(id int64, in *model.EnvironmentHistory) (*model.EnvironmentHistory, error) {
	out := new(model.EnvironmentHistory)
	err := c.patch(fmt.Sprintf("/environment_history/%d", id), in, out)
	return out, err
}
