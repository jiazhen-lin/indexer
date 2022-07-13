migrate_up:
	docker run --rm -v "$(shell pwd)/db/migrations:/migrations" --network indexer_default migrate/migrate -path=/migrations/ -database "mysql://mysql:mysql@tcp(db:3306)/test" up

up:
	docker-compose up -d
