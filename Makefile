

.PHONY: build
build:
	go build -o bashelp cmd/main.go

.PHONY: install
install: build
	mv bashelp ~/.local/bin/bashelp
