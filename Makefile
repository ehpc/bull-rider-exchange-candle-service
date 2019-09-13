BINDIR=$(CURDIR)/bin
MODNAME=ehpc.io/bull-rider-exchange-candle-service

GO=go
GOBUILD=$(GO) build

all: clean build

build:
	$(GOBUILD) -o $(BINDIR)/binance-candle-service $(MODNAME)/cmd/binance-candle-service

clean:
	@rm -rf $(BINDIR)
