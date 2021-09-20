DATE = $(shell LC_ALL=C.UTF-8 date -u +'%Y-%m-%dT%H:%M:%SZ')
RMCMD = rm -rf
LINKCMD = ln -sf $(OUTPUT_BIN_DIR)/$(OUTPUT_BINARY) $(OUTPUT_LINK)
MKDIRCMD = mkdir -p

XBUILDCMD := \
	GOOS=linux GOARCH=amd64 $(MAKE) build ; \
	GOOS=linux GOARCH=arm64 $(MAKE) build ; \
	GOOS=linux GOARCH=arm GOARM=7 $(MAKE) build ; \
	GOOS=windows GOARCH=amd64 $(MAKE) build ; \
	GOOS=darwin GOARCH=amd64 $(MAKE) build ; \
	GOOS=darwin GOARCH=arm64 $(MAKE) build

XCOMPRESSCMD := \
	GOOS=linux GOARCH=amd64 $(MAKE) compress ; \
	GOOS=linux GOARCH=arm64 $(MAKE) compress ; \
	GOOS=linux GOARCH=arm GOARM=7 $(MAKE) compress ; \
	GOOS=windows GOARCH=amd64 $(MAKE) compress ; \
	GOOS=darwin GOARCH=amd64 $(MAKE) compress ; \
	GOOS=darwin GOARCH=arm64 $(MAKE) compress

ifeq ($(shell uname), Linux)
	TRANSFORM_OPT := --xform 'flags=r;s|^$(OUTPUT_BINARY)$$|$(OUTPUT_LINK)|'
else
	TRANSFORM_OPT := -s '/^$(OUTPUT_BINARY)$$/$(OUTPUT_LINK)/'
endif

TARGZCMD := tar $(TRANSFORM_OPT) -czf '$(OUTPUT_DIST_DIR)/$(OUTPUT_ARCHIVE)' \
	$(EXTRA_ASSETS) -C '$(OUTPUT_BIN_DIR)' '$(OUTPUT_BINARY)'
