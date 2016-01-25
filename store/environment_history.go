package store

import (
	"github.com/MEDIGO/feature-flag/model"
	"github.com/russross/meddler"
	sq "gopkg.in/Masterminds/squirrel.v1"
	"time"
)

func (s *store) GetEnvironmentHistoryById(id int64) (*model.EnvironmentHistory, error) {
	environmentHistory := new(model.EnvironmentHistory)
	err := meddler.Load(s.db, "environment_history", environmentHistory, id)

	return environmentHistory, err
}

func (s *store) ListEnvironmentHistory(featureId *int64, name *string, enabled *bool, createdFrom *time.Time, createdUntil *time.Time, timestampFrom *time.Time, timestampUntil *time.Time) ([]*model.EnvironmentHistory, error) {
	query := sq.Select("*").From("environment_history")

	if featureId != nil {
		query = query.Where(sq.Eq{"feature_id": featureId})
	}

	if name != nil {
		query = query.Where(sq.Eq{"name": name})
	}

	if createdFrom != nil {
		query = query.Where(sq.Gt{"created_at": createdFrom})
	}

	if createdUntil != nil {
		query = query.Where(sq.Lt{"created_at": createdUntil})
	}

	if timestampFrom != nil {
		query = query.Where(sq.Gt{"timestamp": timestampFrom})
	}

	if timestampUntil != nil {
		query = query.Where(sq.Lt{"timestamp": timestampUntil})
	}

	if enabled != nil {
		query = query.Where(sq.Eq{"enabled": enabled})
	}

	sql, args, err := query.Limit(10).ToSql()
	if err != nil {
		return nil, err
	}

	environmentHistory := []*model.EnvironmentHistory{}
	err = meddler.QueryAll(s.db, &environmentHistory, sql, args...)

	return environmentHistory, err
}

func (s *store) CreateEnvironmentHistory(environmentHistory *model.EnvironmentHistory) error {
	return meddler.Insert(s.db, "environment_history", environmentHistory)
}

func (s *store) UpdateEnvironmentHistory(environmentHistory *model.EnvironmentHistory) error {
	return meddler.Update(s.db, "environment_history", environmentHistory)
}
