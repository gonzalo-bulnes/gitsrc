define GOLANG_CMD
	docker run --rm -it \
		--volume "$$GOPATH":/gopath \
		--volume "$$(pwd)":/app \
		--env "GOPATH=/gopath" \
		--workdir /app \
		golang:1.12-alpine
endef

.PHONY: all
all: clean build

.PHONY: build
build: gitsrc-linux-amd64 gitsrc-darwin-amd64

.PHONY: clean
clean:
	@rm -f gitsrc-linux-amd64
	@rm -f gitsrc-darwin-amd64

.PHONY: install-for-testing
install-for-testing:
	go get -t ./...
	go get github.com/redbubble/go-passe

gitsrc-linux-amd64: *.go **/*.go
	@${GOLANG_CMD} \
		sh -c 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a --installsuffix cgo --ldflags="-s" -o gitsrc-linux-amd64 ./cmd'

gitsrc-darwin-amd64: *.go **/*.go
	@${GOLANG_CMD} \
		sh -c 'CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a --installsuffix cgo --ldflags="-s" -o gitsrc-darwin-amd64 ./cmd'

.PHONY: test
test:
	go test -json ./... | go-passe
