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

PATH_LVGL_LIB=$HOME/git-karboggy/lv_port_linux/build/lvgl/lib
PATH_LVGL_INCLUDE=$HOME/git-karboggy/lv_port_linux

COMMON_FLAGS="-mfloat-abi=softfp -mfpu=neon-vfpv4 -mcpu=cortex-a7 -O3 -ffast-math"
GCC="${TOOLCHAIN}clang"
GCC_FLAGS="-w -fPIC"

echo "Put lvgl libs in build directory (easier deployement)"
cp $PATH_LVGL_LIB/liblvgl* build/

echo "Building..."
mkdir -p build
${GCC} ${GCC_FLAGS} ${COMMON_FLAGS} \
    -o build/example-lvgl \
    examples/lvgl/lvgl-example.c \
    -llvgl \
    -I${PATH_LVGL_INCLUDE} \
    -L${PATH_LVGL_LIB} \
    -v

echo "Deploying..."
ssh root@$VECTOR_IP_ADDRESS "mkdir -p $TARGET_DIRECTORY"
rsync -av --exclude 'libjpeg-turbo' ./build/* root@$VECTOR_IP_ADDRESS:$TARGET_DIRECTORY/

echo -e "\nPlease run on Vector:\nLD_LIBRARY_PATH=$TARGET_DIRECTORY:/anki/lib $TARGET_DIRECTORY/example-lvgl\n"
ssh root@$VECTOR_IP_ADDRESS