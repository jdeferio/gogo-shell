.DEFAULT_GOAL := help
BINARY_NAME=main.out

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  build                 Command that builds the underlying application code/artifacts"
	@echo "  clean                 Used for projects to clean up build artifacts or other resources"
	@echo "  lint                  Command that the application can use to ensure that the code conforms to a general standards"
	@echo "  test                  Command that runs the necessary tests against the built application code/artifacts"
	@echo "  format                Command to invoke code formatting using Go's built-in"

lint:
	$(info lint)
	golangci-lint run 

build: clean lint
	$(info Running [$@])
	$(info Environment [$*])
	go build -o ${BINARY_NAME} main.go

clean:
	$(info cleaning)
	go clean
	rm ${BINARY_NAME}

test: build
	$(info test)
	go test -v main.go

format: lint
	$(info formatting)
	go fmt

