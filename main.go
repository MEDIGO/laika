package main

import (
	"github.com/DataDog/datadog-go/statsd"
	"github.com/MEDIGO/feature-flag/api"
	"github.com/MEDIGO/feature-flag/store"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	graceful "gopkg.in/tylerb/graceful.v1"
	"os"
	"time"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	app := cli.NewApp()
	app.Name = "Feature-Flag"
	app.Usage = "MEDIGO Feature-Flag Service"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "port",
			Value:  "8000",
			Usage:  "Service port",
			EnvVar: "FEATURE_FLAG_PORT",
		},
		cli.IntFlag{
			Name:   "timeout",
			Value:  10,
			Usage:  "Shutdown timeout",
			EnvVar: "FEATURE_FLAG_TIMEOUT",
		},
		cli.StringFlag{
			Name:   "mysql-host",
			Value:  "mysql",
			Usage:  "MySQL host",
			EnvVar: "MYSQL_HOST",
		},
		cli.StringFlag{
			Name:   "mysql-port",
			Value:  "3306",
			Usage:  "MySQL port",
			EnvVar: "MYSQL_PORT",
		},
		cli.StringFlag{
			Name:   "mysql-username",
			Value:  "root",
			Usage:  "MySQL username",
			EnvVar: "MYSQL_USERNAME",
		},
		cli.StringFlag{
			Name:   "mysql-password",
			Value:  "root",
			Usage:  "MySQL password",
			EnvVar: "MYSQL_PASSWORD",
		},
		cli.StringFlag{
			Name:   "mysql-dbname",
			Value:  "feature-flag-db",
			Usage:  "MySQL dbname",
			EnvVar: "MYSQL_DBNAME",
		},
		cli.StringFlag{
			Name:   "statsd-host",
			Value:  "localhost",
			Usage:  "Statsd host",
			EnvVar: "STATSD_HOST",
		},
		cli.StringFlag{
			Name:   "statsd-port",
			Value:  "8125",
			Usage:  "Statsd port",
			EnvVar: "STATSD_PORT",
		},
	}

	app.Action = func(c *cli.Context) {

		store, err := store.NewStore(
			c.String("mysql-host"),
			c.String("mysql-port"),
			c.String("mysql-username"),
			c.String("mysql-password"),
			c.String("mysql-dbname"),
		)

		if err != nil {
			log.Fatal("failed to create Store: ", err)
		}

		stats, err := statsd.New(c.String("statsd-host") + ":" + c.String("statsd-port"))
		if err != nil {
			log.Fatal("failed to create Statsd client: ", err)
		}

		server := api.NewServer(store, stats)

		log.Info("Starting server on port ", c.String("port"))
		graceful.Run(":"+c.String("port"), time.Duration(c.Int("timeout"))*time.Second, server)
	}

	app.Run(os.Args)

}
