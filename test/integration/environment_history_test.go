package integration

import (
	"github.com/MEDIGO/feature-flag/model"
	"github.com/MEDIGO/feature-flag/util"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type EnvironmentHistoryIntegrationSuite struct {
	FeatureFlagSuite
}

func (s *EnvironmentHistoryIntegrationSuite) TestEnvironmentHistoryCRU() {

	name := util.Token()
	input := &model.EnvironmentHistory{
		CreatedAt: model.Time(time.Now()),
		Enabled:   model.Bool(true),
		FeatureId: model.Int(1),
		Name:      model.String(name),
		Timestamp: model.Time(time.Now()),
	}

	created, err := s.client.EnvironmentHistoryCreate(input)
	require.NoError(s.T(), err)
	require.NotEqual(s.T(), 0, created.Id)
	require.Equal(s.T(), model.Bool(true), created.Enabled)
	require.Equal(s.T(), model.Int(1), created.FeatureId)
	require.Equal(s.T(), model.String(name), created.Name)

	found, err := s.client.EnvironmentHistoryGet(created.Id)
	require.NoError(s.T(), err)
	require.Equal(s.T(), created.Id, found.Id)
	require.Equal(s.T(), created.Enabled, input.Enabled)
	require.Equal(s.T(), created.FeatureId, input.FeatureId)
	require.Equal(s.T(), created.Name, input.Name)

	listed, err := s.client.EnvironmentHistoryList()
	require.NoError(s.T(), err)
	require.NotEqual(s.T(), len(listed), 0)

	newname := util.Token()
	input = &model.EnvironmentHistory{Enabled: model.Bool(false), FeatureId: model.Int(5), Name: model.String(newname)}

	updated, err := s.client.EnvironmentHistoryUpdate(found.Id, input)
	require.NoError(s.T(), err)
	require.Equal(s.T(), found.Id, updated.Id)
	require.Equal(s.T(), model.Bool(false), updated.Enabled)
	require.Equal(s.T(), model.Int(5), updated.FeatureId)
	require.Equal(s.T(), model.String(newname), updated.Name)
}

func TestEnvironmentHistoryIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentHistoryIntegrationSuite))
}
