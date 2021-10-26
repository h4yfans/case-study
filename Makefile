SHELL = /bin/sh

help:
	@echo 'Management commands for Case Study Service:'
	@echo
	@echo 'Usage:'
	@echo '   make build		Build or rebuild services'
	@echo '   make up		Create and start containers'
	@echo '   make lint		Run linter'
	@echo



.PHONY: createtable
createtable:
	migrate create -ext sql -dir db/migrations -seq create_$(filter-out $@,$(MAKECMDGOALS))_table
%:

.PHONY: lint	
lint:
	golangci-lint run

.PHONY: build
build:
	docker-compose build

.PHONY: up
up:
	docker-compose up

.PHONY: up
test:
	go test ./...


