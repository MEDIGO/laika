package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/MEDIGO/feature-flag/model"
	"github.com/MEDIGO/feature-flag/util"
)

type FeatureIntegrationSuite struct {
	FeatureFlagSuite
}

func (s *FeatureIntegrationSuite) TestFeatureCRU() {
	name := util.Token()
	input := &model.Feature{
		Name: model.String(name),
	}

	created, err := s.client.FeatureCreate(input)
	require.NoError(s.T(), err)
	require.NotEqual(s.T(), 0, created.Id)
	require.Equal(s.T(), model.String(name), created.Name)

	found, err := s.client.FeatureGet(created.Id)
	require.NoError(s.T(), err)
	require.Equal(s.T(), created.Id, found.Id)
	require.Equal(s.T(), created.Name, input.Name)

	listed, err := s.client.FeatureList()
	require.NoError(s.T(), err)
	require.NotEqual(s.T(), len(listed), 0)
	require.Equal(s.T(), found.Id, listed[len(listed)-1].Id)

	newname := util.Token()
	input = &model.Feature{Name: model.String(newname)}

	updated, err := s.client.FeatureUpdate(found.Id, input)
	require.NoError(s.T(), err)
	require.Equal(s.T(), found.Id, updated.Id)
	require.Equal(s.T(), model.String(newname), updated.Name)
}

func TestFeatureIntegrationSuite(t *testing.T) {
	suite.Run(t, new(FeatureIntegrationSuite))
}
