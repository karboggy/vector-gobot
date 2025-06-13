#!/bin/bash
set -e

#----------------------------------------------
# Tips
#----------------------------------------------
# before using this script, SSH vector and do:
# systemctl stop anki-robot.target
# mount -o rw,remount,exec /data
#----------------------------------------------

VICOS_SDK_DIRECTORY=${VICOS_SDK_DIRECTORY:=$HOME/git-karboggy/toolchains/vicos-sdk_5.2.1-r06_amd64-linux}
TOOLCHAIN=${TOOLCHAIN:=$VICOS_SDK_DIRECTORY/prebuilt/bin/arm-oe-linux-gnueabi-}
VECTOR_IP_ADDRESS=192.168.1.24
TARGET_DIRECTORY=/data/karboggy

# screen
EXAMPLE_CATEGORY=screen

# screen-helloworld, screen-battery, screen-drawing-gg
EXAMPLE_NAME=screen-drawing-gg

echo "Building..."
CC="${TOOLCHAIN}clang" \
CGO_LDFLAGS="-Lbuild" \
GOARM=7 \
GOARCH=arm \
CGO_ENABLED=1 \
go build -o build/example-$EXAMPLE_NAME examples/$EXAMPLE_CATEGORY/$EXAMPLE_NAME.go

echo "Copy res in build..."
cp -r ./res ./build/res

echo "Deploying..."
ssh root@$VECTOR_IP_ADDRESS "mkdir -p $TARGET_DIRECTORY"
rsync -av --exclude 'libjpeg-turbo' ./build/* root@$VECTOR_IP_ADDRESS:$TARGET_DIRECTORY/

echo -e "Please run on Vector:\nLD_LIBRARY_PATH=$TARGET_DIRECTORY:/anki/lib $TARGET_DIRECTORY/example-$EXAMPLE_NAME"
ssh root@$VECTOR_IP_ADDRESS