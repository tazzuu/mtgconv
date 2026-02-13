SHELL:=/bin/bash

format:
	go fmt ./...

tidy:
	go mod tidy

# cli interface for the program
SRC:=cmd/mtgconv

# need at least 1 tag for this to work
GIT_TAG:=$(shell git describe --tags)
# name of output binary file
BIN:=mtgconv

# compile for the current system
$(BIN): build

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


# test:
# 	go test ./...

# run a select set of unit tests
test:
	go test ./pkg/mtgconv2/sources/shinycsv


# run all end to end test cases
SHINY_FILE:=ShinyExport-7ae9f94e61ef4c9785c441ca55df9194.csv
test-cases: $(BIN)
	./$(BIN) --verbose --output-dir my-decks-tmp search --page-size 2 --sort-type views archidekt.com
	./$(BIN) --output-dir my-decks-tmp --verbose --save-json --user-agent "$$MOXKEY" convert --output-filename auto --output-format dck https://moxfield.com/decks/TiS_BYhhnUWp_3aq24hbFA
	[ -f "$(SHINY_FILE)" ] && ./$(BIN) convert --input-source shiny-csv  "$(SHINY_FILE)" || exit 0
