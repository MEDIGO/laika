all: build vendor test
.PHONY: all

build:
	@echo "===>  Building project..."
	@docker-compose build
.PHONY: build

vendor:
	@echo "Installing dependencies..."
	@docker-compose run laika glide install
.PHONY: vendor

init:
	@echo "===>  Creating database..."
	@docker-compose run laika mysql -h mysql -u root -proot laika-db < schema/laikadbschema.sql
.PHONY: init

test:
	@echo "===> Running tests..."
	@docker-compose run laika go test . ./client ./test/integration
.PHONY: test

run:
	@echo "===> Running services..."
	@docker-compose up laika
.PHONY: run

clean:
	@echo "===>  Cleaning environment..."
	@docker-compose stop
	@docker-compose rm -f -v
.PHONY: clean
