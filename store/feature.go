package store

//TODO replace prints with logs
//TODO remove result prints

import (
  "github.com/russross/meddler"
  sq "gopkg.in/Masterminds/squirrel.v1"
  "time"
  "fmt"
  "github.com/MEDIGO/feature-flag/model"
)

func (s *store) GetFeatureByID(id int64) (*model.Feature, error) {
  feature := new(model.Feature)
  err := meddler.Load(s.db, "feature", feature, id)

  fmt.Printf("%d %s %s\n", feature.Id, *feature.FeatureKey, feature.CreationDate)

  return feature, err
}

func (s *store) ListFeatures(featureKey *string, from *time.Time, to *time.Time) ([]*model.Feature ,error){
  query := sq.Select("*").From("feature")

  if featureKey != nil {
		query = query.Where(sq.Eq{"feature_key": featureKey})
	}

  switch {
    case from != nil && to != nil:
      fmt.Printf("ONE\n")
      query = query.Where(sq.Gt{"creation_date": from})
      query = query.Where(sq.Lt{"creation_date": to})
      break
    case from != nil:
      fmt.Printf("TWO\n")
      query = query.Where(sq.Gt{"creation_date": from})
      break
    case to != nil:
      fmt.Printf("THREE\n")
      query = query.Where(sq.Lt{"creation_date": to})
      break
  }

	sql, args, err := query.Limit(10).ToSql()
  if err != nil {
    return nil, err
  }

  fmt.Printf(sql+"\n")

	features := []*model.Feature{}
	err = meddler.QueryAll(s.db, &features, sql, args...)

  fmt.Printf("%d\n",len(features))

  for i:=0; i<len(features); i++ {
    fmt.Printf("%d %s %s\n", features[i].Id, *features[i].FeatureKey, features[i].CreationDate)
  }

  return features, err
}

func (s *store) CreateFeature(featureKey *string) error {
  feature := model.NewFeature(featureKey)
  return meddler.Insert(s.db, "feature", feature)
}
