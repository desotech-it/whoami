GOCMD ?= go
GOFMT ?= gofmt

PACKAGE := $(shell $(GOCMD) list)
NAME := $(notdir $(PACKAGE))
FULLCOMMIT := $(shell git rev-parse HEAD)
TAG := $(shell git describe --long --dirty)
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
