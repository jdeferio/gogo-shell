.DEFAULT_GOAL := help
BUILD_PATH := cmd/gogo_shell

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  all                   Command that runs tets"
	@echo "  build                 Command that builds the underlying application code/artifacts"
	@echo "  lint                  Command that the application can use to ensure that the code conforms to a general standards"
	@echo "  format                Command to invoke code formatting using Go's built-in"
	@echo "  clean                 Used for projects to clean up build artifacts or other resources"
	@echo "  test                  Command that runs the necessary tests against the built application code/artifacts"
	@echo "  run                   Run the application locally"

all: test

vet:
	@go list ./cmd/ | grep -v vendor | xargs go vet

build: clean lint
	$(info building)
	@go build -o ${BUILD_PATH} ./cmd/main.go

lint:
	$(info lint)
	# brew install golangci-lint
	@golangci-lint run 

format: lint
	$(info formatting)
	go fmt cmd/

clean:
	$(info cleaning)
	go clean
	@rm -rf dist

test: build
	$(info test)
	@go test -v ./cmd/ -coverprofile=coverage.out

run:
	@go run ./cmd/

.PHONY: run test clean format lint build vet all
