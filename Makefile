install:
	sudo apt install docker-compose \
	&& sudo usermod -aG docker $$USER \
	&& sudo service docker restart

mvn:
	mvn clean package



build:
	mvn clean package -DskipTests
	docker compose up --build

stop:
	docker compose stop

up:
	docker compose up -d

down:
	docker compose down

db-start:
	docker compose up -d db

db-stop:
	docker compose stop db


generate-protos:
	protos/protoc -I proto proto/sso/sso.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative








