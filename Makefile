SHELL := /bin/bash
APP_DOCKER_COMPOSE=docker-compose.yml


# // To generate a private/public key PEM file.
# openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
# openssl rsa -pubout -in private.pem -out public.pem
# ./sales-admin genkey


# hey -m GET -c 100 -n 10000 -H "Authorization: Bearer ${TOKEN}" http://localhost:3000/v1/users
# zipkin: http://localhost:9411
# expvarmon -ports=":4000" -vars="build,requests,goroutines,errors,mem:memstats.Alloc"


run:
	go run app/sales-api/main.go

tests:
	go test ./...

tidy:
	go mod tidy
	go mod vendor




all: serve


.PHONY: build
build:
	docker-compose -f ${APP_DOCKER_COMPOSE} build 

.PHONY: serve
serve:
	@docker-compose -f ${APP_DOCKER_COMPOSE} up --remove-orphans

.PHONY: migrate
migrate:
	@docker-compose -f ${APP_DOCKER_COMPOSE} run --rm --entrypoint="./admin-cli migrate" sales-api

.PHONY: seed
seed:
	@docker-compose -f ${APP_DOCKER_COMPOSE} run --rm --entrypoint="./admin-cli seed" sales-api