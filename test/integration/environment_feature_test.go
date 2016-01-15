package integration

import (
  "testing"
  "github.com/MEDIGO/feature-flag/model"
  "github.com/stretchr/testify/require"
  "github.com/stretchr/testify/suite"
)

type EnvironmentFeatureIntegrationSuite struct {
  suite.Suite
  Key string
  FeatureID int64
  EnvironmentID int64
  ExampleEnvironmentFeature *model.EnvironmentFeature
}

func (suite *EnvironmentFeatureIntegrationSuite) SetupTest() {
  suite.Key = "featureKey"
  suite.FeatureID = 1
  suite.EnvironmentID = 2
  suite.ExampleEnvironmentFeature = model.NewEnvironmentFeature(&suite.FeatureID, &suite.EnvironmentID)
}

func (suite *EnvironmentFeatureIntegrationSuite) TestExample() {

	require.NotNil(suite.T(), suite.ExampleEnvironmentFeature.Id)
  require.Equal(suite.T(), *suite.ExampleEnvironmentFeature.FeatureID, suite.FeatureID)
  require.Equal(suite.T(), *suite.ExampleEnvironmentFeature.EnvironmentID, suite.EnvironmentID)
  require.Equal(suite.T(), *suite.ExampleEnvironmentFeature.Enabled, false)
  require.NotNil(suite.T(), suite.ExampleEnvironmentFeature.CreationDate)

}

func TestEnvironmentFeatureIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentFeatureIntegrationSuite))
}
