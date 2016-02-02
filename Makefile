build:
	@echo "===>  Building project..."
	@docker-compose build
.PHONY: build

init:
	@echo "===>  Creating database..."
	@docker-compose run feature-flag mysql -h mysql -u root -proot feature-flag-db < data/feature-flagdbschema.sql
.PHONY: init

test:
	@echo "===> Running tests..."
	@docker-compose run feature-flag go test ./...
.PHONY: test

run:
	@echo "===> Running services..."
	@docker-compose up feature-flag
.PHONY: run

clean:
	@echo "===>  Cleaning environment..."
	@docker-compose stop
	@docker-compose rm -f -v
.PHONY: clean
