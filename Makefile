.PHONY: local swaggo test

# ==============================================================================
# Tools commands

swaggo:
	echo "Starting swagger generating"
	swag init -g **/**/*.go


# ==============================================================================
# Main
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