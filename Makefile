GOCMD ?= go
GOFMT ?= gofmt

PACKAGE := $(shell $(GOCMD) list)
NAME := $(notdir $(PACKAGE))
FULLCOMMIT := $(shell git rev-parse HEAD)
TAG := $(shell git describe --tags --dirty)
VERSION := $(TAG:v%=%)
BUILD_VERSION := $(subst -,.,$(VERSION))

GOOS := $(shell $(GOCMD) env GOOS)
GOARCH := $(shell $(GOCMD) env GOARCH)
GOARM := $(shell $(GOCMD) env GOARM)
GOEXE := $(shell $(GOCMD) env GOEXE)

BASENAME = $(NAME)-$(BUILD_VERSION)-$(GOOS)-$(GOARCH)$(GOARM)
OUTPUT_BINARY = $(BASENAME)$(GOEXE)
OUTPUT_LINK = $(NAME)$(GOEXE)
OUTPUT_BIN_DIR = bin
OUTPUT_DIST_DIR = dist

DOCKER_REGISTRY ?= r.deso.tech
DOCKER_PROJECT ?= whoami
DOCKER_IMAGE := $(DOCKER_REGISTRY)/$(DOCKER_PROJECT)/$(NAME)

ARCHIVE_EXT := .tar.gz

OUTPUT_ARCHIVE=$(BASENAME)$(ARCHIVE_EXT)
EXTRA_ASSETS := 'static' 'template'

include unix.mk

LDFLAGS = $(ADDITIONAL_LDFLAGS) -s -w \
	-X '$(PACKAGE)/app.version=$(VERSION)' \
	-X '$(PACKAGE)/app.commit=$(FULLCOMMIT)' \
	-X '$(PACKAGE)/app.date=$(DATE)'

BUILD_FLAGS = -trimpath -ldflags "$(LDFLAGS)" -o '$(OUTPUT_BIN_DIR)/$(OUTPUT_BINARY)'

GOBUILD = $(GOCMD) build $(BUILD_FLAGS)

COMPRESSCMD := $(TARGZCMD)

.DEFAULT_GOAL := link

.PHONY: build
build:
	$(GOBUILD)

.PHONY: xbuild
xbuild:
	$(XBUILDCMD)

.PHONY: link
link: build
	$(LINKCMD)

.PHONY: compress
compress: build
	$(MKDIRCMD) '$(OUTPUT_DIST_DIR)'
	$(COMPRESSCMD)

.PHONY: xcompress
xcompress:
	$(XCOMPRESSCMD)

.PHONY: clean
clean:
	$(GOCMD) clean
	$(RMCMD) '$(OUTPUT_BIN_DIR)'
	$(RMCMD) '$(OUTPUT_DIST_DIR)'

.PHONY: fmt
fmt:
	$(GOFMT) -s -w .

.PHONY: docker-linux
docker-linux:
	docker buildx build --push --platform linux/amd64 -t '$(DOCKER_IMAGE):amd64-$(VERSION)' -t '$(DOCKER_IMAGE):amd64' .
	docker buildx build --push --platform linux/arm64/v8 -t '$(DOCKER_IMAGE):arm64v8-$(VERSION)' -t '$(DOCKER_IMAGE):arm64v8' .

.PHONY: docker-shared
docker-shared:
	docker manifest create '$(DOCKER_IMAGE):$(VERSION)' \
		'$(DOCKER_IMAGE):amd64-$(VERSION)' \
		'$(DOCKER_IMAGE):arm64v8-$(VERSION)'
	docker manifest push '$(DOCKER_IMAGE):$(VERSION)'

	docker manifest create '$(DOCKER_IMAGE):latest' \
		'$(DOCKER_IMAGE):amd64' \
		'$(DOCKER_IMAGE):arm64v8'
	docker manifest push '$(DOCKER_IMAGE):latest'
