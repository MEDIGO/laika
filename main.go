package main

import (
	"os"
	"time"
	
	"github.com/DataDog/datadog-go/statsd"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	graceful "gopkg.in/tylerb/graceful.v1"

	"github.com/MEDIGO/laika/api"
	"github.com/MEDIGO/laika/notifier"
	"github.com/MEDIGO/laika/store"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}
//TODO set Prefix by echo Router library?
func main() {
	app := cli.NewApp()
	app.Author = "MEDIGO GmbH"
	os.Setenv("PREFIX_ROUTE", "/laika") //Added a prefix as an environment variable
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "port",
			Value:  "8000",
			Usage:  "Service port",
			EnvVar: "LAIKA_PORT",
		},
		cli.IntFlag{
			Name:   "timeout",
			Value:  10,
			Usage:  "Shutdown timeout",
			EnvVar: "LAIKA_TIMEOUT",
		},
		cli.StringFlag{
			Name:   "mysql-host",
			Value:  "localhost",
			Usage:  "MySQL host",
			EnvVar: "LAIKA_MYSQL_HOST",
		},
		cli.StringFlag{
			Name:   "mysql-port",
			Value:  "3306",
			Usage:  "MySQL port",
			EnvVar: "LAIKA_MYSQL_PORT",
		},
		cli.StringFlag{
			Name:   "mysql-username",
			Value:  "root",
			Usage:  "MySQL username",
			EnvVar: "LAIKA_MYSQL_USERNAME",
		},
		cli.StringFlag{
			Name:   "mysql-password",
			Value:  "root",
			Usage:  "MySQL password",
			EnvVar: "LAIKA_MYSQL_PASSWORD",
		},
		cli.StringFlag{
			Name:   "mysql-dbname",
			Value:  "laika",
			Usage:  "MySQL dbname",
			EnvVar: "LAIKA_MYSQL_DBNAME",
		},
		cli.StringFlag{
			Name:   "statsd-host",
			Value:  "localhost",
			Usage:  "Statsd host",
			EnvVar: "LAIKA_STATSD_HOST",
		},
		cli.StringFlag{
			Name:   "statsd-port",
			Value:  "8125",
			Usage:  "Statsd port",
			EnvVar: "LAIKA_STATSD_PORT",
		},
		cli.StringFlag{
			Name:   "root-username",
			Usage:  "Root username",
			Value:  "root",
			EnvVar: "LAIKA_ROOT_USERNAME",
		},
		cli.StringFlag{
			Name:   "root-password",
			Usage:  "Root password",
			Value:  "root",
			EnvVar: "LAIKA_ROOT_PASSWORD",
		},
		cli.StringFlag{
			Name:   "slack-webhook-url",
			Usage:  "Slack webhook URL",
			EnvVar: "LAIKA_SLACK_WEBHOOK_URL",
		},
		cli.StringFlag{
			Name: "route-prefix",
			Usage: "Route prefix",
			EnvVar: "ROUTE_PREFIX",
			Value: "/laika",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "Runs laika's feature flag service",
			Action: func(c *cli.Context) {
				store, err := store.NewMySQLStore(
					c.GlobalString("mysql-username"),
					c.GlobalString("mysql-password"),
					c.GlobalString("mysql-host"),
					c.GlobalString("mysql-port"),
					c.GlobalString("mysql-dbname"),
				)

				if err != nil {
					log.Fatal("Failed to create Store: ", err)
				}

				stats, err := statsd.New(c.GlobalString("statsd-host") + ":" + c.GlobalString("statsd-port"))
				if err != nil {
					log.Fatal("Failed to create Statsd client: ", err)
				}

				notifier := notifier.NewSlackNotifier(c.GlobalString("slack-webhook-url"))

				server, err := api.NewServer(api.ServerConfig{
					RootUsername: c.GlobalString("root-username"),
					RootPassword: c.GlobalString("root-password"),
					Store:        store,
					Stats:        stats,
					Notifier:     notifier,
					PrefixRoute:  c.GlobalString("route-prefix"),
				})
				if err != nil {
					log.Fatal("Failed to create server: ", err)

				}

				log.Info("Starting server on port ", c.GlobalString("port"))
				graceful.Run(":"+c.GlobalString("port"), time.Duration(c.Int("timeout"))*time.Second, server)
			},
		},
		{
			Name:  "migrate",
			Usage: "Migrates the store schema to the latest available version",
			Action: func(c *cli.Context) error {
				store, err := store.NewMySQLStore(
					c.GlobalString("mysql-username"),
					c.GlobalString("mysql-password"),
					c.GlobalString("mysql-host"),
					c.GlobalString("mysql-port"),
					c.GlobalString("mysql-dbname"),
				)

				if err != nil {
					log.Fatal("Failed to create Store: ", err)
				}

				if err := store.Ping(); err != nil {
					log.Fatal("Failed to connect with store: ", err)
				}

				if err := store.Migrate(); err != nil {
					log.Fatal("Failed to migrate store schema: ", err)
				}

				return nil

			},
		},
	}
	app.Run(os.Args)
}
