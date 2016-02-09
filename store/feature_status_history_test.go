package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FeatureStatusHistoryIntegrationSuite struct {
	suite.Suite
}

func (suite *FeatureStatusHistoryIntegrationSuite) TestExample() {
	fId := int64(2)
	eId := int64(2)
	fsId := int64(3)
	enabled := true
	exampleFeatureStatusHistory := NewFeatureStatusHistory(time.Now(), enabled, fId, eId, fsId, time.Now())
	require.NotNil(suite.T(), exampleFeatureStatusHistory.Id)
	require.Equal(suite.T(), *exampleFeatureStatusHistory.Enabled, enabled)
	require.Equal(suite.T(), *exampleFeatureStatusHistory.FeatureId, fId)
	require.Equal(suite.T(), *exampleFeatureStatusHistory.EnvironmentId, eId)
	require.Equal(suite.T(), *exampleFeatureStatusHistory.FeatureStatusId, fsId)
}

func TestFeatureStatusHistoryIntegrationSuite(t *testing.T) {
	suite.Run(t, new(FeatureStatusHistoryIntegrationSuite))
}
