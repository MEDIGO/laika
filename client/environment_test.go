package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/MEDIGO/feature-flag/model"
)

type EnvironmentServiceSuite struct {
	ClientSuite
}

func (s *EnvironmentServiceSuite) TestEnvironmentGet() {
	s.mux.HandleFunc("/features/featureName/environments/environmentName", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)
		fmt.Fprint(w, `{"id": 1, "feature_name": "featureName", "environment_name": "environmentName"}`)
	})

	found, err := s.client.EnvironmentGet("featureName", "environmentName")
	assert.NoError(s.T(), err)

	expected := &model.Environment{Id: 1}
	assert.Equal(s.T(), expected, found)
}

func (s *EnvironmentServiceSuite) TestEnvironmentCreate() {
	input := &model.Environment{Id: 2, Name: model.String("ftest")}

	s.mux.HandleFunc("/environments", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "POST", r.Method)

		received := new(model.Environment)
		json.NewDecoder(r.Body).Decode(received)
		assert.Equal(s.T(), input, received)

		fmt.Fprint(w, `{"id": 2, "name": "ftest"}`)
	})

	found, err := s.client.EnvironmentCreate(input)
	assert.NoError(s.T(), err)

	expected := &model.Environment{Id: 2, Name: model.String("ftest")}
	assert.Equal(s.T(), expected, found)
}

func (s *EnvironmentServiceSuite) TestListEnvironments() {
	s.mux.HandleFunc("/environments", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)

		fmt.Fprint(w, `[{"id": 1, "name": "f1"}]`)
	})

	found, err := s.client.EnvironmentList()
	assert.NoError(s.T(), err)
	assert.Len(s.T(), found, 1)

	expected := &model.Environment{Id: 1, Name: model.String("f1")}
	assert.Equal(s.T(), expected, found[0])
}

func (s *EnvironmentServiceSuite) TestEnvironmentUpdate() {
	input := &model.Environment{Name: model.String("f1")}

	s.mux.HandleFunc("/environments/1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "PATCH", r.Method)

		received := new(model.Environment)
		json.NewDecoder(r.Body).Decode(received)
		assert.Equal(s.T(), input, received)

		fmt.Fprint(w, `{"id": 1, "name": "f1.1"}`)
	})

	found, err := s.client.EnvironmentUpdate(1, input)
	assert.NoError(s.T(), err)

	expected := &model.Environment{Id: 1, Name: model.String("f1.1")}
	assert.Equal(s.T(), expected, found)
}

func TestEnvironmentServiceSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentServiceSuite))
}
