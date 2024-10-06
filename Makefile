NAMESPACE = `echo go-service`
BUILD_VERSION = `git describe --tag`

.PHONY: build
build:
	@go build -ldflags "-X main.Namespace=${NAMESPACE} \
		-X main.BuildVersion=${BUILD_VERSION}" \
		-race --tags=dynamic -o ./build/app ./cmd/

.PHONY: remove-swag
remove-swag:
	@/bin/rm -f ./docs/docs.go ./docs/swagger.json ./docs/swagger.yaml

		
.PHONY: swag
swag: remove-swag
	@`go env GOPATH`/bin/swag init -g ./cmd/main.go

.PHONY: run
run: build
	@./build/app

.PHONY: run-debug
run-debug: build
	@./build/app -debug

