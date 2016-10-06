# Laika

[![CircleCI](https://circleci.com/gh/MEDIGO/laika.svg?style=shield)](https://circleci.com/gh/MEDIGO/laika) [![Coverage Status](https://coveralls.io/repos/github/MEDIGO/laika/badge.svg)](https://coveralls.io/github/MEDIGO/laika)

Laika is a feature flag/feature toggle service, written in Go, that allows the creation of flags and their activation/deactivation for specific environments. This way it is possible to control in which environments each feature is available. For instance, when a new feature is developed and released, it would make sense if it was only made available, at first, in a testing or Q&A environment, and only later in production. With Laika this can be achieved by simply going to a web page, selecting the feature, and changing its status on the desired environments.

Using Laika in a project thus allows for fast and continuous feature release and deployment.

`laika-php` is a PHP library that handles communication with Laika's API. You can get it [here](https://github.com/MEDIGO/laika-php).

## API Reference

| Method  | Endpoint                  | Description                |
| ------- | ------------------------- | -------------------------- |
| `GET`   | `/api/health`             | Check the service health   |
| `GET`   | `/api/features`           | List all features          |
| `GET`   | `/api/features/:name`     | Get a feature by name      |
| `POST`  | `/api/features`           | Create a feature           |
| `PATCH` | `/api/features/:name`     | Update a feature           |
| `GET`   | `/api/environments`       | List all environments      |
| `GET`   | `/api/environments/:name` | Get an environment by name |
| `POST`  | `/api/environments`       | Create an environment      |
| `PATCH` | `/api/environments/:name` | Update an environment      |
| `GET`   | `/api/users/:username`    | Get a user by username     |
| `POST`  | `/api/users`              | Create a user              |

## Client

Laika contains a polling HTTP client that allows to easily check for enabled/disabled features on Go code. It can be found in the `client` package. While Laika uses the `vendor` directory to store external dependencies, `client` can be imported without any vendoring.

### Install

```
$ go get -u github.com/MEDIGO/laika/client
```

### Usage

```go
package main

import (
	"log"

	"github.com/MEDIGO/laika/client"
)

func main() {
	c, err := client.NewClient(client.Config{
		Addr:        "127.0.0.1:8000",
		Username:    "my-username",
		Password:    "my-password",
		Environment: "prod",
	})

	if err != nil {
		log.Fatal(err)
	}

	if c.IsEnabled("my-awesome-feature", false) {
		log.Print("IT'S ALIVE!")
	} else {
		log.Print("Move along. Nothing to see here.")
	}
}
```

## Developing

You will need [Docker Compose](https://docs.docker.com/compose/) to easily run Laika locally. The application can be boostraped with the following steps:

```sh
# setup the root username and password
$ echo "LAIKA_ROOT_USERNAME=my-username" > .env
$ echo "LAIKA_ROOT_PASSWORD=my-password" >> .env

# build the application
$ make build

# migrate the database
$ make migrate

# run the application
$ make run

# open the web UI
$ open http://localhost:8000
```

## Testing

Laika contains an integration tests suite that requires a available MySQL instance. When using Docker Compose, the whole test suite can be run with:

```
$ make test
```

## Current state of the project

In the current release of Laika, it is possible to create feature flags and enable/disable them in the existing environments.

### Wishlist

- Specify country access (e.g. feature only enabled in Germany).
- Specify user access with percentage (e.g. feature only enabled for 30% of the user base).
- Have a field for environment creation on the web page.
- History for flag status changes.
- New flags auto-registering when seen for the first time by laika.

## Copyright and license

Copyright © 2016 MEDIGO GmbH.

Laika is licensed under the MIT License. See [LICENSE](LICENSE) for the full license text.
