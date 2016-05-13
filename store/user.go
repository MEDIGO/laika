package store

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	sq "github.com/Masterminds/squirrel"
	log "github.com/Sirupsen/logrus"
	"github.com/russross/meddler"
)

type User struct {
	ID           int64      `meddler:"id,pk"`
	Username     string     `meddler:"username"`
	PasswordHash string     `meddler:"password_hash"`
	CreatedAt    time.Time  `meddler:"created_at"`
	UpdatedAt    *time.Time `meddler:"updated_at"`
}

func (s *store) GetUserByUsername(username string) (*User, error) {
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

func (s *store) CreateUser(user *User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(passwordHash)

	user.CreatedAt = time.Now()
	return meddler.Insert(s.db, "user", user)
}
