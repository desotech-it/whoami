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

ifeq ($(GOOS), windows)
	ARCHIVE_EXT := .zip
else
	ARCHIVE_EXT := .tar.gz
endif

OUTPUT_ARCHIVE=$(BASENAME)$(ARCHIVE_EXT)
EXTRA_ASSETS := 'static' 'template'

ifeq ($(OS), Windows_NT)
	include windows.mk
else
	include unix.mk
endif

LDFLAGS = $(ADDITIONAL_LDFLAGS) -s -w \
	-X '$(PACKAGE)/app.version=$(VERSION)' \
	-X '$(PACKAGE)/app.commit=$(FULLCOMMIT)' \
	-X '$(PACKAGE)/app.date=$(DATE)'

BUILD_FLAGS = -trimpath -ldflags "$(LDFLAGS)" -o '$(OUTPUT_BIN_DIR)/$(OUTPUT_BINARY)'

GOBUILD = $(GOCMD) build $(BUILD_FLAGS)

ifeq ($(ARCHIVE_EXT), .tar.gz)
	COMPRESSCMD := $(TARGZCMD)
else
	COMPRESSCMD := 7z a -bso0 -bsp0 -tzip '$(OUTPUT_DIST_DIR)/$(OUTPUT_ARCHIVE)' $(EXTRA_ASSETS) '$(OUTPUT_BIN_DIR)/$(OUTPUT_BINARY)' ; \
		7z rn -bso0 -bsp0 '$(OUTPUT_DIST_DIR)/$(OUTPUT_ARCHIVE)' '$(OUTPUT_BIN_DIR)/$(OUTPUT_BINARY)' '$(OUTPUT_LINK)'
endif

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
	docker build --platform linux/amd64 -t '$(DOCKER_IMAGE):amd64-$(VERSION)' -t '$(DOCKER_IMAGE):amd64' .
	docker build --platform linux/arm64/v8 -t '$(DOCKER_IMAGE):arm64v8-$(VERSION)' \
		-t '$(DOCKER_IMAGE):arm64v8' .
	docker push '$(DOCKER_IMAGE):amd64-$(VERSION)'
	docker push '$(DOCKER_IMAGE):amd64'
	docker push '$(DOCKER_IMAGE):arm64v8-$(VERSION)'
	docker push '$(DOCKER_IMAGE):arm64v8'

.PHONY: docker-windows
docker-windows:
	docker build --platform linux/amd64 -t '$(DOCKER_IMAGE):$(VERSION)-windowsservercore' \
		-t '$(DOCKER_IMAGE):windowsservercore' -f windows.dockerfile .
	docker push '$(DOCKER_IMAGE):$(VERSION)-windowsservercore'
	docker push '$(DOCKER_IMAGE):windowsservercore'

.PHONY: docker-shared
docker-shared:
	docker create manifest '$(DOCKER_IMAGE):$(VERSION)' \
		'$(DOCKER_IMAGE):amd64-$(VERSION)' \
		'$(DOCKER_IMAGE):arm64v8-$(VERSION)' \
		'$(DOCKER_IMAGE):$(VERSION)-windowsservercore'
	docker create push '$(DOCKER_IMAGE):$(VERSION)'

	docker create manifest '$(DOCKER_IMAGE):latest' \
		'$(DOCKER_IMAGE):amd64' \
		'$(DOCKER_IMAGE):arm64v8' \
		'$(DOCKER_IMAGE):windowsservercore'
	docker create push '$(DOCKER_IMAGE):latest'
