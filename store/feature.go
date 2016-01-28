package store

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/russross/meddler"
	sq "gopkg.in/Masterminds/squirrel.v1"

	"github.com/MEDIGO/feature-flag/model"
)

func (s *store) GetFeatureById(id int64) (*model.Feature, error) {
	feature := new(model.Feature)
	err := meddler.Load(s.db, "feature", feature, id)

	return feature, err
}

func (s *store) GetFeatureByName(name string) (*model.Feature, error) {
	feature := new(model.Feature)

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

func (s *store) ListFeatures(name *string, from *time.Time, to *time.Time) ([]*model.Feature, error) {
	query := sq.Select("*").From("feature")

	if name != nil {
		query = query.Where(sq.Eq{"name": name})
	}

	if from != nil {
		query = query.Where(sq.Gt{"created_at": from})
	}

	if to != nil {
		query = query.Where(sq.Lt{"created_at": to})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	features := []*model.Feature{}
	err = meddler.QueryAll(s.db, &features, sql, args...)

	return features, err
}

func (s *store) CreateFeature(feature *model.Feature) error {
	feature.CreatedAt = model.Time(time.Now())
	return meddler.Insert(s.db, "feature", feature)
}

func (s *store) UpdateFeature(feature *model.Feature) error {
	return meddler.Update(s.db, "feature", feature)
}
