SHELL:=/bin/sh
.PHONY: build clean fmt
export GO111MODULE=on

#path relate
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))
RELEASE_DIR := $(MKFILE_DIR)bin

RELEASE?=0.0.1

#build flags
GO_LD_FLAGS="-s -w"

#targets
TARGET_ALPHA_LOCAL=${RELEASE_DIR}/alpha

build: build_alpha_local

build_alpha_local:
				@echo "build alpha local ${MKFILE_DIR}"
				cd ${MKFILE_DIR} && \
				CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -v -trimpath -ldflags ${GO_LD_FLAGS} \
				-o ${TARGET_ALPHA_LOCAL} ${MKFILE_DIR}cmd/local

clean:
				rm -rf ${RELEASE_DIR}

#build server docker image

fmt:
				cd ${MKFILE_DIR} && go fmt ./...