package model

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type EnvironmentIntegrationSuite struct {
	suite.Suite
}

func (suite *EnvironmentIntegrationSuite) TestExample() {
	name := "environmentName"
	fId := int64(1)
	exampleEnvironment := NewEnvironment(name, false, fId)
	require.NotNil(suite.T(), exampleEnvironment.Id)
	require.Equal(suite.T(), *exampleEnvironment.Name, name)
	require.Equal(suite.T(), *exampleEnvironment.Enabled, false)
	require.Equal(suite.T(), *exampleEnvironment.FeatureId, fId)
}

func TestEnvironmentIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentIntegrationSuite))
}
