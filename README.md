# Laika

[![CircleCI](https://circleci.com/gh/MEDIGO/laika.svg?style=shield)](https://circleci.com/gh/MEDIGO/laika) [![Coverage Status](https://coveralls.io/repos/github/MEDIGO/laika/badge.svg)](https://coveralls.io/github/MEDIGO/laika)

Laika is a feature flag/feature toggle service, written in Go, that allows the creation of flags and their activation/deactivation for specific environments. This way it is possible to control in which environments each feature is available. For instance, when a new feature is developed and released, it would make sense if it was only made available, at first, in a testing or Q&A environment, and only later in production. With Laika this can be achieved by simply going to a web page, selecting the feature, and changing its status on the desired environments.

Using Laika in a project thus allows for fast and continuous feature release and deployment.

`laika-php` is a PHP library that handles communication with Laika's API. You can get it [here](https://github.com/MEDIGO/laika-php).

## API Reference

Laika uses CQRS, the query endpoints are as follows.

| Method  | Endpoint                  | Description                |
| ------- | ------------------------- | -------------------------- |
| `GET`   | `/api/health`             | Check the service health   |
| `GET`   | `/api/features`           | List all features          |
| `GET`   | `/api/features/:name`     | Get a feature by name      |
| `GET`   | `/api/environments`       | List all environments      |

The command endpoint is for manipulating data.

| Method  | Endpoint                          | Example body                                                   | Description               |
| ------- | --------------------------------- | -------------------------------------------------------------- | ------------------------- |
| `POST`  | `/api/events/environment_created` | `{"name":"staging"}`                                           | Create a new environment. |
| `POST`  | `/api/events/feature_created`     | `{"name":"feature1"}`                                          | Create a new feature.     |
| `POST`  | `/api/events/user`                | `{"username":"admin","password":"secret"}`                     | Create a new user.        |
| `POST`  | `/api/events/feature_toggled`     | `{"feature":"feature1","environment":"staging","status":true}` | Toggle a feature.         |
| `POST`  | `/api/events/feature_deleted`     | `{"name":"feature1"}`                                          | Delete a feature.         |
| `POST`  | `/api/events/environments_ordered`| `{"order":["dev","staging"]}`                                  | Change env display order. |

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

To develop Laika you need to have the following tools installed in your machine:

* [Go](https://golang.org/doc/install)
* [Node.JS](https://nodejs.org/en/download/)
* [Glide](https://github.com/Masterminds/glide)
* [MySQL](https://dev.mysql.com/downloads/installer/)

Then install all the Go and Javascript dependencies with:

```sh
$ make install
```


Build continuously the server and UI with:

```sh
$ make develop
```

And start hacking!

```sh
$ open http://localhost:8000
```

## Testing

The whole test suite can be executed with:

```
$ make test
```

Some test require a MySQL instance, you can pass the configuration to them with the following
environment variables:

```
LAIKA_TEST_MYSQL_HOST=localhost
LAIKA_TEST_MYSQL_PORT=3306
LAIKA_TEST_MYSQL_USERNAME=root
LAIKA_TEST_MYSQL_PASSWORD=root
LAIKA_TEST_MYSQL_DBNAME=test
```

## Current state of the project

In the current release of Laika, it is possible to create feature flags and enable/disable them in the existing environments.

### Wishlist

- Specify country access (e.g. feature only enabled in Germany).
- Specify user access with percentage (e.g. feature only enabled for 30% of the user base).
- Have a field for environment creation on the web page.
- History for flag status changes.

## Copyright and license

Copyright © 2016 MEDIGO GmbH.

Laika is licensed under the MIT License. See [LICENSE](LICENSE) for the full license text.
