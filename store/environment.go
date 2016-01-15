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

func (s *store) GetEnvironmentByID(id int64) (*model.Environment, error) {
  environment := new(model.Environment)
  err := meddler.Load(s.db, "environment", environment, id)

  fmt.Printf("%d %s %t %s\n", environment.Id, *environment.EnvironmentKey, *environment.Enabled, environment.CreationDate) //TODO REMOVE

  return environment, err
}

func (s *store) ListEnvironments(environmentKey *string, enabled *bool , from *time.Time, to *time.Time) ([]*model.Environment ,error){
  query := sq.Select("*").From("environment")

  if environmentKey != nil {
		query = query.Where(sq.Eq{"environment_key": environmentKey})
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

	environments := []*model.Environment{}
	err = meddler.QueryAll(s.db, &environments, sql, args...)

  fmt.Printf("%d\n",len(environments))

  for i:=0; i<len(environments); i++ {
    fmt.Printf("%d %s %t %s\n", environments[i].Id, *environments[i].EnvironmentKey, *environments[i].Enabled, environments[i].CreationDate)
  }

  return environments, err
}

func (s *store) CreateEnvironment(environmentKey *string) error {
  environment := model.NewEnvironment(environmentKey)
  return meddler.Insert(s.db, "environment", environment)
}

func (s *store) ChangeEnvironmentState(environment *model.Environment) error {
  return meddler.Update(s.db, "environment", environment)
}
