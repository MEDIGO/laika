package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/MEDIGO/laika/api"
	"github.com/MEDIGO/laika/store"
)

type EnvironmentServiceSuite struct {
	ClientSuite
}

func (s *EnvironmentServiceSuite) TestEnvironmentGet() {
	s.mux.HandleFunc("/api/environments/e1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)
		fmt.Fprint(w, `{"name": "e1"}`)
	})

	found, err := s.client.EnvironmentGet("e1")
	assert.NoError(s.T(), err)

	expected := &api.Environment{Name: store.String("e1")}
	assert.Equal(s.T(), expected, found)
}

func (s *EnvironmentServiceSuite) TestEnvironmentCreate() {
	input := &api.Environment{Id: 2, Name: store.String("etest")}

	s.mux.HandleFunc("/api/environments", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "POST", r.Method)

		received := new(api.Environment)
		json.NewDecoder(r.Body).Decode(received)
		assert.Equal(s.T(), input, received)

		fmt.Fprint(w, `{"id": 2, "name": "etest"}`)
	})

	found, err := s.client.EnvironmentCreate(input)
	assert.NoError(s.T(), err)

	expected := &api.Environment{Id: 2, Name: store.String("etest")}
	assert.Equal(s.T(), expected, found)
}

func (s *EnvironmentServiceSuite) TestListEnvironments() {
	s.mux.HandleFunc("/api/environments", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)

		fmt.Fprint(w, `[{"id": 1, "name": "e1"}]`)
	})

	found, err := s.client.EnvironmentList()
	assert.NoError(s.T(), err)
	assert.Len(s.T(), found, 1)

	expected := &api.Environment{Id: 1, Name: store.String("e1")}
	assert.Equal(s.T(), expected, found[0])
}

func (s *EnvironmentServiceSuite) TestEnvironmentUpdate() {
	input := &api.Environment{Name: store.String("e1")}

	s.mux.HandleFunc("/api/environments/e1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "PATCH", r.Method)

		received := new(api.Environment)
		json.NewDecoder(r.Body).Decode(received)
		assert.Equal(s.T(), input, received)

		fmt.Fprint(w, `{"id": 1, "name": "e1.1"}`)
	})

	found, err := s.client.EnvironmentUpdate("e1", input)
	assert.NoError(s.T(), err)

	expected := &api.Environment{Id: 1, Name: store.String("e1.1")}
	assert.Equal(s.T(), expected, found)
}

func TestEnvironmentServiceSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentServiceSuite))
}
