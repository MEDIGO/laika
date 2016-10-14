package store

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MEDIGO/laika/models"
	"github.com/MEDIGO/laika/store/schema"
	sq "github.com/Masterminds/squirrel"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rubenv/sql-migrate"
	"github.com/russross/meddler"
	"golang.org/x/crypto/bcrypt"
)

type status struct {
	ID            int64     `meddler:"id,pk"`
	CreatedAt     time.Time `meddler:"created_at"`
	Enabled       bool      `meddler:"enabled"`
	FeatureID     int64     `meddler:"feature_id"`
	EnvironmentID int64     `meddler:"environment_id"`
}

type mySQLStore struct {
	db *sql.DB
}

// NewMySQLStore creates a new store that uses MySQL as a backend.
func NewMySQLStore(username, password, host, port, dbname string) (Store, error) {
	creds := username
	if password != "" {
		creds += ":" + password
	}

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", creds, host, port, dbname)
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

func (s *mySQLStore) GetFeatureByName(name string) (*models.Feature, error) {
	feature := new(models.Feature)

	query := sq.Select("*").From("feature").Where(sq.Eq{"name": name})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	if err := meddler.QueryRow(s.db, feature, sql, args...); err != nil {
		return nil, err
	}

	envs, err := s.ListEnvironments()
	if err != nil {
		return nil, err
	}

	if len(envs) == 0 {
		return feature, nil
	}

	feature.Status = make(map[string]bool)
	envNames := make(map[int64]string)
	for _, env := range envs {
		feature.Status[env.Name] = false
		envNames[env.ID] = env.Name
	}

	status, err := s.listStatusByFeatureID(feature.ID)
	if err != nil {
		return nil, err
	}

	for _, st := range status {
		name := envNames[st.EnvironmentID]
		if name == "" {
			continue
		}

		feature.Status[name] = st.Enabled
	}

	return feature, nil
}

func (s *mySQLStore) ListFeatures() ([]*models.Feature, error) {
	query := sq.Select("*").From("feature")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	features := []*models.Feature{}
	err = meddler.QueryAll(s.db, &features, sql, args...)
	if err != nil {
		return nil, err
	}

	featuresByID := make(map[int64]*models.Feature)
	envsByID := make(map[int64]*models.Environment)

	envs, err := s.ListEnvironments()
	if err != nil {
		return nil, err
	}

	for _, env := range envs {
		envsByID[env.ID] = env
	}

	for _, feature := range features {
		featuresByID[feature.ID] = feature
		feature.Status = make(map[string]bool)
		for _, env := range envs {
			feature.Status[env.Name] = false
		}
	}

	stats, err := s.listStatus()
	if err != nil {
		return nil, err
	}

	for _, stat := range stats {
		feature := featuresByID[stat.FeatureID]
		env := envsByID[stat.EnvironmentID]

		feature.Status[env.Name] = stat.Enabled
	}

	return features, err
}

func (s *mySQLStore) CreateFeature(feature *models.Feature) error {
	feature.CreatedAt = time.Now()

	if err := meddler.Insert(s.db, "feature", feature); err != nil {
		return err
	}

	// update the feature to create all the status
	return s.UpdateFeature(feature)
}

func (s *mySQLStore) UpdateFeature(feature *models.Feature) error {
	if err := meddler.Update(s.db, "feature", feature); err != nil {
		return err
	}

	envs, err := s.ListEnvironments()
	if err != nil {
		return err
	}

	envsByName := make(map[string]*models.Environment)
	for _, env := range envs {
		envsByName[env.Name] = env
	}

	stats, err := s.listStatusByFeatureID(feature.ID)
	if err != nil {
		return err
	}

	statusByEnvironmentID := make(map[int64]*status)
	for _, stat := range stats {
		statusByEnvironmentID[stat.EnvironmentID] = stat
	}

	for envName, enabled := range feature.Status {
		env := envsByName[envName]
		if env == nil {
			return ErrNoRows
		}

		stat := statusByEnvironmentID[env.ID]
		if stat == nil {
			err := s.createStatus(&status{
				FeatureID:     feature.ID,
				EnvironmentID: env.ID,
				Enabled:       enabled,
			})
			if err != nil {
				return err
			}
		} else {
			if stat.Enabled == enabled {
				// no changes
				continue
			}

			stat.Enabled = enabled
			err := s.updateStatus(stat)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *mySQLStore) GetUserByUsername(username string) (*models.User, error) {
	user := new(models.User)

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

func (s *mySQLStore) CreateUser(user *models.User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = ""
	user.PasswordHash = string(passwordHash)

	user.CreatedAt = time.Now()
	return meddler.Insert(s.db, "user", user)
}

func (s *mySQLStore) GetEnvironmentByName(name string) (*models.Environment, error) {
	environment := new(models.Environment)

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

func (s *mySQLStore) ListEnvironments() ([]*models.Environment, error) {
	query := sq.Select("*").From("environment")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	environments := []*models.Environment{}
	err = meddler.QueryAll(s.db, &environments, sql, args...)

	return environments, err
}

func (s *mySQLStore) CreateEnvironment(environment *models.Environment) error {
	environment.CreatedAt = time.Now()
	return meddler.Insert(s.db, "environment", environment)
}

func (s *mySQLStore) UpdateEnvironment(environment *models.Environment) error {
	return meddler.Update(s.db, "environment", environment)
}

func (s *mySQLStore) listStatusByFeatureID(featureID int64) ([]*status, error) {
	query := sq.Select("*").From("feature_status").Where(sq.Eq{"feature_id": featureID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	status := []*status{}
	err = meddler.QueryAll(s.db, &status, sql, args...)

	return status, err
}

func (s *mySQLStore) listStatus() ([]*status, error) {
	query := sq.Select("*").From("feature_status")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	status := []*status{}
	err = meddler.QueryAll(s.db, &status, sql, args...)

	return status, err
}

func (s *mySQLStore) createStatus(status *status) error {
	status.CreatedAt = time.Now()
	return meddler.Insert(s.db, "feature_status", status)
}

func (s *mySQLStore) updateStatus(status *status) error {
	return meddler.Update(s.db, "feature_status", status)
}
