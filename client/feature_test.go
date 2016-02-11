package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/MEDIGO/feature-flag/api"
	"github.com/MEDIGO/feature-flag/store"
)

type FeatureServiceSuite struct {
	ClientSuite
}

func (s *FeatureServiceSuite) TestFeatureGet() {
	s.mux.HandleFunc("/api/features/f1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)
		fmt.Fprint(w, `{"name": "f1"}`)
	})

	found, err := s.client.FeatureGet("f1")
	assert.NoError(s.T(), err)

	expected := &api.Feature{Name: store.String("f1")}
	assert.Equal(s.T(), expected, found)
}

func (s *FeatureServiceSuite) TestFeatureCreate() {
	input := &api.Feature{Id: 2, Name: store.String("ftest")}

	s.mux.HandleFunc("/api/features", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "POST", r.Method)

		received := new(api.Feature)
		json.NewDecoder(r.Body).Decode(received)
		assert.Equal(s.T(), input, received)

		fmt.Fprint(w, `{"id": 2, "name": "ftest"}`)
	})

	found, err := s.client.FeatureCreate(input)
	assert.NoError(s.T(), err)

	expected := &api.Feature{Id: 2, Name: store.String("ftest")}
	assert.Equal(s.T(), expected, found)
}

func (s *FeatureServiceSuite) TestListFeatures() {
	s.mux.HandleFunc("/api/features", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)

		fmt.Fprint(w, `[{"id": 1, "name": "f1"}]`)
	})

	found, err := s.client.FeatureList()
	assert.NoError(s.T(), err)
	assert.Len(s.T(), found, 1)

	expected := &api.Feature{Id: 1, Name: store.String("f1")}
	assert.Equal(s.T(), expected, found[0])
}

func (s *FeatureServiceSuite) TestFeatureUpdate() {
	input := &api.Feature{Name: store.String("f1")}

	s.mux.HandleFunc("/api/features/f1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "PATCH", r.Method)

		received := new(api.Feature)
		json.NewDecoder(r.Body).Decode(received)
		assert.Equal(s.T(), input, received)

		fmt.Fprint(w, `{"id": 1, "name": "f1.1"}`)
	})

	found, err := s.client.FeatureUpdate("f1", input)
	assert.NoError(s.T(), err)

	expected := &api.Feature{Id: 1, Name: store.String("f1.1")}
	assert.Equal(s.T(), expected, found)
}

func TestFeatureServiceSuite(t *testing.T) {
	suite.Run(t, new(FeatureServiceSuite))
}
