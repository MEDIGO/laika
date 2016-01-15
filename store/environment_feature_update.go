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

func (s *store) GetEnvironmentFeatureUpdateByID(id int64) (*model.EnvironmentFeatureUpdate, error) {
  environmentFeatureUpdate := new(model.EnvironmentFeatureUpdate)
  err := meddler.Load(s.db, "environment_feature_update", environmentFeatureUpdate, id)

  fmt.Printf("%d %d %d %t %s\n", environmentFeatureUpdate.Id, *environmentFeatureUpdate.FeatureID, *environmentFeatureUpdate.EnvironmentID, *environmentFeatureUpdate.Enabled, environmentFeatureUpdate.Date)

  return environmentFeatureUpdate, err
}

func (s *store) ListEnvironmentFeatureUpdates(featureID *int64, environmentID *int64, enabled *bool , from *time.Time, to *time.Time) ([]*model.EnvironmentFeatureUpdate ,error){
  query := sq.Select("*").From("environment_feature_update")

  if featureID != nil {
		query = query.Where(sq.Eq{"feature_id": featureID})
	}

  if environmentID != nil {
		query = query.Where(sq.Eq{"environment_id": environmentID})
	}

  switch {
    case from != nil && to != nil:
      fmt.Printf("ONE\n")
      query = query.Where(sq.Gt{"date": from})
      query = query.Where(sq.Lt{"date": to})
      break
    case from != nil:
      fmt.Printf("TWO\n")
      query = query.Where(sq.Gt{"date": from})
      break
    case to != nil:
      fmt.Printf("THREE\n")
      query = query.Where(sq.Lt{"date": to})
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

	environmentFeatureUpdates := []*model.EnvironmentFeatureUpdate{}
	err = meddler.QueryAll(s.db, &environmentFeatureUpdates, sql, args...)

  fmt.Printf("%d\n",len(environmentFeatureUpdates))

  for i:=0; i<len(environmentFeatureUpdates); i++ {
    fmt.Printf("%d %d %d %t %s\n", environmentFeatureUpdates[i].Id, *environmentFeatureUpdates[i].FeatureID, *environmentFeatureUpdates[i].EnvironmentID, *environmentFeatureUpdates[i].Enabled, environmentFeatureUpdates[i].Date)
  }

  return environmentFeatureUpdates, err
}

func (s *store) CreateEnvironmentFeatureUpdate(featureID, environmentID *int64, enabled *bool) error {
  environmentFeatureUpdate := model.NewEnvironmentFeatureUpdate(featureID, environmentID, enabled)
  return meddler.Insert(s.db, "environment_feature_update", environmentFeatureUpdate)
}
