package main

import (
  "fmt"
  log "github.com/Sirupsen/logrus"
  _ "github.com/go-sql-driver/mysql"
  "github.com/MEDIGO/feature-flag/store"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

  store, err := store.NewStore(
    "root",
    "root",
    "192.168.99.100",
    "3306",
    "featflagdb",
  )

  if err != nil {
    fmt.Printf("Failed to create Store: ", err) //TODO REPLACE WITH LOG
  }

  store.ListEnvironments(nil, nil, nil, nil)

}
