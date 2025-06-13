#!/bin/bash
set -e

# $HOME/git-karboggy/toolchains/vicos-sdk_5.2.1-r06_amd64-linux
VICOS_SDK_DIRECTORY=${VICOS_SDK_DIRECTORY:=$HOME/git-karboggy/toolchains/vicos-sdk_5.2.1-r06_amd64-linux}
TOOLCHAIN=${TOOLCHAIN:=$VICOS_SDK_DIRECTORY/prebuilt/bin/arm-oe-linux-gnueabi-}

# clean all
rm -rf build
rm -rf libjpeg-turbo/build

# build libs
TOOLCHAIN=$TOOLCHAIN make vector-gobot
TOOLCHAIN=$TOOLCHAIN make libjpeg-turbo
TOOLCHAIN=$TOOLCHAIN make jpeg_interface

# auto build example mjpeg
# VICOS_SDK_DIRECTORY=$VICOS_SDK_DIRECTORY TOOLCHAIN=$TOOLCHAIN ./build-example-mjpeg.sh