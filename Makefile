SHELL := /bin/bash
BASEDIR = $(shell pwd)

versionDir = "cobra_cli/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: gotool
	@go build -v -ldflags ${ldflags} .
clean:
	rm -f cobra_cli
gotool:
	gofmt -w .
	go vet .

help:
	@echo "make - compile the source code"
	@echo "make clean - remove binary file"
	@echo "make gotool - run go tool 'fmt' and 'vet'"

.PHONY: clean gotool help
