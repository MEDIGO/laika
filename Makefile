all: install build report publish clean
.PHONY: all

build:
	@echo "Building source code..."
	@scripts/build.sh
.PHONY: build

install:
	@echo "Installing dependencies..."
	@scripts/install.sh
.PHONY: install

generate:
	@echo "Generating source code..."
	@scripts/generate.sh
.PHONY: generate

lint:
	@echo "Linting sourcecode..."
	@scripts/lint.sh
.PHONY: lint

test:
	@echo "Running tests..."
	@scripts/test.sh
.PHONY: test

develop:
	@echo "Running server..."
	@scripts/develop.sh
.PHONY: develop

report:
	@echo "Reporting coverage..."
	@scripts/report.sh
.PHONY: report

image:
	@echo "Building Docker image..."
	@scripts/image.sh
.PHONY: image

publish:
	@echo "Publishing docker image..."
	@scripts/publish.sh
.PHONY: publish

clean:
	@echo "Cleaning environment..."
	@rm -rf bin public
.PHONY: clean
