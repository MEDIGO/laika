# laika

Laika is a feature flag/feature toggle service, written in Go, that allows the creation of flags and their activation/deactivation for specific environments. This way it is possible to control in which environments each feature is available. For instance, when a new feature is developed and released, it would make sense if it was only made available, at first, in a testing or Q&A environment, and only later in production. With Laika this can be achieved by simply going to a web page, selecting the feature, and changing its status on the desired environments.

Using Laika in a project thus allows for fast and continuous feature release and deployment.

laika-php is a PHP library that handles communication with Laika's API. You can get it here https://github.com/MEDIGO/laika-php.

##API

GET /api/health - Health check

GET /api/features - Get all features
GET /api/features/:name - Get feature by name
POST /api/features - Create feature
PATCH /api/features/:name - Update feature

GET /api/environments - Get all environments
GET /api/environments/:name - Get environment by name
POST /api/environments - Create environment
PATCH /api/environments/:name - Update environment

## Setup

Setup docker and docker-compose in your machine (https://docs.docker.com/mac/ and https://docs.docker.com/compose/install/)

Clone the Laika project, open Docker Quickstart Terminal and go to the project directory.

Create a .env file with a username and password

```
LAIKA_AUTH_USERNAME=myusername
LAIKA_AUTH_PASSWORD=mypassword
```

Build the application by running `make build` inside its directory and, afterwards, run `make init`.

Finally, you can run the application using the `make run` command.

After Laika is running, open a browser and enter the IP of your docker machine (get it by running `docker-machine ip`).

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

##Current state of the project
In the current release of Laika, it is possible to create feature flags and enable/disable them in the existing environments.

##Wishlist

- Specify country access (e.g. feature only enabled in Germany)
- Specify user access with percentage (e.g. feature only enabled for 30% of the user base)
- Have a field for environment creation on the web page
- SSL support
- History for flag status changes
- Slack notifications / webhooks
- New flags auto-registering when seen for the first time by laika

##How to contribute

- Fork the repository
- Clone it
- `git remote add upstream git@github.com:MEDIGO/laika.git`
- Code code code code code
- Make a pull request

###Thank you!
