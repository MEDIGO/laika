package store

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MEDIGO/laika/store/schema"
	sq "github.com/Masterminds/squirrel"
	log "github.com/Sirupsen/logrus"
	"github.com/rubenv/sql-migrate"
	"github.com/russross/meddler"
	"golang.org/x/crypto/bcrypt"
)

type mySQLStore struct {
	db *sql.DB
}

// NewMySQLStore creates a new store that uses MySQL as a backend.
func NewMySQLStore(username, password, host, port, dbname string) (Store, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbname)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return &mySQLStore{db}, nil
}

func (s *mySQLStore) Ping() error {
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

func (s *mySQLStore) Migrate() error {
	migrations := &migrate.AssetMigrationSource{
		Asset:    schema.Asset,
		AssetDir: schema.AssetDir,
		Dir:      "store/schema",
	}

	_, err := migrate.Exec(s.db, "mysql", migrations, migrate.Up)
	return err
}

func (s *mySQLStore) Reset() error {
	tables := []string{
		"feature_status_history",
		"feature_status",
		"environment",
		"feature",
		"user",
	}

	if _, err := s.db.Exec("SET FOREIGN_KEY_CHECKS=0"); err != nil {
		return err
	}

	for _, table := range tables {
		if _, err := s.db.Exec("TRUNCATE TABLE " + table); err != nil {
			return err
		}
	}

	if _, err := s.db.Exec("SET FOREIGN_KEY_CHECKS=1"); err != nil {
		return err
	}

	return nil
}

func (s *mySQLStore) GetFeatureByName(name string) (*Feature, error) {
	feature := new(Feature)

	query := sq.Select("*").From("feature")
	query = query.Where(sq.Eq{"name": name})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	err = meddler.QueryRow(s.db, feature, sql, args...)

	return feature, err
}

func (s *mySQLStore) ListFeatures() ([]*Feature, error) {
	query := sq.Select("*").From("feature")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	features := []*Feature{}
	err = meddler.QueryAll(s.db, &features, sql, args...)

	return features, err
}

func (s *mySQLStore) CreateFeature(feature *Feature) error {
	feature.CreatedAt = Time(time.Now())
	return meddler.Insert(s.db, "feature", feature)
}

func (s *mySQLStore) UpdateFeature(feature *Feature) error {
	return meddler.Update(s.db, "feature", feature)
}

func (s *mySQLStore) GetUserByUsername(username string) (*User, error) {
	user := new(User)

	query := sq.Select("*").From("user")
	query = query.Where(sq.Eq{"username": username})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	err = meddler.QueryRow(s.db, user, sql, args...)

	return user, err
}

func (s *mySQLStore) CreateUser(user *User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(passwordHash)

	user.CreatedAt = time.Now()
	return meddler.Insert(s.db, "user", user)
}

func (s *mySQLStore) GetFeatureStatus(featureId int64, environmentId int64) (*FeatureStatus, error) {
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

func (s *mySQLStore) ListFeatureStatus(featureId *int64, environmentId *int64) ([]*FeatureStatus, error) {
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

func (s *mySQLStore) CreateFeatureStatus(featureStatus *FeatureStatus) error {
	featureStatus.CreatedAt = Time(time.Now())
	return meddler.Insert(s.db, "feature_status", featureStatus)
}

func (s *mySQLStore) UpdateFeatureStatus(featureStatus *FeatureStatus) error {
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

func (s *mySQLStore) ListFeatureStatusHistory(featureId *int64, environmentId *int64, featureStatusId *int64) ([]*FeatureStatusHistory, error) {
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

func (s *mySQLStore) GetEnvironmentByName(name string) (*Environment, error) {
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

func (s *mySQLStore) ListEnvironments() ([]*Environment, error) {
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

func (s *mySQLStore) CreateEnvironment(environment *Environment) error {
	environment.CreatedAt = Time(time.Now())
	return meddler.Insert(s.db, "environment", environment)
}

func (s *mySQLStore) UpdateEnvironment(environment *Environment) error {
	return meddler.Update(s.db, "environment", environment)
}
