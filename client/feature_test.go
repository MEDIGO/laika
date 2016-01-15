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

type FeatureServiceSuite struct {
	ClientSuite
}

func (s *FeatureServiceSuite) TestFeatureGet() {
	s.mux.HandleFunc("/features/1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)
		fmt.Fprint(w, `{"id": 1}`)
	})

	found, err := s.client.FeatureGet(1)
	assert.NoError(s.T(), err)

	expected := &model.Feature{Id: 1}
	assert.Equal(s.T(), expected, found)
}

func (s *FeatureServiceSuite) TestFeatureCreate() {
	input := &model.Feature{Id: 2, Name: model.String("ftest")}

	s.mux.HandleFunc("/features", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "POST", r.Method)

		received := new(model.Feature)
		json.NewDecoder(r.Body).Decode(received)
		assert.Equal(s.T(), input, received)

		fmt.Fprint(w, `{"id": 2, "name": "ftest"}`)
	})

	found, err := s.client.FeatureCreate(input)
	assert.NoError(s.T(), err)

	expected := &model.Feature{Id: 2, Name: model.String("ftest")}
	assert.Equal(s.T(), expected, found)
}

func (s *FeatureServiceSuite) TestListFeatures() {
	s.mux.HandleFunc("/features", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "GET", r.Method)

		fmt.Fprint(w, `[{"id": 1, "name": "f1"}]`)
	})

	found, err := s.client.FeatureList()
	assert.NoError(s.T(), err)
	assert.Len(s.T(), found, 1)

	expected := &model.Feature{Id: 1, Name: model.String("f1")}
	assert.Equal(s.T(), expected, found[0])
}

func (s *FeatureServiceSuite) TestFeatureUpdate() {
	input := &model.Feature{Name: model.String("f1")}

	s.mux.HandleFunc("/features/1", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(s.T(), "PATCH", r.Method)

		received := new(model.Feature)
		json.NewDecoder(r.Body).Decode(received)
		assert.Equal(s.T(), input, received)

		fmt.Fprint(w, `{"id": 1, "name": "f1.1"}`)
	})

	found, err := s.client.FeatureUpdate(1, input)
	assert.NoError(s.T(), err)

	expected := &model.Feature{Id: 1, Name: model.String("f1.1")}
	assert.Equal(s.T(), expected, found)
}

func TestFeatureServiceSuite(t *testing.T) {
	suite.Run(t, new(FeatureServiceSuite))
}
