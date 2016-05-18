package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFeatureCacheAdd(t *testing.T) {
	cache := NewFeatureCache()

	feature := &Feature{
		Name: "awesome_feature",
	}

	cache.Add(feature)

	found := cache.Get("awesome_feature")
	require.NotNil(t, found)
	require.Equal(t, "awesome_feature", found.Name)
}

func TestFeatureCacheAddAll(t *testing.T) {
	cache := NewFeatureCache()

	features := []*Feature{
		&Feature{
			Name: "awesome_feature_1",
		},
		&Feature{
			Name: "awesome_feature_2",
		},
	}

	cache.AddAll(features)

	found := cache.Get("awesome_feature_1")
	require.NotNil(t, found)
	require.Equal(t, "awesome_feature_1", found.Name)

	found = cache.Get("awesome_feature_2")
	require.NotNil(t, found)
	require.Equal(t, "awesome_feature_2", found.Name)
}
