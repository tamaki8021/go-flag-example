EXEC_FILE = cli-tool
VERSION = 0.0.1
BUILD_FLAGS = -ldflags "\
                -X main.Version=$(VERSION) \
                "

all: clean deps test build

fmt:
	go fmt ./...

deps:
	go get -d -v ./...

test:
	go test -v ./...

build: fmt
	go build $(BUILD_FLAGS) -o "$(EXEC_FILE)" .

clean:
	go clean

.PHONY: fmt clean deps build