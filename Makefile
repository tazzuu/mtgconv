SHELL:=/bin/bash

format:
	go fmt ./...

tidy:
	go mod tidy

# cli interface for the program
SRC:=cmd/mtgconv


test:
	go test ./...


# need at least 1 tag for this to work
GIT_TAG:=$(shell git describe --tags)
# name of output binary file
BIN:=mtgconv

# compile for the current system
build:
	CGO_ENABLED=0 go build -ldflags="-X 'main.version=$(GIT_TAG)'" -trimpath -o ./$(BIN) ./$(SRC)
.PHONY:build

# cross-compile for all available OS and arch types
# use this for releases
build-all:
	mkdir -p build ; \
	for os in darwin linux windows; do \
	for arch in amd64 arm64; do \
	output="build/$(BIN)-v$(GIT_TAG)-$$os-$$arch" ; \
	if [ "$${os}" == "windows" ]; then output="$${output}.exe"; fi ; \
	echo "building: $$output" ; \
	CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build  -ldflags="-X 'main.version=$(GIT_TAG)'" -trimpath -o "$${output}" ./$(SRC) ; \
	done ; \
	done


# brew install goreleaser
release:
	goreleaser release --snapshot --clean