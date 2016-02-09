package store

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type FeatureIntegrationSuite struct {
	suite.Suite
}

func (suite *FeatureIntegrationSuite) TestExample() {
	name := "featureKey"
	exampleFeature := NewFeature(name)
	require.NotNil(suite.T(), exampleFeature.Id)
	require.Equal(suite.T(), *exampleFeature.Name, name)
}

func TestFeatureIntegrationSuite(t *testing.T) {
	suite.Run(t, new(FeatureIntegrationSuite))
}
