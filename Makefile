echo "Variables"
GO_VERSION := 1.23.3
DOCKER_IMAGE := my-webserver
DOCKER_TAG := latest

echo "Install Go dependencies"
.PHONY: install-deps
install-deps:
	@echo "Installing Go dependencies..."
	go mod tidy

echo "Build Go application"
.PHONY: build
build:
	@echo "Building Go application..."
	go build -o webserver ./cmd/webserver/main.go

echo "Run tests"
.PHONY: test
test:
	@echo "Running Go tests..."
	go test ./...

echo "Lint Go code"
.PHONY: lint
lint:
	@echo "Running Go lint..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...

echo "Build Docker image"
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

echo "Run Docker container"
.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

echo "Run the application (inside Docker)"
.PHONY: run
run: docker-build docker-run

echo "Docker Compose up"
.PHONY: docker-compose-up
docker-compose-up:
	@echo "Running Docker Compose..."
	docker-compose up -d

echo "Docker Compose down"
.PHONY: docker-compose-down
docker-compose-down:
	@echo "Stopping Docker Compose..."
	docker-compose down

echo "Clean the build (remove binaries)"
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -f webserver
