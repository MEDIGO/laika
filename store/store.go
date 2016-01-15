package store

import (
	"database/sql"
	"fmt"
	"github.com/MEDIGO/feature-flag/model"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var ErrNoRows = sql.ErrNoRows
var ErrTxDone = sql.ErrTxDone

type Store interface {
	GetFeatureById(id int64) (*model.Feature, error)
	ListFeatures(name *string, from *time.Time, to *time.Time) ([]*model.Feature, error)
	CreateFeature(feature *model.Feature) error
	UpdateFeature(feature *model.Feature) error

	GetEnvironmentById(id int64) (*model.Environment, error)
	ListEnvironments(name *string, featureId *int64, enabled *bool, from *time.Time, to *time.Time) ([]*model.Environment, error)
	CreateEnvironment(environment *model.Environment) error
	UpdateEnvironment(environment *model.Environment) error

	GetEnvironmentHistoryById(id int64) (*model.EnvironmentHistory, error)
	ListEnvironmentHistory(featureId *int64, name *string, enabled *bool, createdFrom *time.Time, createdUntil *time.Time, timestampFrom *time.Time, timestampUntil *time.Time) ([]*model.EnvironmentHistory, error)
	CreateEnvironmentHistory(environmentHistory *model.EnvironmentHistory) error
	UpdateEnvironmentHistory(environmentHistory *model.EnvironmentHistory) error

	Ping() error
}

type store struct {
	db *sql.DB
}

func NewStore(username, password, host, port, dbname string) (Store, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return &store{db}, nil
}

func (s *store) Ping() error {
	var err error

	for i := 0; i < 10; i++ {
		err = s.db.Ping()
		if err == nil {
			return nil
		}

		log.Warn("Failed to ping the database. Retry in 1s.")
		time.Sleep(time.Second)

	}

	return err
}
