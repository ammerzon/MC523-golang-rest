# OS
OSNAME 				    :=
BINARY_NAME_FILE  :=
ifeq ($(OS),Windows_NT)
	OSNAME=windows
else
	UNAME_S :=$(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OSNAME=linux
	endif
	ifeq ($(UNAME_S),Darwin)
		OSNAME=darwin
	endif
endif

# Env
GOCMD=go
GOARCH=amd64
BINARY_NAME=server
BINARY_NAME_FILE =./dist/$(OSNAME)/$(BINARY_NAME)
BINARY_NAME_LINUX=./dist/linux/$(BINARY_NAME)
BINARY_NAME_MACOS=./dist/darwin/$(BINARY_NAME)
BINARY_NAME_WIN=./dist/windows/$(BINARY_NAME)
IMAGE=golang-rest
prebuild:
	mkdir -p ./dist/$(OSNAME)/
prebuild-all:
	mkdir -p $(BINARY_NAME_LINUX)
	mkdir -p $(BINARY_NAME_MACOS)
	mkdir -p $(BINARY_NAME_WIN)
_dist_os:
	$(GOCMD) build -o $(BINARY_NAME_FILE) ./cmd/...
build: prebuild _dist_os
build-linux:
	CC="x86_64-linux-musl-gcc" CXX="x86_64-linux-musl-g++" GOOS=linux $(GOCMD) build -o $(BINARY_NAME_LINUX) ./cmd/...
build-mac:
	GOOS=darwin $(GOCMD) build -o $(BINARY_NAME_MACOS) ./cmd/...
build-win:
	CC="x86_64-w64-mingw32-gcc" GOOS=windows $(GOCMD) build -o $(BINARY_NAME_WIN) ./cmd/...
test:
	$(GOCMD) test -v ./cmd/...
clean:
	$(GOCMD) clean ./...
	rm -rf ./dist/
download:
	go mod download
build-all: build-mac build-win build-linux
all: test prebuild-all build-all
build-image:
	docker build -t ${DOCKER_USERNAME}/$(IMAGE):latest .
	GIT_SHA="$(git rev-parse --short HEAD)"
	docker tag ${DOCKER_USERNAME}/$(IMAGE):latest ${DOCKER_USERNAME}/$(IMAGE):$GIT_SHA
	docker push ${DOCKER_USERNAME}/$(IMAGE):latest
	docker push ${DOCKER_USERNAME}/$(IMAGE):$GIT_SHA
run:
	$(GOCMD) run ./cmd start
build-stack:
	docker-compose build
run-docker:
	docker-compose up --build