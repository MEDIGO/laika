all: build vendor validate publish
.PHONY: all

build:
	@echo "===> Building project..."
	@docker-compose build
.PHONY: build

vendor:
	@echo "===> Installing dependencies..."
	@docker-compose run laika glide install
.PHONY: vendor

init:
	@echo "===> Initialising database..."
	@docker-compose run laika mysql -h mysql -u root -proot laika-db < schema/1.sql
.PHONY: init

validate: lint test
.PHONY: validate

lint:
	@echo "===> Running gulp eslint..."
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

clean:
	@echo "===> Cleaning environment..."
	@docker-compose stop
	@docker-compose rm -f -v
.PHONY: clean
