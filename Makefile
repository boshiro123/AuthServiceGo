generate-protos:
	protoc -I . pkg/proto/schemas/sso.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative


install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
	&& go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


run:
	go run cmd/auth-service/main.go

config:
	export CONFIG_PATH=./compose/env/local/config-local.yaml

db:
	docker compose \
  --env-file compose/env/local/.env \
  -f compose/base/docker-compose.db.yml \
  up -d