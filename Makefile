run: fmt
	go run *.go $(ARGS)

fmt:
	go fmt ./...

build:
	go build -o iterm2 *.go
