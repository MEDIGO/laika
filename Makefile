PACKAGES = ./api/... ./client/... ./notifier/... ./store/... ./test/... ./util/...

all: build deps init lint test publish

build:
	@echo "===> Building project..."
	@docker-compose build

deps:
	@echo "===> Installing dependencies..."
	@docker-compose run laika npm install
	@docker-compose run laika bower --allow-root install

init:
	@echo "===> Initialising database..."
	@docker-compose run laika mysql -h mysql -u root -proot laika < schema/1.sql
	@docker-compose run laika mysql -h mysql -u root -proot laika < schema/2.sql

lint:
	@echo "===> Linting sourcecode..."
	@docker-compose run laika go vet $(PACKAGES)
	@docker-compose run laika ./node_modules/gulp-cli/bin/gulp.js eslint

test:
	@echo "===> Running tests..."
	@docker-compose run laika go test $(PACKAGES)

run:
	@echo "===> Running services..."
	@docker-compose up laika

shell:
	@echo "===> Opening shell..."
	@docker-compose run laika sh

publish:
	@echo "===> Publishing docker image..."
	@docker build -t quay.io/medigo/laika .
	@docker tag -f quay.io/medigo/laika:latest quay.io/medigo/laika:$(shell git rev-parse HEAD)
	@docker push quay.io/medigo/laika
	@docker push quay.io/medigo/laika:$(shell git rev-parse HEAD)

deploy:
	@echo "Deploying docker image..."
	@docker pull quay.io/medigo/laika:$(shell git rev-parse HEAD)
	@aws ecs register-task-definition --family $(ECS_FAMILY) --container-definitions '$(shell ./ecs-container-definitions.sh)'
	@aws ecs update-service --service $(ECS_FAMILY) --task-definition $(ECS_FAMILY)

clean:
	@echo "===> Cleaning environment..."
	@docker-compose stop
	@docker-compose rm -f -v

.PHONY: all build deps init lint test run shell publish deploy clean
