SHELL := powershell.exe
.SHELLFLAGS := -NoProfile -Command

DATE := $(shell $$now = (Get-Date).ToUniversalTime() ; Get-Date $$now -UFormat %Y-%m-%dT%H:%M:%SZ)
# Hyphen before the command tells make to ignore the error code since Remove-Item returns 2 if the
# target folder does not exist - even if ErrorAction is Ignore or Continue
RMCMD := -Remove-Item -ErrorAction Ignore -Force -Recurse -LiteralPath
MKDIRCMD := -New-Item >$$NUL -Type Directory -ErrorAction Ignore -Name
LINKCMD := New-Item >$$NUL -Type SymbolicLink -Name '$(OUTPUT_LINK)' -Target '$(OUTPUT_BIN_DIR)/$(OUTPUT_BINARY)'

XBUILDCMD := \
	$$Env:GOOS = 'linux' ; $$Env:GOARCH = 'amd64' ; $(MAKE) build ; \
	$$Env:GOOS = 'linux' ; $$Env:GOARCH = 'arm64' ; $(MAKE) build ; \
	$$Env:GOOS = 'linux' ; $$Env:GOARCH = 'arm' ; $$Env:GOARM = 7 ; $(MAKE) build ; \
	$$Env:GOOS = 'windows' ; $$Env:GOARCH = 'amd64' ; $(MAKE) build ; \
	$$Env:GOOS = 'darwin' ; $$Env:GOARCH = 'amd64' ; $(MAKE) build ; \
	$$Env:GOOS = 'darwin' ; $$Env:GOARCH = 'arm64' ; $(MAKE) build

XCOMPRESSCMD := \
	$$Env:GOOS = 'linux' ; $$Env:GOARCH = 'amd64' ; $(MAKE) compress ; \
	$$Env:GOOS = 'linux' ; $$Env:GOARCH = 'arm64' ; $(MAKE) compress ; \
	$$Env:GOOS = 'linux' ; $$Env:GOARCH = 'arm' ; $$Env:GOARM = 7 ; $(MAKE) compress ; \
	$$Env:GOOS = 'windows' ; $$Env:GOARCH = 'amd64' ; $(MAKE) compress ; \
	$$Env:GOOS = 'darwin' ; $$Env:GOARCH = 'amd64' ; $(MAKE) compress ; \
	$$Env:GOOS = 'darwin' ; $$Env:GOARCH = 'arm64' ; $(MAKE) compress

# This is a hack to make up for the lack of a -s/--xform option in Windows tar
TARGZCMD := tar -cf '$(OUTPUT_DIST_DIR)/$(OUTPUT_ARCHIVE:%.gz=%)' $(EXTRA_ASSETS) \
	-C '$(OUTPUT_BIN_DIR)' '$(OUTPUT_BINARY)' ; \
	7z rn -bso0 -bsp0 '$(OUTPUT_DIST_DIR)/$(OUTPUT_ARCHIVE:%.gz=%)' '$(OUTPUT_BINARY)' '$(OUTPUT_LINK)' ; \
	7z a -bso0 -bsp0 -tgzip -sdel '$(OUTPUT_DIST_DIR)/$(OUTPUT_ARCHIVE)' '$(OUTPUT_DIST_DIR)/$(OUTPUT_ARCHIVE:%.gz=%)'
