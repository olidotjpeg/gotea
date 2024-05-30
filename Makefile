.PHONY: run
## run: runs the api
run:
	cd backend && go mod tidy
	docker compose up --build

.PHONY: migrate-up
## migrate-up: runs the migration up script
migrate-up: 
	migrate -path migrations -database "postgresql://docker:postgres123@localhost/gotea?sslmode=disable" -verbose up

.PHONY: migrate-down
## migrate-down: runs the migration down script
migrate-down:
	migrate -path migrations -database "postgresql://docker:postgres123@localhost/gotea?sslmode=disable" -verbose down

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'