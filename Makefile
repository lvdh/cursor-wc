.PHONY: build b convey c cover cc coverage ccc
.PHONY: deps d init i run r test t help h
.DEFAULT_GOAL := test
all: init deps convey
build: b
b:
	go build .
convey: c
c:
	$(GOPATH)/bin/goconvey
cover: cc
cc:
	go test -cover
coverage: ccc
ccc:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
deps: d
d:
	go install github.com/smartystreets/goconvey@latest
init: i
i:
	go mod init sandbox
run: r
r:
	go run .
test: t
t:
	go test
help: h
h:
	@sed -nr 's/^([a-z]{2,}: [a-z])/\1/p' Makefile
