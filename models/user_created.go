package models

import (
	"errors"
	"time"

	"github.com/MEDIGO/laika/notifier"
	"golang.org/x/crypto/bcrypt"
)

type UserCreated struct {
	Username     string  `json:"username"`
	Password     *string `json:"password,omitempty"`
	PasswordHash string  `json:"password_hash"`
}

func (e *UserCreated) Validate(s *State) (error, error) {
	if e.Username == "" {
		return errors.New("Username must not be empty"), nil
	}

	if s.getUserByName(e.Username) != nil {
		return errors.New("Name is already in use"), nil
	}

	if e.Password == nil && e.PasswordHash == "" ||
		e.Password != nil && e.PasswordHash != "" {
		return errors.New("Exactly one of either password or password hash is required"), nil
	}

	return nil, nil
}

func (e *UserCreated) Update(s *State, t time.Time) *State {
	state := *s

	state.Users = append(state.Users, User{
		Username:     e.Username,
		PasswordHash: e.PasswordHash,
	})

	return &state
}

func (e *UserCreated) PrePersist(s *State) (Event, error) {
	if e.Password == nil {
		// already hashed
		return e, nil
	}

	event := *e

	hash, err := bcrypt.GenerateFromPassword([]byte(*e.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// discard plain text
	event.Password = nil
	event.PasswordHash = string(hash)

	return &event, nil
}

func (e *UserCreated) Notify(*State, notifier.Notifier) {
}
