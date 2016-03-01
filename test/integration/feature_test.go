package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/MEDIGO/laika/api"
	"github.com/MEDIGO/laika/store"
	"github.com/MEDIGO/laika/util"
)

type FeatureIntegrationSuite struct {
	FeatureFlagSuite
}

func (s *FeatureIntegrationSuite) TestFeatureCRUD() {
	name := util.Token()
	input := &api.Feature{
		Name: store.String(name),
	}

	created, err := s.client.FeatureCreate(input)
	require.NoError(s.T(), err)
	require.NotEqual(s.T(), 0, created.Id)
	require.Equal(s.T(), store.String(name), created.Name)

	found, err := s.client.FeatureGet(*created.Name)
	require.NoError(s.T(), err)
	require.Equal(s.T(), created.Id, found.Id)
	require.Equal(s.T(), created.Name, input.Name)

	listed, err := s.client.FeatureList()
	require.NoError(s.T(), err)
	require.NotEqual(s.T(), len(listed), 0)
	require.Equal(s.T(), found.Id, listed[len(listed)-1].Id)

	newName := util.Token()
	input = &api.Feature{Name: store.String(newName)}

	updated, err := s.client.FeatureUpdate(*found.Name, input)
	require.NoError(s.T(), err)
	require.Equal(s.T(), found.Id, updated.Id)
	require.Equal(s.T(), store.String(newName), updated.Name)
}

func TestFeatureIntegrationSuite(t *testing.T) {
	suite.Run(t, new(FeatureIntegrationSuite))
}
