# Development Dependencies Setup

This document explains how to set up the required dependencies for local development of the backend service.

## WebP Library Dependency

The backend now uses the `github.com/kolesa-team/go-webp` library for converting images to WebP format. This library requires CGO to be enabled and the WebP C library to be installed on your system.

### macOS

Install the WebP library using Homebrew:

```bash
brew install webp
```

### Ubuntu/Debian

Install the WebP development libraries:

```bash
sudo apt-get update
sudo apt-get install libwebp-dev
```

### CentOS/RHEL/Fedora

Install the WebP development libraries:

```bash
# For CentOS/RHEL
sudo yum install libwebp-devel

# For Fedora
sudo dnf install libwebp-devel
```

## Building the Backend

When building the backend locally, make sure CGO is enabled:

```bash
cd services/backend
CGO_ENABLED=1 go build -o myapplication src/main.go
```

Note: The Docker build process already handles this automatically by installing the required libraries and enabling CGO.
