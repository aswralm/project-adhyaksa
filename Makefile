
compose-up: ### Run docker-compose
	docker-compose up --build -d && docker-compose logs -f
.PHONY: compose-up

compose-up-integration-test: ### Run docker-compose with integration test
	docker-compose up --build --abort-on-container-exit --exit-code-from integration
.PHONY: compose-up-integration-test

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans
.PHONY: compose-down

docker-rm-volume: ### remove docker volume
	docker volume rm go-clean-template_pg-data
.PHONY: docker-rm-volume

migrate-create-db:  ### create new migration
	dbmate create
.PHONY: migrate-create-db

migrate-up: ### migration up
	dbmate up
.PHONY: migrate-up

migrate-down: ### migration down
	dbmate down
.PHONY: migrate-down

delete-container:
	docker ps -q --filter name=$(CONTAINER_NAME) | xargs -r docker stop && docker ps -aq --filter name=$(CONTAINER_NAME) | xargs -r docker rm

build-prod:
	docker build -t tongsang_be:$(APP_VERSION) . --no-cache

run-prod-docker:
	@make delete-container && docker run -d -v ./internal/log:/internal/log -v ./.env:/.env -p $(APP_PORT):8080 --name $(CONTAINER_NAME) tongsang_be:$(APP_VERSION)

reload-prod-docker:
	docker restart $(CONTAINER_NAME)

docker-seed:
	docker exec -it tongsang_be-dev go run cmd/seeder/main.go $(TABLE_NAME)

gen-proto:
	protoc --go_out=./services/event/internal/delivery --go-grpc_out=./services/event/internal/delivery ./services/event/internal/delivery/proto/*.proto

bin-deps:
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@latest

watch:
	CompileDaemon --build="go build -o main services/event/cmd/main.go" --command="./main"
.PHONY: watch