package store

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/russross/meddler"
	sq "gopkg.in/Masterminds/squirrel.v1"
)

type FeatureStatusHistory struct {
	Id              int64      `json:"id"                            meddler:"id,pk"`
	CreatedAt       *time.Time `json:"created_at,omitempty"          meddler:"created_at"`
	Enabled         *bool      `json:"enabled,omitempty"             meddler:"enabled"`
	FeatureId       *int64     `json:"feature_id,omitempty"          meddler:"feature_id"`
	EnvironmentId   *int64     `json:"environment_id,omitempty"      meddler:"environment_id"`
	FeatureStatusId *int64     `json:"feature_status_id,omitempty"   meddler:"feature_status_id"`
	Timestamp       *time.Time `json:"timestamp,omitempty"           meddler:"timestamp"`
}

func NewFeatureStatusHistory(createdAt time.Time, enabled bool, featureId int64, environmentId int64, featureStatusId int64, timestamp time.Time) *FeatureStatusHistory {
	featureStatusHistory := new(FeatureStatusHistory)

	featureStatusHistory.CreatedAt = &createdAt
	featureStatusHistory.Enabled = &enabled
	featureStatusHistory.FeatureId = &featureId
	featureStatusHistory.EnvironmentId = &environmentId
	featureStatusHistory.FeatureStatusId = &featureStatusId
	featureStatusHistory.Timestamp = &createdAt

	return featureStatusHistory
}

func (s *store) ListFeatureStatusHistory(featureId *int64, environmentId *int64, featureStatusId *int64) ([]*FeatureStatusHistory, error) {
	query := sq.Select("*").From("feature_status_history")

	if featureId != nil {
		query = query.Where(sq.Eq{"feature_id": featureId})
	}

	if environmentId != nil {
		query = query.Where(sq.Eq{"environment_id": environmentId})
	}

	if featureStatusId != nil {
		query = query.Where(sq.Eq{"environment_id": environmentId})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	featuresStatusHistory := []*FeatureStatusHistory{}
	err = meddler.QueryAll(s.db, &featuresStatusHistory, sql, args...)

	return featuresStatusHistory, err
}

func (s *store) CreateFeatureStatusHistory(featureStatusHistory *FeatureStatusHistory) error {
	featureStatusHistory.CreatedAt = Time(time.Now())
	return meddler.Insert(s.db, "feature_status_history", featureStatusHistory)
}

func (s *store) UpdateFeatureStatusHistory(featureStatusHistory *FeatureStatusHistory) error {
	return meddler.Update(s.db, "feature_status_history", featureStatusHistory)
}
