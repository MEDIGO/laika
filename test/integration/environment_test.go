package integration

import (
  "testing"
  "github.com/MEDIGO/feature-flag/model"
  "github.com/stretchr/testify/require"
  "github.com/stretchr/testify/suite"
)

type EnvironmentIntegrationSuite struct {
  suite.Suite
  Key string
  ExampleEnvironment *model.Environment
}

func (suite *EnvironmentIntegrationSuite) SetupTest() {
  suite.Key = "environmentKey"
  suite.ExampleEnvironment = model.NewEnvironment(&suite.Key)
}

func (suite *EnvironmentIntegrationSuite) TestExample() {

	require.NotNil(suite.T(), suite.ExampleEnvironment.Id)
	require.Equal(suite.T(), *suite.ExampleEnvironment.EnvironmentKey, suite.Key)
  require.Equal(suite.T(), *suite.ExampleEnvironment.Enabled, false)
  require.NotNil(suite.T(), suite.ExampleEnvironment.CreationDate)

}

func TestEnvironmentIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentIntegrationSuite))
}
