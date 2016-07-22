pkgs := . ./api/... ./client/... ./notifier/... ./store/...

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
	@$(dc) run laika npm install
	@$(dc) run laika bower --allow-root install
.PHONY: vendor

schema:
	@echo "Generating schema..."
	@go-bindata -pkg schema -o store/schema/schema.go -ignore \.go store/schema/...
.PHONY: schema

lint:
	@echo "Linting sourcecode..."
	@$(dc) run laika go vet $(pkgs)
	@$(dc) run laika ./node_modules/gulp-cli/bin/gulp.js eslint
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

clean:
	@echo "Cleaning environment..."
	@$(dc) stop
	@$(dc) rm -fva
.PHONY: clean
