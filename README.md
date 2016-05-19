# Laika

Laika is a feature flag/feature toggle service, written in Go, that allows the creation of flags and their activation/deactivation for specific environments. This way it is possible to control in which environments each feature is available. For instance, when a new feature is developed and released, it would make sense if it was only made available, at first, in a testing or Q&A environment, and only later in production. With Laika this can be achieved by simply going to a web page, selecting the feature, and changing its status on the desired environments.

Using Laika in a project thus allows for fast and continuous feature release and deployment.

`laika-php` is a PHP library that handles communication with Laika's API. You can get it here https://github.com/MEDIGO/laika-php.

## API

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

Laika contains a polling HTTP client that allows to easily check for enabled/disabled
features on Go code. It can be found in the `client` pacakge. While Laika uses
the `vendor` directory to store external dependencies, `client` can be imported
without any vendoring.


### Install

```
go get github.com/MEDIGO/laika/client
```

### Usage

```go
package main

import (
	"log"

	"github.com/MEDIGO/laika/client"
)

func main() {
	cfg := client.Config{
		Addr:        "127.0.0.1:8000",
		Username:    "my-username",
		Password:    "my-password",
		Environment: "prod",
	}

	c, err := client.NewClient(cfg)
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

## Setup

Setup docker and docker-compose in your machine (https://docs.docker.com/mac/ and https://docs.docker.com/compose/install/)

Clone the Laika project, open Docker Quickstart Terminal and go to the project directory.

Create a .env file with a username and password

```
LAIKA_ROOT_USERNAME=myusername
LAIKA_ROOT_PASSWORD=mypassword
```

Build the application by running `make build` inside its directory.

Finally, you can run the application using the `make run` command.

After Laika is running, open a browser and enter the IP of your docker machine (get it by running `docker-machine ip`).

Login by using root or user credentials.

Currently, when you create a new feature on the web page, you will not be able to enable/disable it yet as there are no environments. To create environments, you have to add them manually to the database. To do so you can use a software like Sequel Pro and insert the following on the corresponding fields:

```
Host: (the IP address you get by running the `docker-machine ip` command)
Username: root
Password: root
Port: (run the command `docker-compose ps`. Under Ports you will have something like `0.0.0.0:11111->2222/tcp`. In this case, your port would be `11111`)
```

Once connected to the database, create an entry to the `environment` table. It will now show on Laika's web page and you will be able to enable and disable features for the new environment.

## Tests

To run tests:

```
make test
```

## Current state of the project
In the current release of Laika, it is possible to create feature flags and enable/disable them in the existing environments.

## Wishlist

- Specify country access (e.g. feature only enabled in Germany)
- Specify user access with percentage (e.g. feature only enabled for 30% of the user base)
- Have a field for environment creation on the web page
- History for flag status changes
- New flags auto-registering when seen for the first time by laika

## How to contribute

- Fork the repository
- Clone it
- `git remote add upstream git@github.com:MEDIGO/laika.git`
- Code code code code code
- Make a pull request

### Thank you!
