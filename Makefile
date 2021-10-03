LOCAL_BIN:=$(CURDIR)/bin
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.42.1

BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)
COMMIT:=$(shell git rev-parse --short HEAD)
PKG_PATH:=url_shortener
TAG:=$(shell git describe --tags |cut -d- -f1)

LDFLAGS:=-ldflags "-X $(PKG_PATH)/internal/version.gitTag=$(TAG) -X $(PKG_PATH)/internal/version.gitCommit=$(COMMIT) -X $(PKG_PATH)/internal/version.gitBranch=$(BRANCH)"

# Check local bin version
ifneq ($(wildcard $(GOLANGCI_BIN)),)
GOLANGCI_BIN_VERSION:=$(shell $(GOLANGCI_BIN) --version)
ifneq ($(GOLANGCI_BIN_VERSION),)
GOLANGCI_BIN_VERSION_SHORT:=$(shell echo "$(GOLANGCI_BIN_VERSION)" | sed -E 's/.* version (.*) built from .* on .*/\1/g')
else
GOLANGCI_BIN_VERSION_SHORT:=0
endif
ifneq "$(GOLANGCI_TAG)" "$(word 1, $(sort $(GOLANGCI_TAG) $(GOLANGCI_BIN_VERSION_SHORT)))"
GOLANGCI_BIN:=
endif
endif

.PHONY: build
build: dep
	go build $(LDFLAGS) -o $(LOCAL_BIN)/url_shortener ./cmd/url_shortener

.PHONY: dep
dep:
	go mod download

.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info #Downloading golangci-lint v$(GOLANGCI_TAG))
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG) && \
		go build -ldflags "-X 'main.version=$(GOLANGCI_TAG)' -X 'main.commit=test' -X 'main.date=test'" -o $(LOCAL_BIN)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint && \
		rm -rf $$tmp
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

.PHONY: lint
lint: install-lint
	$(info #Running lint...)
	$(GOLANGCI_BIN) run --config=.golang-lint.ci.yml ./...

.PHONY: run
run:
	go run ./cmd/url_shortener

.PHONY: test
test:
	go test -v ./...

.DEFAULT_GOAL=run