package store

import (
	"time"

	"github.com/russross/meddler"

	"github.com/MEDIGO/feature-flag/model"
)

func (s *store) CreateEnvironmentHistory(environment *model.Environment) error {
	environmentHistory := model.NewEnvironmentHistory(*environment.CreatedAt, *environment.Enabled, *environment.FeatureId, *environment.Name)

	if err := environmentHistory.Validate(); err != nil {
		return err
	}

	environmentHistory.Timestamp = model.Time(time.Now())

	return meddler.Insert(s.db, "environment_history", environmentHistory)
}
