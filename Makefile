# based on templates:
# * https://gist.github.com/turtlemonvh/38bd3d73e61769767c35931d8c70ccb4
# * https://vincent.bernat.im/fr/blog/2017-makefile-pour-golang

BINARY = gopher-ball
GOARCH = amd64

VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

CURRENT_DIR=$(shell pwd)
BUILD_DIRNAME = dist
BUILD_DIR=${CURRENT_DIR}/${BUILD_DIRNAME}
ASSETS_DIRNAME=assets

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-s -X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# Build the project
all: clean prepare linux darwin windows

prepare:
	mkdir -p dist

linux:
	echo "Skipping Linux ..."

windows:
	echo "Skipping Windows ..."

darwin:
	CGO_ENABLED=1 GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-darwin-${GOARCH}.app .

	mkdir -p ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}
	mv ${BINARY}-darwin-${GOARCH}.app ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}/

	cp -R ./${ASSETS_DIRNAME} ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}/${ASSETS_DIRNAME}
	rm -rf ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}/${ASSETS_DIRNAME}/originals

	cd ./${BUILD_DIRNAME}; \
	zip -r ${BINARY}-darwin-${GOARCH}.zip ${BINARY}-darwin-${GOARCH}; \
	cd - >/dev/null

clean:
	-rm -rf ${BUILD_DIR}/*

.PHONY: clean prepare linux darwin windows