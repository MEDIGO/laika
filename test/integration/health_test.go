package integration

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type HealthIntegrationSuite struct {
	FeatureFlagSuite
}

func (s *HealthIntegrationSuite) TestHealthCheck() {
	err := s.client.HealthCheck()
	require.NoError(s.T(), err)
}

func TestHealthIntegrationSuite(t *testing.T) {
	suite.Run(t, new(HealthIntegrationSuite))
}
