package store

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/russross/meddler"
	sq "github.com/Masterminds/squirrel"
)

type FeatureStatus struct {
	Id            int64      `json:"id"                         meddler:"id,pk"`
	CreatedAt     *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Enabled       *bool      `json:"enabled,omitempty"          meddler:"enabled"`
	FeatureId     *int64     `json:"feature_id,omitempty"       meddler:"feature_id"`
	EnvironmentId *int64     `json:"environment_id,omitempty"   meddler:"environment_id"`
}

func (s *store) GetFeatureStatus(featureId int64, environmentId int64) (*FeatureStatus, error) {
	featureStatus := new(FeatureStatus)

	query := sq.Select("*").From("feature_status")
	query = query.Where(sq.Eq{"feature_id": featureId})
	query = query.Where(sq.Eq{"environment_id": environmentId})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	err = meddler.QueryRow(s.db, featureStatus, sql, args...)

	return featureStatus, err
}

func (s *store) ListFeatureStatus(featureId *int64, environmentId *int64) ([]*FeatureStatus, error) {
	query := sq.Select("*").From("feature_status")

	if featureId != nil {
		query = query.Where(sq.Eq{"feature_id": featureId})
	}

	if environmentId != nil {
		query = query.Where(sq.Eq{"environment_id": environmentId})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	featuresStatus := []*FeatureStatus{}
	err = meddler.QueryAll(s.db, &featuresStatus, sql, args...)

	return featuresStatus, err
}

func (s *store) CreateFeatureStatus(featureStatus *FeatureStatus) error {
	featureStatus.CreatedAt = Time(time.Now())
	return meddler.Insert(s.db, "feature_status", featureStatus)
}

func (s *store) UpdateFeatureStatus(featureStatus *FeatureStatus) error {
	featureStatusHistory := &FeatureStatusHistory{
		CreatedAt:       Time(time.Now()),
		Enabled:         featureStatus.Enabled,
		FeatureId:       featureStatus.FeatureId,
		EnvironmentId:   featureStatus.EnvironmentId,
		FeatureStatusId: &featureStatus.Id,
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
		return
	}()

	if err := meddler.Insert(tx, "feature_status_history", featureStatusHistory); err != nil {
		return err
	}
	return meddler.Update(tx, "feature_status", featureStatus)
}
