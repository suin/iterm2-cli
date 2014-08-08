run: fmt
	go run *.go $(ARGS)

fmt:
	go fmt ./...

build:
	go build -o bin/iterm2 -ldflags "-X main.version $(version)" *.go

release:
	@test $(version)
	go build -o bin/iterm2 -ldflags "-X main.version $(version)" *.go
	git add bin
	git commit -c "release $(version)"
	git tag $(version)
