PASSWORD=secret
USER=wizard
TEST_PORT=5434
DEV_PORT=5432
HOST=localhost
TEST_NAME=simple_bank_test
DEV_NAME=simple_bank_dev
BACKEND=simple_bank_backend
DB_NAME=simple_bank
IMAGE=postgres
DB_IP = 10.10.10.5
BACKEND_IP = 10.10.10.3
DB_DATA_VOLUME_MOUNT_TEST=~/Documents/SimpleBank/db_data/test_db
DB_DATA_VOLUME_MOUNT_DEV=~/Documents/SimpleBank/db_data/dev_db
MIGRATE_PATH=./database/migration
DATABASE_CONNECTION_URI_TEST=postgresql://$(USER):$(PASSWORD)@$(HOST):$(TEST_PORT)/$(DB_NAME)?sslmode=disable
DATABASE_CONNECTION_URI_DEV=postgresql://$(USER):$(PASSWORD)@$(HOST):$(DEV_PORT)/$(DB_NAME)?sslmode=disable
VERSION = 1
NETWORK_NAME = simple-bank-net
NETWORK_SUBNET = 10.10.10.0/24

start-test-server : setup-test-env
	path="." name="test" ext="env" go run main.go 

start-dev-server : create-docker-network setup-dev-env
	docker run --name $(BACKEND) --network $(NETWORK_NAME) --ip $(BACKEND_IP) -p 8080:8080 -e GIN_MODE=release -d simple_bank:latest 

build-dev-image :
	docker build --no-cache -t simple_bank:latest .

test: setup-test-env wait-for-postgres-test migrate-up-test test-test-db clean-test-env


test-dev: create-docker-network setup-dev-env wait-for-postgres-dev migrate-up-dev test-dev-db clean-dev-env

setup-test-env:
	docker run --name $(TEST_NAME) -p $(TEST_PORT):5432 -e POSTGRES_PASSWORD=$(PASSWORD) -e POSTGRES_USER=$(USER) -e POSTGRES_DB=$(DB_NAME) -v $(DB_DATA_VOLUME_MOUNT_TEST):/var/lib/postgresql/data -d $(IMAGE)

create-docker-network:
	docker network create --driver bridge --subnet $(NETWORK_SUBNET) $(NETWORK_NAME)

setup-dev-env:
	docker run --name $(DEV_NAME) --network $(NETWORK_NAME) --ip $(DB_IP) -p $(DEV_PORT):5432 -e POSTGRES_PASSWORD=$(PASSWORD) -e POSTGRES_USER=$(USER) -e POSTGRES_DB=$(DB_NAME) -v $(DB_DATA_VOLUME_MOUNT_DEV):/var/lib/postgresql/data -d $(IMAGE)

clean-test-env:
	docker rm -f -v $(TEST_NAME)

clean-dev-env:
	docker rm -f -v $(DEV_NAME)
	docker rm -f -v $(BACKEND)
	docker network rm -f $(NETWORK_NAME)

test-test-db:
	cd ./database && path="../../" name="test" ext="env" go test -v -cover -count=1 ./...

test-dev-db:
	cd ./database && path="../../" name="dev" ext="env" go test -v -cover -count=1 ./...

wait-for-postgres-test:
	@echo "Waiting for PostgreSQL container to be ready..."
	until docker exec $(TEST_NAME) pg_isready -h localhost -p 5432 > /dev/null 2>&1; do \
		echo "Waiting for PostgreSQL to be ready..."; \
		sleep 1; \
	done
	@echo "PostgreSQL is ready!"

wait-for-postgres-dev:
	@echo "Waiting for PostgreSQL container to be ready..."
	until docker exec $(DEV_NAME) pg_isready -h localhost -p 5432 > /dev/null 2>&1; do \
		echo "Waiting for PostgreSQL to be ready..."; \
		sleep 1; \
	done
	@echo "PostgreSQL is ready!"

test-api :
	go test -v -cover -count=1 ./api

create-new-migration:
	migrate create -ext sql -dir ./database/migration -seq add_users

migrate-up-test :
	migrate -path $(MIGRATE_PATH) -database $(DATABASE_CONNECTION_URI_TEST) -verbose up
migrate-up-dev :
	migrate -path $(MIGRATE_PATH) -database $(DATABASE_CONNECTION_URI_DEV) -verbose up
migrate-up-test-by-one :
	migrate -path $(MIGRATE_PATH) -database $(DATABASE_CONNECTION_URI_TEST) -verbose up $(VERSION)

migrate-up-dev-by-one :
	migrate -path $(MIGRATE_PATH) -database $(DATABASE_CONNECTION_URI_DEV) -verbose up $(VERSION)

migrate-down-test-by-one :
	migrate -path $(MIGRATE_PATH) -database $(DATABASE_CONNECTION_URI_TEST) -verbose down $(VERSION)

migrate-down-dev-by-one :
	migrate -path $(MIGRATE_PATH) -database $(DATABASE_CONNECTION_URI_DEV) -verbose down $(VERSION)

migrate-down-test :
	migrate -path $(MIGRATE_PATH) -database $(DATABASE_CONNECTION_URI_TEST) -verbose down

migrate-down-dev :
	migrate -path $(MIGRATE_PATH) -database $(DATABASE_CONNECTION_URI_DEV) -verbose down 
gen-mock-db:	
	mockgen -package mock_db -destination ./database/mock/store.go  github.com/REZ-OAN/simplebank/database/sqlc Store

gen-query:
	cd ./database && sqlc generate


.PHONY: start-server test test-dev setup-test-env setup-dev-env clean-test-env clean-dev-env test-test-db test-dev-db gen-query