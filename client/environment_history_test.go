package client

import (
	"encoding/json"
	"fmt"
	"github.com/MEDIGO/feature-flag/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type EnvironmentHistoryServiceSuite struct {
	ClientSuite
}

func (s *EnvironmentHistoryServiceSuite) TestEnvironmentHistoryGet() {
	s.mux.HandleFunc("/environment_history/1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)
		fmt.Fprint(w, `{"id": 1}`)
	})

	found, err := s.client.EnvironmentHistoryGet(1)
	assert.NoError(s.T(), err)

	expected := &model.EnvironmentHistory{Id: 1}
	assert.Equal(s.T(), expected, found)
}

func (s *EnvironmentHistoryServiceSuite) TestEnvironmentHistoryCreate() {
	input := &model.EnvironmentHistory{Id: 2, Name: model.String("ftest")}

	s.mux.HandleFunc("/environment_history", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "POST", r.Method)

		received := new(model.EnvironmentHistory)
		json.NewDecoder(r.Body).Decode(received)
		assert.Equal(s.T(), input, received)

		fmt.Fprint(w, `{"id": 2, "name": "ftest"}`)
	})

	found, err := s.client.EnvironmentHistoryCreate(input)
	assert.NoError(s.T(), err)

	expected := &model.EnvironmentHistory{Id: 2, Name: model.String("ftest")}
	assert.Equal(s.T(), expected, found)
}

func (s *EnvironmentHistoryServiceSuite) TestListEnvironmentHistory() {
	s.mux.HandleFunc("/environment_history", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)

		fmt.Fprint(w, `[{"id": 1, "name": "f1"}]`)
	})

	found, err := s.client.EnvironmentHistoryList()
	assert.NoError(s.T(), err)
	assert.Len(s.T(), found, 1)

	expected := &model.EnvironmentHistory{Id: 1, Name: model.String("f1")}
	assert.Equal(s.T(), expected, found[0])
}

func (s *EnvironmentHistoryServiceSuite) TestEnvironmentUpdate() {
	input := &model.EnvironmentHistory{Name: model.String("f1")}

	s.mux.HandleFunc("/environment_history/1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "PATCH", r.Method)

		received := new(model.EnvironmentHistory)
		json.NewDecoder(r.Body).Decode(received)
		assert.Equal(s.T(), input, received)

		fmt.Fprint(w, `{"id": 1, "name": "f1.1"}`)
	})

	found, err := s.client.EnvironmentHistoryUpdate(1, input)
	assert.NoError(s.T(), err)

	expected := &model.EnvironmentHistory{Id: 1, Name: model.String("f1.1")}
	assert.Equal(s.T(), expected, found)
}

func TestEnvironmentHistoryServiceSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentHistoryServiceSuite))
}
