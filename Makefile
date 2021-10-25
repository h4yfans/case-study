SHELL = /bin/sh

help:
	@echo 'Management commands for Case Study Service:'
	@echo
	@echo 'Usage:'
	@echo '   make createtable <file_name>	Generate new db table'
	@echo '   make lint			Run linter'
	@echo



.PHONY: createtable
createtable:
	migrate create -ext sql -dir db/migrations -seq create_$(filter-out $@,$(MAKECMDGOALS))_table
%:


.PHONY: lint	
lint:
	golangci-lint run







