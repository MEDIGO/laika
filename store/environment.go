package store

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/russross/meddler"
	sq "gopkg.in/Masterminds/squirrel.v1"

	"github.com/MEDIGO/feature-flag/model"
)

func (s *store) GetEnvironmentById(id int64) (*model.Environment, error) {
	environment := new(model.Environment)
	err := meddler.Load(s.db, "environment", environment, id)

	return environment, err
}

func (s *store) GetEnvironment(name string, featureId int64) (*model.Environment, error) {
	environment := new(model.Environment)

	query := sq.Select("*").From("environment")
	query = query.Where(sq.Eq{"name": name})
	query = query.Where(sq.Eq{"feature_id": featureId})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	err = meddler.QueryRow(s.db, environment, sql, args...)

	return environment, err
}

func (s *store) ListEnvironments(name *string, featureId *int64, enabled *bool, from *time.Time, to *time.Time) ([]*model.Environment, error) {
	query := sq.Select("*").From("environment")

	if name != nil {
		query = query.Where(sq.Eq{"name": name})
	}

	if featureId != nil {
		query = query.Where(sq.Eq{"feature_id": featureId})
	}

	if from != nil {
		query = query.Where(sq.Gt{"created_at": from})
	}

	if to != nil {
		query = query.Where(sq.Lt{"created_at": to})
	}

	if enabled != nil {
		query = query.Where(sq.Eq{"enabled": enabled})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	environments := []*model.Environment{}
	err = meddler.QueryAll(s.db, &environments, sql, args...)

	return environments, err
}

func (s *store) CreateEnvironment(environment *model.Environment) error {
	environment.CreatedAt = model.Time(time.Now())
	return meddler.Insert(s.db, "environment", environment)
}

func (s *store) UpdateEnvironment(environment *model.Environment) error {
	return meddler.Update(s.db, "environment", environment)
}
