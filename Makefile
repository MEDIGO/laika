dc = docker-compose
ifeq ($(CI), true)
	dc = docker-compose -f docker-compose-ci.yml
endif

all: build install lint migrate test
.PHONY: all

build:
	@echo "Building project..."
	@$(dc) build
.PHONY: build

install:
	@echo "Installing dependencies..."
	@$(dc) run laika scripts/install.sh
.PHONY: install

generate:
	@echo "Generating source code..."
	@$(dc) run laika scripts/generate.sh
.PHONY: generate

lint:
	@echo "Linting sourcecode..."
	@$(dc) run laika scripts/lint.sh
.PHONY: lint

test:
	@echo "Running tests..."
	@$(dc) run laika scripts/test.sh
.PHONY: test

up:
	@echo "Running services..."
	@$(dc) up laika
.PHONY: run

migrate:
	@echo "Migrating DB..."
	@$(dc) run laika scripts/migrate.sh
.PHONY: migrate

shell:
	@echo "Opening shell..."
	@$(dc) run laika sh
.PHONY: shell

report:
	@echo "Reporting coverage..."
	@$(dc) run laika scripts/report.sh
.PHONY: report

publish:
	@echo "Publishing docker image..."
	scripts/publish.sh
.PHONY: publish

deploy:
	@echo "Deploying docker image..."
	scripts/deploy.sh
.PHONY: deploy

clean:
	@echo "Cleaning environment..."
	@$(dc) stop
	@$(dc) rm -fva
.PHONY: clean
