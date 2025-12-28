# Project parameters
LIB_NAME = typr-io
VERSION = v0.1.0
REPO_URL = https://github.com/ZiedYousfi/typr-io/releases/download/$(VERSION)

# All supported platforms and architectures
PLATFORMS = linux-arm64 linux-x86_64 macos-arm64 macos-x86_64 windows-arm64 windows-x64

# Helper to fix library names (remove version suffixes and symlinks)
define FIX_LIBS
	@echo "Fixing library names in lib/$(1)..."
	@find lib/$(1) -type f -name "libtypr_io.so.*" -exec mv {} lib/$(1)/libtypr_io.so \; 2>/dev/null || true
	@find lib/$(1) -type f -name "libtypr_io.*.dylib" -exec mv {} lib/$(1)/libtypr_io.dylib \; 2>/dev/null || true
	@find lib/$(1) -type l -name "libtypr_io.*" -delete 2>/dev/null || true
endef

.PHONY: all install install-all install-current clean

all: install-all

# Install all platform libraries
install-all:
	@echo "Downloading all typr-io releases..."
	@mkdir -p include lib

	# Download and extract Linux ARM64
	@echo "Fetching linux-arm64..."
	@curl -sL $(REPO_URL)/$(LIB_NAME)-$(VERSION)-linux-arm64.tar.gz -o linux-arm64.tar.gz
	@mkdir -p tmp-linux-arm64 lib/linux-arm64
	@tar -xzf linux-arm64.tar.gz -C tmp-linux-arm64
	@cp -R tmp-linux-arm64/include/* include/ 2>/dev/null || true
	@cp -R tmp-linux-arm64/lib/lib* lib/linux-arm64/
	@rm -rf tmp-linux-arm64 linux-arm64.tar.gz
	$(call FIX_LIBS,linux-arm64)

	# Download and extract Linux x86_64
	@echo "Fetching linux-x86_64..."
	@curl -sL $(REPO_URL)/$(LIB_NAME)-$(VERSION)-linux-x86_64.tar.gz -o linux-x86_64.tar.gz
	@mkdir -p tmp-linux-x86_64 lib/linux-x86_64
	@tar -xzf linux-x86_64.tar.gz -C tmp-linux-x86_64
	@cp -R tmp-linux-x86_64/include/* include/ 2>/dev/null || true
	@cp -R tmp-linux-x86_64/lib/lib* lib/linux-x86_64/
	@rm -rf tmp-linux-x86_64 linux-x86_64.tar.gz
	$(call FIX_LIBS,linux-x86_64)

	# Download and extract macOS ARM64
	@echo "Fetching macos-arm64..."
	@curl -sL $(REPO_URL)/$(LIB_NAME)-$(VERSION)-macos-arm64.tar.gz -o macos-arm64.tar.gz
	@mkdir -p tmp-macos-arm64 lib/macos-arm64
	@tar -xzf macos-arm64.tar.gz -C tmp-macos-arm64
	@cp -R tmp-macos-arm64/include/* include/ 2>/dev/null || true
	@cp -R tmp-macos-arm64/lib/lib* lib/macos-arm64/
	@rm -rf tmp-macos-arm64 macos-arm64.tar.gz
	$(call FIX_LIBS,macos-arm64)

	# Download and extract macOS x86_64
	@echo "Fetching macos-x86_64..."
	@curl -sL $(REPO_URL)/$(LIB_NAME)-$(VERSION)-macos-x86_64.tar.gz -o macos-x86_64.tar.gz
	@mkdir -p tmp-macos-x86_64 lib/macos-x86_64
	@tar -xzf macos-x86_64.tar.gz -C tmp-macos-x86_64
	@cp -R tmp-macos-x86_64/include/* include/ 2>/dev/null || true
	@cp -R tmp-macos-x86_64/lib/lib* lib/macos-x86_64/
	@rm -rf tmp-macos-x86_64 macos-x86_64.tar.gz
	$(call FIX_LIBS,macos-x86_64)

	# Download and extract Windows ARM64
	@echo "Fetching windows-arm64..."
	@curl -sL $(REPO_URL)/$(LIB_NAME)-$(VERSION)-windows-arm64.zip -o windows-arm64.zip
	@mkdir -p tmp-windows-arm64 lib/windows-arm64
	@unzip -q windows-arm64.zip -d tmp-windows-arm64
	@cp -R tmp-windows-arm64/include/* include/ 2>/dev/null || true
	@cp -R tmp-windows-arm64/lib/* lib/windows-arm64/ 2>/dev/null || true
	@cp -R tmp-windows-arm64/bin/* lib/windows-arm64/ 2>/dev/null || true
	@rm -rf tmp-windows-arm64 windows-arm64.zip

	# Download and extract Windows x64
	@echo "Fetching windows-x64..."
	@curl -sL $(REPO_URL)/$(LIB_NAME)-$(VERSION)-windows-x64.zip -o windows-x64.zip
	@mkdir -p tmp-windows-x64 lib/windows-x64
	@unzip -q windows-x64.zip -d tmp-windows-x64
	@cp -R tmp-windows-x64/include/* include/ 2>/dev/null || true
	@cp -R tmp-windows-x64/lib/* lib/windows-x64/ 2>/dev/null || true
	@cp -R tmp-windows-x64/bin/* lib/windows-x64/ 2>/dev/null || true
	@rm -rf tmp-windows-x64 windows-x64.zip

	@echo "Done! All platform libraries installed."
	@echo ""
	@echo "Library structure:"
	@ls -la lib/

# Install only for current platform (legacy behavior)
install-current:
	$(eval OS := $(shell uname -s | tr '[:upper:]' '[:lower:]'))
	$(eval ARCH := $(shell uname -m))
	$(eval PLATFORM := $(if $(filter darwin,$(OS)),macos,linux))
	$(eval GOARCH := $(if $(filter x86_64,$(ARCH)),x86_64,$(if $(filter arm64 aarch64,$(ARCH)),arm64,$(ARCH))))
	$(eval RELEASE_DIR := $(LIB_NAME)-$(VERSION)-$(PLATFORM)-$(GOARCH))
	$(eval TARBALL := $(RELEASE_DIR).tar.gz)

	@echo "Download of the typr-io release for $(PLATFORM)-$(GOARCH)..."
	@curl -L $(REPO_URL)/$(TARBALL) -o $(TARBALL)

	@echo "Extracting..."
	@mkdir -p $(RELEASE_DIR)
	@tar -xzf $(TARBALL) -C $(RELEASE_DIR)

	@echo "Structuring directories..."
	@mkdir -p include lib/$(PLATFORM)-$(GOARCH)

	@cp -R $(RELEASE_DIR)/include/* include/
	@cp -R $(RELEASE_DIR)/lib/lib* lib/$(PLATFORM)-$(GOARCH)/

	$(call FIX_LIBS,$(PLATFORM)-$(GOARCH))

	@echo "Cleaning up temporary files..."
	@rm -rf $(RELEASE_DIR)
	@rm $(TARBALL)
	@echo "Done!"

clean:
	@rm -rf include lib
	@echo "Local directories cleaned."