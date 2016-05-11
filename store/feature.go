package store

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/russross/meddler"
	sq "github.com/Masterminds/squirrel"
)

type Feature struct {
	Id        int64      `json:"id"                         meddler:"id,pk"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Name      *string    `json:"name,omitempty"             meddler:"name"`
}

func (s *store) GetFeatureByName(name string) (*Feature, error) {
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

func (s *store) ListFeatures() ([]*Feature, error) {
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

func (s *store) CreateFeature(feature *Feature) error {
	feature.CreatedAt = Time(time.Now())
	return meddler.Insert(s.db, "feature", feature)
}

func (s *store) UpdateFeature(feature *Feature) error {
	return meddler.Update(s.db, "feature", feature)
}
