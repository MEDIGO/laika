package store

import (
	"database/sql"
	"fmt"
	"time"
  _ "github.com/go-sql-driver/mysql"
	"github.com/MEDIGO/feature-flag/model"
)

//TODO ??? - REMOVE IF UNNECESSARY
//var ErrNoRows = sql.ErrNoRows
//var ErrTxDone = sql.ErrTxDone

type Store interface {
  GetFeatureByID(id int64) (*model.Feature, error)
  ListFeatures(featureKey *string, from *time.Time, to *time.Time) ([]*model.Feature ,error)
  CreateFeature(featureKey *string) error

  GetEnvironmentByID(id int64) (*model.Environment, error)
  ListEnvironments(environmentKey *string, enabled *bool , from *time.Time, to *time.Time) ([]*model.Environment ,error)
  CreateEnvironment(environmentKey *string) error
  ChangeEnvironmentState(environment *model.Environment) error

  GetEnvironmentFeatureByID(id int64) (*model.EnvironmentFeature, error)
  ListEnvironmentFeatures(environmentID *int64, featureID *int64, enabled *bool, from *time.Time, to *time.Time) ([]*model.EnvironmentFeature ,error)
  CreateEnvironmentFeature(featureID, environmentID *int64) error
  ChangeEnvironmentFeatureState(environmentFeature *model.EnvironmentFeature) error

  GetEnvironmentFeatureUpdateByID(id int64) (*model.EnvironmentFeatureUpdate, error)
  ListEnvironmentFeatureUpdates(featureID *int64, environmentID *int64, enabled *bool , from *time.Time, to *time.Time) ([]*model.EnvironmentFeatureUpdate ,error)
  CreateEnvironmentFeatureUpdate(featureID, environmentID *int64, enabled *bool) error

  Ping() error
}

type store struct {
  db *sql.DB
}

func NewStore(username, password, host, port, dbname string) (Store, error){
  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbname)
  db, err := sql.Open("mysql", dsn)

  if err != nil {
    return nil, err
  }

  return &store{db}, nil
}

func (s *store) Ping() error {
  var err error

  for i := 0; i < 10 ; i++ {
    err = s.db.Ping()
    if err == nil {
      return nil
    }

    fmt.Printf("Failed to ping the database. Retry in 1s.") //TODO REMOVE
    time.Sleep(time.Second)

  }

  return err
}
