package model

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type EnvironmentIntegrationSuite struct {
	suite.Suite
	Name               string
	ExampleEnvironment *Environment
}

func (suite *EnvironmentIntegrationSuite) TestExample() {
	suite.Name = "environmentName"
	var fId int64 = 1
	suite.ExampleEnvironment = NewEnvironment(suite.Name, time.Now(), false, fId)
	require.NotNil(suite.T(), suite.ExampleEnvironment.Id)
	require.Equal(suite.T(), *suite.ExampleEnvironment.Name, suite.Name)
	require.Equal(suite.T(), *suite.ExampleEnvironment.Enabled, false)
	require.Equal(suite.T(), *suite.ExampleEnvironment.FeatureId, fId)
	require.NotNil(suite.T(), suite.ExampleEnvironment.CreatedAt)
}

func TestEnvironmentIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentIntegrationSuite))
}
