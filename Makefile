DOCKER_IMAGE_MIGRATION_NAME=db_migration

build_migration:
	@echo "start to build docker image - $(DOCKER_IMAGE_MIGRATION_NAME)"
	docker build -t $(DOCKER_IMAGE_MIGRATION_NAME) -f docker/migration/Dockerfile .

migrate_script:
	docker run --rm --network indexer_default --env-file docker/migration/.env $(DOCKER_IMAGE_MIGRATION_NAME)

migrate_up:
	docker run --rm -v "$(shell pwd)/db/migrations:/migrations" --network indexer_default migrate/migrate -path=/migrations/ -database "mysql://mysql:mysql@tcp(db:3306)/test" up

up:
	docker-compose up -d
