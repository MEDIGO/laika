all: build vendor init validate publish
.PHONY: all

build:
	@echo "===> Building project..."
	@docker-compose build
.PHONY: build

vendor:
	@echo "===> Installing dependencies..."
	@docker-compose run laika npm install
	@docker-compose run laika bower --allow-root install
.PHONY: vendor

init:
	@echo "===> Initialising database..."
	@docker-compose run laika mysql -h mysql -u root -proot laika < schema/1.sql
	@docker-compose run laika mysql -h mysql -u root -proot laika < schema/2.sql
.PHONY: init

validate: lint test
.PHONY: validate

lint:
	@echo "===> Linting sourcecode..."
	@docker-compose run laika ./node_modules/gulp-cli/bin/gulp.js eslint
.PHONY: lint

test:
	@echo "===> Running tests..."
	@docker-compose run laika go test . ./client ./test/integration
.PHONY: test

run:
	@echo "===> Running services..."
	@docker-compose up laika
.PHONY: run

shell:
	@echo "===> Opening shell..."
	@docker-compose run laika sh
.PHONY: shell

publish:
	@echo "===> Publishing docker image..."
	@docker build -t quay.io/medigo/laika .
	@docker tag -f quay.io/medigo/laika:latest quay.io/medigo/laika:$(shell git rev-parse HEAD)
	@docker push quay.io/medigo/laika
	@docker push quay.io/medigo/laika:$(shell git rev-parse HEAD)
.PHONY: publish

deploy:
	@echo "Deploying docker image..."
	@docker pull quay.io/medigo/laika:$(shell git rev-parse HEAD)
	@aws ecs register-task-definition --family $(ECS_FAMILY) --container-definitions '$(shell ./ecs-container-definitions.sh)'
	@aws ecs update-service --service $(ECS_FAMILY) --task-definition $(ECS_FAMILY)
.PHONY: deploy

clean:
	@echo "===> Cleaning environment..."
	@docker-compose stop
	@docker-compose rm -f -v
.PHONY: clean
