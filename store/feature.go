package store

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/russross/meddler"
	sq "gopkg.in/Masterminds/squirrel.v1"
)

type Feature struct {
	Id        int64      `json:"id"                         meddler:"id,pk"`
	CreatedAt *time.Time `json:"created_at,omitempty"       meddler:"created_at"`
	Name      *string    `json:"name,omitempty"             meddler:"name"`
}

func NewFeature(name string) *Feature {
	feature := new(Feature)

	feature.Name = &name

	return feature
}

func (f *Feature) Validate() error {
	if f.Name == nil {
		return CustomError{
			"Name: non zero value required;",
		}
	}
	return nil
}

func (s *store) GetFeatureById(id int64) (*Feature, error) {
	feature := new(Feature)
	err := meddler.Load(s.db, "feature", feature, id)

	return feature, err
}

func (s *store) ListFeatures() ([]Feature, error) {
	query := sq.Select("*").From("feature")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	log.Debug(sql)

	features := []Feature{}
	err = meddler.QueryAll(s.db, &features, sql, args...)

	return features, err
}

func (s *store) CreateFeature(feature Feature) error {
	feature.CreatedAt = Time(time.Now())
	return meddler.Insert(s.db, "feature", feature)
}

func (s *store) UpdateFeature(feature Feature) error {
	return meddler.Update(s.db, "feature", feature)
}
