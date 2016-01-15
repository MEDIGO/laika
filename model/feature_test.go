package model

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type FeatureIntegrationSuite struct {
	suite.Suite
	Name           string
	ExampleFeature *Feature
}

func (suite *FeatureIntegrationSuite) TestExample() {
	suite.Name = "featureKey"
	suite.ExampleFeature = NewFeature(suite.Name, time.Now())
	require.NotNil(suite.T(), suite.ExampleFeature.Id)
	require.Equal(suite.T(), *suite.ExampleFeature.Name, suite.Name)
	require.NotNil(suite.T(), suite.ExampleFeature.CreatedAt)
}

func TestFeatureIntegrationSuite(t *testing.T) {
	suite.Run(t, new(FeatureIntegrationSuite))
}
