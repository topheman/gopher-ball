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
all: help

help:
	@echo ""
	@echo "Usage:"
	@echo ""
	@echo "\tmake [command]"
	@echo ""
	@echo "The commands are:"
	@echo ""
	@echo "\tdarwin\t\t\tcompiles a bundle for MacOS in ./dist/gopher-ball-darwin-amd64.app (and zips it)"
	@echo "\twindows\t\tcompiles a bundle for Windows [not yet implemented]"
	@echo "\tlinux\t\t\t\tcompiles a bundle for Linux [not yet implemented]"
	@echo ""
	@echo "\tdarwin-dev\tsame as go build - but creates a file named gopher-ball.app (so that you can interact with it in the finder)"
	@echo ""
	@echo "\tclean\t\t\t\tcleans up ./dist folder (executed for each tasks above)"
	@echo "\tprepare\t\tcreates ./dist folder if doesn't exist (executed for each tasks above)"
	@echo ""

prepare:
	@echo "[WARNING] The development of this part is still in progress ..."
	@echo "[INFO] Creating ./dist folder"
	@mkdir -p dist

linux: clean prepare
	@echo "Skipping Linux ..."

windows: clean prepare
	@echo "Skipping Windows ..."

darwin: clean prepare
	go build -o ${BINARY}-darwin-${GOARCH} .

	mkdir -p ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/{MacOS,Frameworks,Resources}
	mv ${BINARY}-darwin-${GOARCH} ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/MacOS
	cp -R ${BUILD_DIR}/darwin/Contents ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app

	cp /usr/local/opt/sdl2_image/lib/libSDL2_image-2.0.0.dylib ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks
	cp /usr/local/opt/sdl2_ttf/lib/libSDL2_ttf-2.0.0.dylib ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks
	cp /usr/local/opt/sdl2/lib/libSDL2-2.0.0.dylib ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks
	cp /usr/local/opt/libpng/lib/libpng16.16.dylib ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks
	cp /usr/local/opt/libtiff/lib/libtiff.5.dylib ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks
	cp /usr/local/opt/webp/lib/libwebp.7.dylib ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks
	cp /usr/local/opt/jpeg/lib/libjpeg.9.dylib ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks
	cp /usr/local/opt/freetype/lib/libfreetype.6.dylib ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks

	cd ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/MacOS; \
	install_name_tool -change /usr/local/opt/sdl2_image/lib/libSDL2_image-2.0.0.dylib @executable_path/../Frameworks/libSDL2_image-2.0.0.dylib ${BINARY}-darwin-${GOARCH}; \
	install_name_tool -change /usr/local/opt/sdl2_ttf/lib/libSDL2_ttf-2.0.0.dylib @executable_path/../Frameworks/libSDL2_ttf-2.0.0.dylib ${BINARY}-darwin-${GOARCH}; \
	install_name_tool -change /usr/local/opt/sdl2/lib/libSDL2-2.0.0.dylib @executable_path/../Frameworks/libSDL2-2.0.0.dylib ${BINARY}-darwin-${GOARCH}; \
	cd - >/dev/null

	cd ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Frameworks; \
	chmod +w ./libSDL2_image-2.0.0.dylib libSDL2_ttf-2.0.0.dylib; \
	install_name_tool -change /usr/local/opt/libpng/lib/libpng16.16.dylib @executable_path/libpng16.16.dylib libSDL2_image-2.0.0.dylib; \
	install_name_tool -change /usr/local/opt/libtiff/lib/libtiff.5.dylib @executable_path/libtiff.5.dylib libSDL2_image-2.0.0.dylib; \
	install_name_tool -change /usr/local/opt/webp/lib/libwebp.7.dylib @executable_path/libwebp.7.dylib libSDL2_image-2.0.0.dylib; \
	install_name_tool -change /usr/local/opt/jpeg/lib/libjpeg.9.dylib @executable_path/libjpeg.9.dylib libSDL2_image-2.0.0.dylib; \
	install_name_tool -change /usr/local/opt/freetype/lib/libfreetype.6.dylib @executable_path/libfreetype.6.dylib libSDL2_ttf-2.0.0.dylib; \
	chmod -w libSDL2_image-2.0.0.dylib libSDL2_ttf-2.0.0.dylib; \
	cd - >/dev/null

	cp -R ./${ASSETS_DIRNAME} ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Resources/${ASSETS_DIRNAME}
	rm -rf ${DIST_DIR}/${BINARY}-darwin-${GOARCH}.app/Contents/Resources/${ASSETS_DIRNAME}/originals

	cd ./${DIST_DIRNAME}; \
	zip -r ${BINARY}-darwin-${GOARCH}.app.zip ${BINARY}-darwin-${GOARCH}.app; \
	cd - >/dev/null

darwin-dev:
	go build -o ${BINARY}.app

clean:
	@echo "[INFO] Cleaning ./dist folder"
	-rm -rf ${DIST_DIRNAME}/*