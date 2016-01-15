package integration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/MEDIGO/feature-flag/model"
	"github.com/MEDIGO/feature-flag/util"
)

type EnvironmentIntegrationSuite struct {
	FeatureFlagSuite
}

func (s *EnvironmentIntegrationSuite) TestEnvironmentCRU() {
	name := util.Token()
	inputFeature := &model.Feature{
		Name: model.String(name),
	}

	createdFeature, featureErr := s.client.FeatureCreate(inputFeature)
	require.NoError(s.T(), featureErr)
	require.NotEqual(s.T(), 0, createdFeature.Id)
	require.Equal(s.T(), model.String(name), createdFeature.Name)

	name = util.Token()
	input := &model.Environment{
		CreatedAt: model.Time(time.Now()),
		Enabled:   model.Bool(true),
		FeatureId: model.Int(createdFeature.Id),
		Name:      model.String(name),
	}

	created, err := s.client.EnvironmentCreate(input)
	require.NoError(s.T(), err)
	require.NotEqual(s.T(), 0, created.Id)
	require.Equal(s.T(), model.Bool(true), created.Enabled)
	require.Equal(s.T(), model.Int(createdFeature.Id), created.FeatureId)
	require.Equal(s.T(), model.String(name), created.Name)

	found, err := s.client.EnvironmentGet(*createdFeature.Name, *created.Name)
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
	input = &model.Environment{Enabled: model.Bool(false), Name: model.String(newname)}

	updated, err := s.client.EnvironmentUpdate(found.Id, input)
	require.NoError(s.T(), err)
	require.Equal(s.T(), found.Id, updated.Id)
	require.Equal(s.T(), model.Bool(false), updated.Enabled)
	require.Equal(s.T(), model.String(newname), updated.Name)
}

func TestEnvironmentIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentIntegrationSuite))
}
