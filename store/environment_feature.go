package store

//TODO replace prints with logs
//TODO remove result prints

import (
  "fmt"
  "github.com/russross/meddler"
  sq "gopkg.in/Masterminds/squirrel.v1"
  "time"
  "github.com/MEDIGO/feature-flag/model"
)

func (s *store) GetEnvironmentFeatureByID(id int64) (*model.EnvironmentFeature, error) {
  environmentFeature := new(model.EnvironmentFeature)
  err := meddler.Load(s.db, "environment_feature", environmentFeature, id)

  fmt.Printf("%d %d %d %t %s\n", environmentFeature.Id, *environmentFeature.FeatureID, *environmentFeature.EnvironmentID, *environmentFeature.Enabled, environmentFeature.CreationDate)

  return environmentFeature, err
}

func (s *store) ListEnvironmentFeatures(featureID *int64, environmentID *int64, enabled *bool , from *time.Time, to *time.Time) ([]*model.EnvironmentFeature ,error){
  query := sq.Select("*").From("environment_feature")

  if featureID != nil {
		query = query.Where(sq.Eq{"feature_id": featureID})
	}

  if environmentID != nil {
		query = query.Where(sq.Eq{"environment_id": environmentID})
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

  if enabled != nil {
		query = query.Where(sq.Eq{"enabled": enabled})
	}

	sql, args, err := query.Limit(10).ToSql()
  if err != nil {
    return nil, err
  }

  fmt.Printf(sql+"\n")

	environmentFeatures := []*model.EnvironmentFeature{}
	err = meddler.QueryAll(s.db, &environmentFeatures, sql, args...)

  fmt.Printf("%d\n",len(environmentFeatures))

  for i:=0; i<len(environmentFeatures); i++ {
    fmt.Printf("%d %d %d %t %s\n", environmentFeatures[i].Id, *environmentFeatures[i].FeatureID, *environmentFeatures[i].EnvironmentID, *environmentFeatures[i].Enabled, environmentFeatures[i].CreationDate)
  }

  return environmentFeatures, err
}

func (s *store) CreateEnvironmentFeature(featureID, environmentID *int64) error {
  environmentFeature := model.NewEnvironmentFeature(featureID, environmentID)
  return meddler.Insert(s.db, "environment_feature", environmentFeature)
}

func (s *store) ChangeEnvironmentFeatureState(environmentFeature *model.EnvironmentFeature) error {
  return meddler.Update(s.db, "environment_feature", environmentFeature)
}
