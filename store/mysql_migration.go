package store

import (
	"encoding/json"
	"time"

	"github.com/MEDIGO/laika/models"
)

func (s *mySQLStore) migrateData() error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	rows, err := s.db.Query("SELECT 1 FROM gorp_migrations WHERE id='<DATA:0>'")
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return nil
	}

	type ev struct {
		time      time.Time
		eventType string
		data      string
	}
	events := []ev{}

	// migrate users
	rows, err = tx.Query("SELECT username, password_hash, created_at FROM user")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.UserCreated{}
		var createdAt time.Time
		if err := rows.Scan(&user.Username, &user.PasswordHash, &createdAt); err != nil {
			return err
		}
		data, _ := json.Marshal(&user)
		events = append(events, ev{
			time:      createdAt,
			eventType: "user_created",
			data:      string(data),
		})
	}

	// migrate envs
	rows, err = tx.Query("SELECT name, created_at FROM environment")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		env := models.EnvironmentCreated{}
		var createdAt time.Time
		if err := rows.Scan(&env.Name, &createdAt); err != nil {
			return err
		}
		data, _ := json.Marshal(&env)
		events = append(events, ev{
			time:      createdAt,
			eventType: "environment_created",
			data:      string(data),
		})
	}

	// migrate features
	rows, err = tx.Query("SELECT name, created_at FROM feature")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		feature := models.FeatureCreated{}
		var createdAt time.Time
		if err := rows.Scan(&feature.Name, &createdAt); err != nil {
			return err
		}
		data, _ := json.Marshal(&feature)
		events = append(events, ev{
			time:      createdAt,
			eventType: "feature_created",
			data:      string(data),
		})
	}

	// migrate status
	rows, err = tx.Query(`
		SELECT
			environment.name, feature.name, feature_status.enabled, feature_status.created_at
		FROM
			feature_status
		JOIN
			environment ON environment.id = feature_status.environment_id
		JOIN
			feature ON feature.id = feature_status.feature_id`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		toggle := models.FeatureToggled{}
		var createdAt time.Time
		if err := rows.Scan(&toggle.Environment, &toggle.Feature, &toggle.Status, &createdAt); err != nil {
			return err
		}
		data, _ := json.Marshal(&toggle)
		events = append(events, ev{
			time:      createdAt,
			eventType: "feature_toggled",
			data:      string(data),
		})
	}

	for _, event := range events {
		if _, err := tx.Exec("INSERT INTO `events` (`time`, `type`, `data`) VALUES (?, ?, ?)",
			event.time, event.eventType, event.data); err != nil {
			return err
		}
	}

	if _, err := tx.Exec("INSERT INTO gorp_migrations (id, applied_at) VALUES ('<DATA:0>', NOW())"); err != nil {
		return err
	}

	return tx.Commit()
}
