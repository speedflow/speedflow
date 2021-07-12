APP_NAME:=speedflow
APP_TAG:=dev
BIN_DIR:=bin
CONTAINER_FILE:=Containerfile
CONTAINER_ENGINE:=docker

generate: ## Run go generate
	go generate

lint: ## Lint code
	golangci-lint run

test: ## Test packages
	go test -count=1 -cover -coverprofile=coverage.out -v ./...

coverage: ## Test coverage with default output
	go tool cover -func=coverage.out

coverage-html: ## Test coverage with html output
	go tool cover -html=coverage.html

clean: ## Clean project
	rm -Rf ./${BIN_DIR}
	rm -Rf coverage.out

build: clean ## Build local binary
	mkdir -p ./${BIN_DIR}
	go build -o ./${BIN_DIR} ./cmd/${APP_NAME}

build-image: ## Build local image
	${CONTAINER_ENGINE} build -f ${CONTAINER_FILE} -t ghcr.io/${APP_NAME}/${APP_NAME}:${APP_TAG} .

run-container: ## Build local container
	${CONTAINER_ENGINE} run --rm -it ghcr.io/${APP_NAME}/${APP_NAME}:${APP_TAG}

run: build ## Run local binary
	./${BIN_DIR}/${APP_NAME}

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: generate lint test coverage coverage-html clean build build-image run run-container env-up env-down env-logs help
