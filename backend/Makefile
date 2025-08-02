# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=main
BINARY_UNIX=$(BINARY_NAME)_unix

# Build targets
.PHONY: all build clean test coverage deps lint run dev docker-build docker-run

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/main.go

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -cover ./...

deps:
	$(GOGET) -d ./...
	go mod tidy

lint:
	golangci-lint run

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/main.go
	./$(BINARY_NAME)

dev:
	air

# Database operations
migrate-up:
	migrate -path migrations -database "postgresql://localhost/smart_learning_dev?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgresql://localhost/smart_learning_dev?sslmode=disable" down

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

# Docker operations
docker-build:
	docker build -t smart-learning-backend .

docker-run:
	docker run -p 8080:8080 smart-learning-backend

# Build for Linux (useful for deployment)
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd/main.go