package store

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/russross/meddler"
	sq "github.com/Masterminds/squirrel"
)

type Environment struct {
	Id        int64      `json:"id"                         meddler:"id,pk"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Name      *string    `json:"name,omitempty"             meddler:"name"`
}

func (s *store) GetEnvironmentByName(name string) (*Environment, error) {
	environment := new(Environment)

	query := sq.Select("*").From("environment")
	query = query.Where(sq.Eq{"name": name})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	err = meddler.QueryRow(s.db, environment, sql, args...)

	return environment, err
}

func (s *store) ListEnvironments() ([]*Environment, error) {
	query := sq.Select("*").From("environment")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	environments := []*Environment{}
	err = meddler.QueryAll(s.db, &environments, sql, args...)

	return environments, err
}

func (s *store) CreateEnvironment(environment *Environment) error {
	environment.CreatedAt = Time(time.Now())
	return meddler.Insert(s.db, "environment", environment)
}

func (s *store) UpdateEnvironment(environment *Environment) error {
	return meddler.Update(s.db, "environment", environment)
}
