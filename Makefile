update-util:
	docker-compose run --rm --entrypoint "go get github.com/vality-prototype/vality-utility-service" app

start:
	docker-compose run --service-ports --rm --entrypoint "sh ./scripts/dev.sh" app

clean:
	docker-compose down --rmi all --volume

migrate:
	docker-compose run --rm --entrypoint "sh ./scripts/migrate.sh up" app

migrate-test:
	docker-compose exec mysql /etc/mysql/test/create_db_test.sh
	docker-compose run --rm --entrypoint "sh ./scripts/migrate_test.sh up" app

tidy:
	docker-compose run --rm --entrypoint "go mod tidy" app

run-test:
	docker-compose run --rm --entrypoint "sh ./scripts/test.sh" app-ut
	# docker-compose run --rm --entrypoint "go test ./test/... -v" app-ut
	# docker-compose run --rm --entrypoint "go test ./app/... ./cmd/... ./test/... -v" app-ut