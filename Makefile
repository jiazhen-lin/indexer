DOCKER_IMAGE_API_NAME=block_api
DOCKER_IMAGE_INDEXER_NAME=indexer
DOCKER_IMAGE_MIGRATION_NAME=db_migration

build_api:
	@echo "start to build docker image - $(DOCKER_IMAGE_API_NAME)"
	docker build -t $(DOCKER_IMAGE_API_NAME) -f docker/api/Dockerfile .

build_indexer:
	@echo "start to build docker image - $(DOCKER_IMAGE_INDEXER_NAME)"
	docker build -t $(DOCKER_IMAGE_INDEXER_NAME) -f docker/indexer/Dockerfile .

build_migration:
	@echo "start to build docker image - $(DOCKER_IMAGE_MIGRATION_NAME)"
	docker build -t $(DOCKER_IMAGE_MIGRATION_NAME) -f docker/migration/Dockerfile .

migrate_script:
	docker run --rm --network indexer_default --env-file docker/migration/.env $(DOCKER_IMAGE_MIGRATION_NAME)

migrate_up:
	docker run --rm -v "$(shell pwd)/db/migrations:/migrations" --network indexer_default migrate/migrate -path=/migrations/ -database "mysql://mysql:mysql@tcp(db:3306)/test" up

deploy_api:
	docker run --name block_api --rm --network indexer_default -p "8081:8081" $(DOCKER_IMAGE_API_NAME) --restart=always

deploy_indexer:
	docker run --name indexer --rm --network indexer_default --env-file docker/config/.env -d $(DOCKER_IMAGE_INDEXER_NAME) --restart=always

up:
	docker-compose up -d
