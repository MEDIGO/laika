package client

import (
	"sync"
	"bytes"
	"encoding/gob"
)

// FeatureCache is a in-memory threadsafe cache for Features.
type FeatureCache struct {
	features map[string]*Feature
	lock     sync.RWMutex
}

// NewFeatureCache creates a new FeatureCache.
func NewFeatureCache() *FeatureCache {
	return &FeatureCache{
		features: make(map[string]*Feature),
	}
}

// Add adds a feature to the cache.
func (fc *FeatureCache) Add(feature *Feature) {
	fc.lock.Lock()
	defer fc.lock.Unlock()

	fc.features[feature.Name] = feature
}

// AddAll adds a list of features to the cache.
func (fc *FeatureCache) AddAll(features []*Feature) {
	fc.lock.Lock()
	defer fc.lock.Unlock()

	for _, feature := range features {
		fc.features[feature.Name] = feature
	}
}

// Get gets a Feature from the cache if it exits.
func (fc *FeatureCache) Get(name string) *Feature {
	fc.lock.RLock()
	defer fc.lock.RUnlock()

	feature, ok := fc.features[name]
	if !ok {
		return nil
	}

	return feature
}

func (fc *FeatureCache) GetAll() map[string]*Feature {
	fc.lock.RLock()
	defer fc.lock.RUnlock()

	// deep clone features map
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	err := enc.Encode(fc.features)
	if err != nil {
		panic(err)
	}

	var deepCopy map[string]*Feature
	err = dec.Decode(&deepCopy)
	if err != nil {
		panic(err)
	}

	return deepCopy
}