SHELL := /bin/bash
version := $(shell git rev-list --count HEAD)
commit := $(shell git describe --always --long --dirty)
built_at := $(shell date +%FT%T%z)
built_by := ${USER}

flags := -gcflags="all=-N -l -c 2"
ldflags := -X main.version=v${version} -X main.commit=${commit}
ldflags += -X main.builtAt=${built_at} -X main.builtBy=${built_by}

upload-path-linux := 's3://tools.harmony.one/release/linux-x86_64/oasis-spammer'

dist := ./dist/oasis-spammer
env := GO111MODULE=on
DIR := ${CURDIR}

all:
	$(env) go build -o $(dist) -ldflags="$(ldflags)" cmd/main.go

static:
	$(env) go build -o $(dist) -ldflags="$(ldflags) -w -extldflags \"-static\"" cmd/main.go

debug:
	$(env) go build $(flags) -o $(dist) -ldflags="$(ldflags)" cmd/main.go

upload-linux:static
	aws s3 cp dist/oasis-spammer ${upload-path-linux} --acl public-read

.PHONY:clean upload-linux

clean:
	@rm -f $(dist)
	@rm -rf ./dist
