MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail
DOCKER_REPO ?= "a3s"
DOCKER_IMAGE ?= "a3s"
DOCKER_TAG ?= "dev"

export GO111MODULE = on

default: lint test a3s cli

lint:
	golangci-lint run \
		--disable-all \
		--exclude-use-default=false \
		--enable=errcheck \
		--enable=goimports \
		--enable=ineffassign \
		--enable=revive \
		--enable=unused \
		--enable=structcheck \
		--enable=staticcheck \
		--enable=varcheck \
		--enable=deadcode \
		--enable=unconvert \
		--enable=misspell \
		--enable=prealloc \
		--enable=nakedret \
		--enable=typecheck \
		--enable=nilerr \
		./...

test:
	go test ./... -race -cover -covermode=atomic -coverprofile=unit_coverage.cov

	@ echo "Converting the coverage file..."
	gocov convert ./unit_coverage.cov | gocov-xml > ./coverage.xml

sec:
	gosec -quiet ./...

codegen:
	cd pkgs/api && make codegen

a3s:
	cd cmd/a3s && CGO_ENABLED=0 go build -ldflags="-w -s"

a3s_linux:
	cd cmd/a3s && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"

cli:
	cd cmd/a3sctl && CGO_ENABLED=0 go install -ldflags="-w -s"

cli_linux:
	cd cmd/a3sctl && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -ldflags="-w -s"

docker: a3s_linux
	mkdir -p docker/in
	cp cmd/a3s/a3s docker/in
	cd docker && docker build -t ${DOCKER_REPO}/${DOCKER_IMAGE}:${DOCKER_TAG} .
