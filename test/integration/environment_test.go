package integration

import (
	"github.com/MEDIGO/feature-flag/model"
	"github.com/MEDIGO/feature-flag/util"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type EnvironmentIntegrationSuite struct {
	FeatureFlagSuite
}

func (s *EnvironmentIntegrationSuite) TestEnvironmentCRU() {

	name := util.Token()
	input := &model.Environment{
		CreatedAt: model.Time(time.Now()),
		Enabled:   model.Bool(true),
		FeatureId: model.Int(1),
		Name:      model.String(name),
	}

	created, err := s.client.EnvironmentCreate(input)
	require.NoError(s.T(), err)
	require.NotEqual(s.T(), 0, created.Id)
	require.Equal(s.T(), model.Bool(true), created.Enabled)
	require.Equal(s.T(), model.Int(1), created.FeatureId)
	require.Equal(s.T(), model.String(name), created.Name)

	found, err := s.client.EnvironmentGet(created.Id)
	require.NoError(s.T(), err)
	require.Equal(s.T(), created.Id, found.Id)
	require.Equal(s.T(), created.Enabled, input.Enabled)
	require.Equal(s.T(), created.FeatureId, input.FeatureId)
	require.Equal(s.T(), created.Name, input.Name)

	listed, err := s.client.EnvironmentList()
	require.NoError(s.T(), err)
	require.NotEqual(s.T(), len(listed), 0)
	require.Equal(s.T(), found.Id, listed[len(listed)-1].Id)

	newname := util.Token()
	input = &model.Environment{Enabled: model.Bool(false), FeatureId: model.Int(5), Name: model.String(newname)}

	updated, err := s.client.EnvironmentUpdate(found.Id, input)
	require.NoError(s.T(), err)
	require.Equal(s.T(), found.Id, updated.Id)
	require.Equal(s.T(), model.Bool(false), updated.Enabled)
	require.Equal(s.T(), model.Int(5), updated.FeatureId)
	require.Equal(s.T(), model.String(newname), updated.Name)
}

func TestEnvironmentIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentIntegrationSuite))
}
