
migrate-create-db:  ### create new migration
	dbmate create
.PHONY: migrate-create-db

migrate-up: ### migration up
	dbmate up
.PHONY: migrate-up

migrate-down: ### migration down
	dbmate down
.PHONY: migrate-down

gen-proto:
	protoc --go_out=./services/event/internal/delivery --go-grpc_out=./services/event/internal/delivery ./services/event/internal/delivery/proto/*.proto

watch:
	CompileDaemon --build="go build -o main services/event/cmd/main.go" --command="./main"
.PHONY: watch

test:
	go test -cover ./...   
.PHONY: test