.PHONY: run default

all: default

run: setting
	go run ./cmd/web -addr=$$SNIPPETBOX_ADDR

setting:
	export SNIPPETBOX_ADDR=:9999

default:
	go run ./cmd/web

