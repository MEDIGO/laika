package integration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/MEDIGO/laika/api"
	"github.com/MEDIGO/laika/store"
	"github.com/MEDIGO/laika/util"
)

type EnvironmentIntegrationSuite struct {
	FeatureFlagSuite
}

func (s *EnvironmentIntegrationSuite) TestEnvironmentCRUD() {
	name := util.Token()

	input := &api.Environment{
		CreatedAt: store.Time(time.Now()),
		Name:      store.String(name),
	}

	created, err := s.client.EnvironmentCreate(input)
	require.NoError(s.T(), err)
	require.NotEqual(s.T(), 0, created.Id)
	require.Equal(s.T(), store.String(name), created.Name)

	found, err := s.client.EnvironmentGet(*created.Name)
	require.NoError(s.T(), err)
	require.Equal(s.T(), created.Id, found.Id)
	require.Equal(s.T(), created.Name, input.Name)

	listed, err := s.client.EnvironmentList()
	require.NoError(s.T(), err)
	require.NotEqual(s.T(), len(listed), 0)
	require.Equal(s.T(), found.Id, listed[len(listed)-1].Id)

	newName := util.Token()
	input = &api.Environment{Name: store.String(newName)}

	updated, err := s.client.EnvironmentUpdate(*found.Name, input)
	require.NoError(s.T(), err)
	require.Equal(s.T(), found.Id, updated.Id)
	require.Equal(s.T(), store.String(newName), updated.Name)
}

func TestEnvironmentIntegrationSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentIntegrationSuite))
}
