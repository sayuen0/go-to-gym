.PHONY: local swaggo test fmt

# ==============================================================================
# Go migrate mysql

force:
	migrate -database 'mysql://root:root@tcp(127.0.0.1:3306)/go_to_gym' -path tools/db/migrate force 1

version:
	migrate -database 'mysql://root:root@tcp(127.0.0.1:3306)/go_to_gym' -path tools/db/migrate version

migrate_up:
	migrate -database 'mysql://root:root@tcp(127.0.0.1:3306)/go_to_gym' -path tools/db/migrate up 1

migrate_down:
	migrate -database 'mysql://root:root@tcp(127.0.0.1:3306)/go_to_gym' -path tools/db/migrate down 1


boiler:
	sqlboiler mysql

# ==============================================================================
# Tools commands

swaggo:
	echo "Starting swagger generating"
	swag init -g **/**/*.go

lint:
	golangci-lint run

fmt:
	gofumpt -l -w .


# ==============================================================================
# Main
debug:
	air

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...

# ==============================================================================
# Modules support
tidy:
	go mod tidy

# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)