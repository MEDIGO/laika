package store

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
	exampleEnvironment := NewEnvironment(name)
	require.NotNil(suite.T(), exampleEnvironment.Id)
	require.Equal(suite.T(), *exampleEnvironment.Name, name)
}

func TestEnvironmentIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentIntegrationSuite))
}
