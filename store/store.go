package store

import (
	"database/sql"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"

	"github.com/MEDIGO/feature-flag/model"
)

var ErrNoRows = sql.ErrNoRows
var ErrTxDone = sql.ErrTxDone

type Store interface {
	GetFeatureById(id int64) (*model.Feature, error)
	GetFeatureByName(name string) (*model.Feature, error)
	ListFeatures(name *string, from *time.Time, to *time.Time) ([]*model.Feature, error)
	CreateFeature(feature *model.Feature) error
	UpdateFeature(feature *model.Feature) error

	GetEnvironment(name string, featureId int64) (*model.Environment, error)
	GetEnvironmentById(id int64) (*model.Environment, error)
	ListEnvironments(name *string, featureId *int64, enabled *bool, from *time.Time, to *time.Time) ([]*model.Environment, error)
	CreateEnvironment(environment *model.Environment) error
	UpdateEnvironment(environment *model.Environment) error

	CreateEnvironmentHistory(environment *model.Environment) error

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
