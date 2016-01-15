package integration

import (
  "testing"
  "github.com/MEDIGO/feature-flag/model"
  "github.com/stretchr/testify/require"
  "github.com/stretchr/testify/suite"
)

type FeatureIntegrationSuite struct {
  suite.Suite
  Key string
  ExampleFeature *model.Feature
}

func (suite *FeatureIntegrationSuite) SetupTest() {
  suite.Key = "featureKey"
  suite.ExampleFeature = model.NewFeature(&suite.Key)
}

func (suite *FeatureIntegrationSuite) TestExample() {

	require.NotNil(suite.T(), suite.ExampleFeature.Id)
	require.Equal(suite.T(), *suite.ExampleFeature.FeatureKey, suite.Key)
  require.NotNil(suite.T(), suite.ExampleFeature.CreationDate)

}

func TestFeatureIntegrationSuite(t *testing.T) {
	suite.Run(t, new(FeatureIntegrationSuite))
}
