package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FeatureStatusIntegrationSuite struct {
	suite.Suite
}

func (suite *FeatureStatusIntegrationSuite) TestExample() {
	fId := int64(1)
	eId := int64(2)
	enabled := true
	exampleFeatureStatus := NewFeatureStatus(time.Now(), enabled, fId, eId)
	require.NotNil(suite.T(), exampleFeatureStatus.Id)
	require.Equal(suite.T(), *exampleFeatureStatus.Enabled, enabled)
	require.Equal(suite.T(), *exampleFeatureStatus.FeatureId, fId)
	require.Equal(suite.T(), *exampleFeatureStatus.EnvironmentId, eId)
}

func TestFeatureStatusIntegrationSuite(t *testing.T) {
	suite.Run(t, new(FeatureStatusIntegrationSuite))
}
