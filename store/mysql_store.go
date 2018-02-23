package store

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/MEDIGO/laika/models"
	"github.com/MEDIGO/laika/store/schema"
	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rubenv/sql-migrate"
)

type mySQLStore struct {
	db          *sql.DB
	mux         sync.Mutex
	lastEventID int
	state       *models.State
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

	return &mySQLStore{db: db, state: models.NewState()}, nil
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

	if _, err := migrate.Exec(s.db, "mysql", migrations, migrate.Up); err != nil {
		return err
	}

	return s.migrateData()
}

func (s *mySQLStore) Reset() error {
	tables := []string{
		"events",
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
		if _, err := s.db.Exec("DELETE FROM " + table); err != nil {
			return err
		}
	}

	if _, err := s.db.Exec("SET FOREIGN_KEY_CHECKS=1"); err != nil {
		return err
	}

	return nil
}

func (s *mySQLStore) Persist(eventType string, data string) (int64, error) {
	res, err := s.db.Exec(
		"INSERT INTO `events` (`time`, `type`, `data`) VALUES (NOW(), ?, ?)",
		eventType, data)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *mySQLStore) State() (*models.State, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	rows, err := s.db.Query(
		"SELECT `id`, `time`, `type`, `data` FROM `events` WHERE `id` > ? ORDER BY `time`, `id`",
		s.lastEventID)
	if err != nil {
		return s.state, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var time time.Time
		var eventType string
		var data string
		if err := rows.Scan(&id, &time, &eventType, &data); err != nil {
			return s.state, err
		}

		event, err := models.EventForType(eventType)
		if err != nil {
			return s.state, fmt.Errorf("error on event %d: %s", id, err)
		}

		if err := json.Unmarshal([]byte(data), event); err != nil {
			return s.state, fmt.Errorf("error on event %d: %s", id, err)
		}

		s.state = event.Update(s.state, time)
		s.lastEventID = id
	}

	return s.state, nil
}
