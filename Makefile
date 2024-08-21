start-test-server : setup-test-env
	path="." name="test" ext="env" go run main.go 
start-dev-server :
	path="." name="dev" ext="env" go run main.go
test: setup-test-env test-test-db clean-test-env

test-dev: setup-dev-env test-dev-db clean-dev-env

setup-test-env:
	cd ./database && docker compose -f docker-compose-test.yml up -d

setup-dev-env:
	cd ./database && docker compose -f docker-compose.yml up -d

clean-test-env:
	cd ./database && docker compose -f docker-compose-test.yml down

clean-dev-env:
	cd ./database && docker compose -f docker-compose.yml down

test-test-db:
	cd ./database && path="../../" name="test" ext="env" go test -v -cover -count=1 ./...

test-dev-db:
	cd ./database && path="../../" name="dev" ext="env" go test -v -cover -count=1 ./...

test-api :
	go test -v -cover -count=1 ./api
gen-mock-db:	
	mockgen -package mock_db -destination database/mock/store.go  github.com/REZ-OAN/simplebank/database/sqlc Store


gen-query:
	cd ./database && sqlc generate


.PHONY: start-server test test-dev setup-test-env setup-dev-env clean-test-env clean-dev-env test-test-db test-dev-db gen-query