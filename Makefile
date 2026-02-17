.PHONY: run dev docker build test coverage coverage-html

# Variables
BINARY_NAME=main
COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html

run:
	go run cmd/main.go

dev:
	go run -gcflags=all="-N -l" cmd/main.go

docker:
	docker compose up --build -d

build:
	go build -o $(BINARY_NAME) .

test:
	go test -v ./...

coverage:
	go test ./... -coverprofile=$(COVERAGE_FILE)
	go tool cover -func=$(COVERAGE_FILE)

coverage-html: coverage
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Coverage report: $(COVERAGE_HTML)"