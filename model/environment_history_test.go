package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type EnvironmentHistoryIntegrationSuite struct {
	suite.Suite
}

func (suite *EnvironmentHistoryIntegrationSuite) TestExample() {
	name := "featureKey"
	featureId := int64(1)
	enabled := true
	exampleEnvironmentHistory := NewEnvironmentHistory(time.Now(), enabled, featureId, name)
	require.NotNil(suite.T(), exampleEnvironmentHistory.Id)
	require.Equal(suite.T(), *exampleEnvironmentHistory.FeatureId, featureId)
	require.Equal(suite.T(), *exampleEnvironmentHistory.Name, name)
	require.Equal(suite.T(), *exampleEnvironmentHistory.Enabled, enabled)
}

func TestEnvironmentHistoryIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentHistoryIntegrationSuite))
}
