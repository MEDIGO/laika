package store

import (
	"database/sql"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

var ErrNoRows = sql.ErrNoRows
var ErrTxDone = sql.ErrTxDone

type Store interface {
	GetFeatureByName(name string) (*Feature, error)
	ListFeatures() ([]*Feature, error)
	CreateFeature(feature *Feature) error
	UpdateFeature(feature *Feature) error

	GetEnvironmentByName(name string) (*Environment, error)
	ListEnvironments() ([]*Environment, error)
	CreateEnvironment(environment *Environment) error
	UpdateEnvironment(environment *Environment) error

	GetFeatureStatus(featureId int64, environmentId int64) (*FeatureStatus, error)
	ListFeatureStatus(featureId *int64, environmentId *int64) ([]*FeatureStatus, error)
	CreateFeatureStatus(featureStatus *FeatureStatus) error
	UpdateFeatureStatus(featureStatus *FeatureStatus) error

	ListFeatureStatusHistory(featureId *int64, environmentId *int64, featureStatusId *int64) ([]*FeatureStatusHistory, error)
	CreateFeatureStatusHistory(featureStatusHistory *FeatureStatusHistory) error
	UpdateFeatureStatusHistory(featureStatusHistory *FeatureStatusHistory) error

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
