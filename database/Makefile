
## Arguments
PASSWORD=secret
USER=root
HOST_PORT=5432
HOST=localhost
NAME=simple-bank
DB_NAME=simple_bank
IMAGE=simple-bank-db
DB_DATA_VOLUME_MOUNT=/home/baymax/Documents/SimpleBank/
MIGRATE_PATH=/home/baymax/Documents/SimpleBank/backend/database/migration
GIT_MIGRATE_PATH=./migration
DATABASE_CONNECTION_URI=postgresql://$(USER):$(PASSWORD)@$(HOST):$(HOST_PORT)/$(DB_NAME)?sslmode=disable
build-db-img:
	docker build -t $(IMAGE) .
launch-db-con:
	docker run --name $(NAME) -p $(HOST_PORT):5432 -e POSTGRES_PASSWORD=$(PASSWORD) -e POSTGRES_USER=$(USER) -v $(DB_DATA_VOLUME_MOUNT):/var/lib/postgresql/data -d $(IMAGE) 
launch-db-con-with-db-create:
	docker run --name $(NAME) -p $(HOST_PORT):5432 -e POSTGRES_PASSWORD=$(PASSWORD) -e POSTGRES_USER=$(USER) -e POSTGRES_DB=$(DB_NAME) -d $(IMAGE) 
create-db:
	docker exec -it $(NAME) createdb --username=$(USER) --owner=$(USER) $(DB_NAME)

drop-db:
	docker exec -it $(NAME) psql -U $(USER) -c "SELECT pid, pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = '$(DB_NAME)' AND pid <> pg_backend_pid();"
	docker exec -it $(NAME) dropdb $(DB_NAME)
delete-all-data-db:
	docker exec -it $(NAME) psql -U $(USER) -d $(DB_NAME) -c "TRUNCATE accounts CASCADE;"

migrate-up-db:
	migrate -path $(MIGRATE_PATH) -database $(DATABASE_CONNECTION_URI) -verbose up

migrate-up-db-git:
	migrate -path $(GIT_MIGRATE_PATH) -database $(DATABASE_CONNECTION_URI) -verbose up
	
migrate-down-db:
	migrate -path $(MIGRATE_PATH) -database $(DATABASE_CONNECTION_URI) -verbose down

gen-query :
	sqlc generate
test-db :
	go test -v -cover -count=1 ./...

.PHONY: build-db-img launch-db-con create-db drop-db migrate-up-db migrate-down-db gen-query test-db