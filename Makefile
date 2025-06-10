APP_NAME := dag-engine
BIN_DIR := bin
DOCKER_IMAGE := dag-engine
VERSION := $(shell git describe --tags --always --dirty)
GO_FILES := $(shell find . -type f -name "*.go" -not -path "./vendor/*")
GO_PKGS := $(shell go list ./.. | grep -v /vendor/)
