OS          := $(shell go env GOOS)
ARCH        := $(shell go env GOARCH)

GO          := $(GOENV) go
GOBUILD     := GOOS=$(OS) GOARCH=$(ARCH) $(GO) build

TARGET_DIR  := ./target
INSTALL_PREFIX=/usr/local


APPS=music-theory

#
# environment

all: $(APPS)

.PHONY: $(APPS) install test fmt

fmt:
	go fmt ./...

test: deps
	go test ./...

$(APPS): deps test
	@echo "# Building $@"
	$(GOBUILD) \
		-ldflags "$(LDFLAGS)" \
		-o $(TARGET_DIR)/$@.$(OS).$(ARCH) \
		$@.go

install: all
	for app in $(APPS); do \
		sudo install -m 0755 $(TARGET_DIR)/$$app.$(OS).$(ARCH) $(INSTALL_PREFIX)/bin/$$app; done \

deps:
	go get -v -t ./...
	mkdir -p $(TARGET_DIR)
