#  Makefile for Go

GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_TEST=$(GO_CMD) test -v
GO_CLEAN=$(GO_CMD) clean
GO_VET=$(GO_CMD) vet
GO_FMT=go fmt
GO_LINT=golint

GOOS ?= $(shell go env GOOS)
ifeq ($(GOOS), windows)
	IS_EXE := .exe
endif

# Packages
PACKAGE_LIST := $(GOPATH)/src/S3FileWriter
NON_VENDOR := $(shell go list ./... | grep -v /vendor)
DIST_DIR := ./build
SOURCE_DIR := $(PACKAGE_LIST)/app
NAME:=build/S3FileWriter
BUILD_FILE:=$(GOPATH)/src/common

info:
	@if [ -d "$(PACKAGE_LIST)" ]; then \
		echo "Reading project info  $$p ..."; \
		n := $(BUILD_FILE)/readYaml.sh $(PACKAGE_LIST)/build.yml ; \
		echo "stored value.."$(n); \
	fi


.PHONY: all build compile test clean fmt vet lint

all: build
compile: clean fmt vet lint

build: vet
	@if [ -d "$(PACKAGE_LIST)" ]; then \
		echo "Build $$p ..."; \
		$(GO_BUILD) -o $(NAME) $(SOURCE_DIR)/*.go || exit 1; \
	fi

test:
	@if [ -d "$(PACKAGE_LIST)" ]; then \
		echo "Unit Testing $$p ..."; \
		$(GO_TEST) $(NON_VENDOR) ; \
	fi

clean:
	@if [ -d "$(DIST_DIR)" ]; then \
		echo "Cleaning $$p ..."; \
		rm -rf $(DIST_DIR)$*; \
	fi

fmt:
	@if [ -d "$(PACKAGE_LIST)" ]; then \
		echo "Formatting ..."; \
		$(GO_FMT) $(NON_VENDOR); \
	fi
vet:
	@if [ -d "$(PACKAGE_LIST)" ]; then \
		echo "Vet $$p ..."; \
		$(GO_VET) $(NON_VENDOR); \
	fi

lint:
	@if [ -d "$(PACKAGE_LIST)" ]; then \
		echo "Lint $$p ... "; \
		go list ./... | grep -v /vendor | xargs -L1 golint; \
	fi
# vim: set noexpandtab shiftwidth=8 softtabstop=0:
