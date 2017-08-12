# based on templates:
# * https://gist.github.com/turtlemonvh/38bd3d73e61769767c35931d8c70ccb4
# * https://vincent.bernat.im/fr/blog/2017-makefile-pour-golang

BINARY = gopher-ball
GOARCH = amd64

VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

CURRENT_DIR=$(shell pwd)
DIST_DIRNAME = dist
DIST_DIR=${CURRENT_DIR}/${DIST_DIRNAME}
BUILD_DIRNAME = build
BUILD_DIR=${CURRENT_DIR}/${BUILD_DIRNAME}
ASSETS_DIRNAME=assets

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-s -X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# Build the project
all: clean prepare linux darwin windows

prepare:
	echo "[WARNING] The development of this part is still in progress ..."
	mkdir -p dist

linux:
	echo "Skipping Linux ..."

windows:
	echo "Skipping Windows ..."

darwin:
	go build -o ${BINARY}-darwin-${GOARCH} .

	mkdir -p ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/{MacOS,Frameworks,Resources}
	mv ${BINARY}-darwin-${GOARCH} ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/MacOS
	cp ${BUILD_DIR}/darwin/Contents/Info.plist ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents

	cp /usr/local/opt/sdl2_image/lib/libSDL2_image-2.0.0.dylib ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks
	cp /usr/local/opt/sdl2_ttf/lib/libSDL2_ttf-2.0.0.dylib ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks
	cp /usr/local/opt/sdl2/lib/libSDL2-2.0.0.dylib ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks

	cd ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/MacOS; \
	install_name_tool -change /usr/local/opt/sdl2_image/lib/libSDL2_image-2.0.0.dylib @executable_path/../Frameworks/libSDL2_image-2.0.0.dylib ${BINARY}-darwin-${GOARCH}; \
	install_name_tool -change /usr/local/opt/sdl2_ttf/lib/libSDL2_ttf-2.0.0.dylib @executable_path/../Frameworks/libSDL2_ttf-2.0.0.dylib ${BINARY}-darwin-${GOARCH}; \
	install_name_tool -change /usr/local/opt/sdl2/lib/libSDL2-2.0.0.dylib @executable_path/../Frameworks/libSDL2-2.0.0.dylib ${BINARY}-darwin-${GOARCH}; \
	cd - >/dev/null

	cp -R ./${ASSETS_DIRNAME} ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Resources/${ASSETS_DIRNAME}
	rm -rf ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/MacOS/${ASSETS_DIRNAME}/originals

darwin-dev:
	go build -o ${BINARY}.app

clean:
	-rm -rf ${DIST_DIR}/*

.PHONY: clean prepare linux darwin windows