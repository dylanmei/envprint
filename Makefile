TARGET   ?= $(shell basename `git rev-parse --show-toplevel`)
VERSION  ?= $(shell git describe --tags --always)
REVISION ?= $(shell git rev-parse HEAD)
LD_FLAGS ?= -s \
	-X github.com/dylanmei/envprint/main.Name=$(TARGET) \
	-X github.com/dylanmei/envprint/main.Version=$(VERSION) \
	-X github.com/dylanmei/envprint/main.Revision=$(REVISION)

default: test build

test:
	go test -v -cover -run=$(RUN)

build: clean
	@go build -v \
		-ldflags "$(LD_FLAGS)+local_changes" \
		-o bin/$(TARGET) .

clean:
	@rm -rf bin/
