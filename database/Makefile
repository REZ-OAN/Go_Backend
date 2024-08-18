test: setup-test-env test-config clean-test-env

test-dev: setup-dev-env dev-config clean-dev-env

setup-test-env:
	docker compose -f docker-compose-test.yml up -d

setup-dev-env:
	docker compose -f docker-compose.yml up -d 

clean-test-env:
	docker compose -f docker-compose-test.yml down

clean-dev-env:
	docker compose -f docker-compose.yml down

test-config:
	path="../../" name="test" ext="env" go test -v -cover -count=1 ./...

dev-config:
	path="../../" name="dev" ext="env" go test -v -cover -count=1 ./...


gen-query :
	sqlc generate


.PHONY: test setup-test-env setup-dev-env clean-test-env clean-dev-env test-config dev-config gen-query