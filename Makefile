BINDIR=$(CURDIR)/bin
MODNAME=ehpc.io/bull-rider/exchange-candle-service

GO=go
GOBUILD=$(GO) build
GOFMT=$(GO) fmt
GOVET=$(GO) vet
GOTEST=$(GO) test
GOPATH=$(shell go env GOPATH)

all: clean format lint test build

clean:
	@rm -rf $(BINDIR)

format:
	$(GOFMT) ./...

lint:
	$(GOPATH)/bin/golangci-lint run

test:
	$(GOTEST) ./...

cover:
	($GOTEST) -cover ./...

build:
	$(GOBUILD) -o $(BINDIR)/binance-candle-service $(MODNAME)/cmd/binance
