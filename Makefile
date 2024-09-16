include .env
LOCAL_BIN:=$(CURDIR)/bin

build:
	go build -o ./bin/service ./cmd/app/main.go
	go build -o ./bin/console ./cmd/console/main.go