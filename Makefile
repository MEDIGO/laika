PACKAGES = ./api/... ./client/... ./notifier/... ./store/... ./test/... ./util/...

dc = docker-compose
ifeq ($(CI), true)
	dc = docker-compose -f docker-compose-ci.yml
endif

all: build deps lint test publish

build:
	@echo "===> Building project..."
	@$(dc) build

deps:
	@echo "===> Installing dependencies..."
	@$(dc) run laika npm install
	@$(dc) run laika bower --allow-root install

schema:
	@echo "===> Generating schema..."
	@go-bindata -pkg schema -o store/schema/schema.go -ignore \.go store/schema/...

lint:
	@echo "===> Linting sourcecode..."
	@$(dc) run laika go vet $(PACKAGES)
	@$(dc) run laika ./node_modules/gulp-cli/bin/gulp.js eslint

test:
	@echo "===> Running tests..."
	@$(dc) run laika go test $(PACKAGES)

run:
	@echo "===> Running services..."
	@$(dc) up laika

shell:
	@echo "===> Opening shell..."
	@$(dc) run laika sh

publish:
	@echo "===> Publishing docker image..."
	@docker build -t medigo/laika:latest .
	@docker login -e $(DOCKER_EMAIL) -u $(DOCKER_USER) -p $(DOCKER_PASS)
	@docker push medigo/laika:latest

deploy:
	@echo "Deploying docker image..."
	@docker pull quay.io/medigo/laika:$(shell git rev-parse HEAD)
	@aws ecs register-task-definition --family $(ECS_FAMILY) --container-definitions '$(shell ./ecs-container-definitions.sh)'
	@aws ecs update-service --service $(ECS_FAMILY) --task-definition $(ECS_FAMILY)

clean:
	@echo "===> Cleaning environment..."
	@$(dc) stop
	@$(dc) rm -f -v

.PHONY: all build deps schema lint test run shell publish deploy clean
