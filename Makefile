# based on templates:
# * https://gist.github.com/turtlemonvh/38bd3d73e61769767c35931d8c70ccb4
# * https://vincent.bernat.im/fr/blog/2017-makefile-pour-golang

BINARY = gopher-ball
GOARCH = amd64

VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

CURRENT_DIR=$(shell pwd)
BUILD_DIR=${CURRENT_DIR}
ASSETS_DIRNAME=assets

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-s -X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# Build the project
all: clean prepare linux darwin windows

prepare:
#	mkdir -p build/dist

linux:
	echo "Skipping Linux ..."

windows:
	echo "Skipping Windows ..."

darwin:
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-darwin-${GOARCH}.app .

	mkdir -p ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}
	mv ${BINARY}-darwin-${GOARCH}.app ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}/

	cp -R ./${ASSETS_DIRNAME} ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}/${ASSETS_DIRNAME}
	rm -rf ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}/${ASSETS_DIRNAME}/originals

	mkdir -p ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}/lib
	cp /usr/local/opt/sdl2_image/lib/libSDL2_image-2.0.0.dylib ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}/lib/
	cp /usr/local/opt/sdl2/lib/libSDL2-2.0.0.dylib ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}/lib/
	cp /usr/local/opt/sdl2_ttf/lib/libSDL2_ttf-2.0.0.dylib ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}/lib/
	cp /usr/lib/libSystem.B.dylib ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}/lib/

	zip -r ${BINARY}-darwin-${GOARCH}.zip ${BINARY}-darwin-${GOARCH}

clean:
	-rm -rf ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}
	-rm -rf ${BUILD_DIR}/${BINARY}-darwin-${GOARCH}.zip

.PHONY: clean prepare linux darwin windows