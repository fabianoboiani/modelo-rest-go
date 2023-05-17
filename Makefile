APP_NAME=modelo-rest-go
VERSION=0.0.1

.PHONY: deps
deps:
	go mod download

.PHONY: get-docs
get-docs:
	go get -u github.com/swaggo/swag/cmd/swag

.PHONY: docs
docs: get-docs
	swag init --dir cmd/${APP_NAME} --parseDependency --output api

.PHONY: build
build:
	go build -o bin/restapi cmd/${APP_NAME}/main.go

.PHONY: run
run:
	go run cmd/${APP_NAME}/main.go -e development

.PHONY: run-prod
run-prod:
	go run cmd/${APP_NAME}/main.go -e production

.PHONY: test
test:
	go test -v ./test/...

.PHONY: build-docker
build-docker: build
	docker build . -t api-rest

.PHONY: run-docker
run-docker: build-docker
	docker run -p 3000:3000 api-rest