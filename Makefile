create-db:
	@echo "Creating database..."
	@docker-compose run feat-flag mysql -h mysql -u root -proot featflagdb < data/featflagdbschema.sql
.PHONY: create-db

compose-db:
	@echo "Composing database..."
	@docker-compose up mysql
.PHONY: compose-db