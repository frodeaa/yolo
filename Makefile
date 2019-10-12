VERSION=0.0.1

DEP?=dep
GO?=go
GOFLAGS?=
GOSRC!=find . -name '*.go'

yolo: $(GOSRC)
	$(DEP) ensure
	$(GO) build $(GOFLAGS) \
		-ldflags "-X main.VERSION=$(VERSION)" \
		-o $@

dep:
	$(GO) get -v -u github.com/golang/dep/cmd/dep

test:
	$(GO) test -v ./...

all: yolo

RM?=rm -f

clean:
	$(RM) yolo

.DEFAULT_GOAL := all

.PHONY: all dep clean test
