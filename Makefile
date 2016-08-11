pkgs := . ./api/... ./client/... ./notifier/... ./store/...
commit := $(shell git rev-parse HEAD)

dc = docker-compose
ifeq ($(CI), true)
	dc = docker-compose -f docker-compose-ci.yml
endif

all: build vendor lint migrate test
.PHONY: all

build:
	@echo "Building project..."
	@$(dc) build
.PHONY: build

vendor:
	@echo "Installing dependencies..."
	@$(dc) run laika bower install --allow-root
	@$(dc) run laika glide install
.PHONY: vendor

schema:
	@echo "Generating schema..."
	@$(dc) run laika go-bindata -pkg schema -o store/schema/schema.go -ignore \.go store/schema/...
.PHONY: schema

lint:
	@echo "Linting sourcecode..."
	@$(dc) run laika go vet $(pkgs)
	@$(dc) run laika eslint .
.PHONY: lint

test:
	@echo "Running tests..."
	@$(dc) run laika go test $(pkgs)
.PHONY: test

up:
	@echo "Running services..."
	@$(dc) up laika
.PHONY: run

migrate:
	@echo "Migrating DB..."
	@$(dc) run laika go run main.go migrate
.PHONY: migrate

shell:
	@echo "Opening shell..."
	@$(dc) run laika sh
.PHONY: shell

publish:
	@echo "Publishing docker image..."
	@docker tag medigo/laika:latest medigo/laika:$(commit)
	@docker login -e $(DOCKER_EMAIL) -u $(DOCKER_USER) -p $(DOCKER_PASS)
	@docker push medigo/laika:latest
	@docker push medigo/laika:$(commit)
.PHONY: publish

deploy:
	@echo "Deploying docker image..."
	@docker pull medigo/laika:$(commit)
	@aws ecs register-task-definition --family $(ECS_FAMILY) --container-definitions '$(shell ./ecs-container-definitions.sh)'
	@aws ecs update-service --service $(ECS_FAMILY) --task-definition $(ECS_FAMILY)
	@aws ecs wait services-stable --services $(ECS_FAMILY)
.PHONY: deploy

clean:
	@echo "Cleaning environment..."
	@$(dc) stop
	@$(dc) rm -fva
.PHONY: clean
