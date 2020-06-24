#!/bin/sh
git submodule update --recursive --init
cd eip1962
cp eip196_header.h ../../
RUSTFLAGS="-C link-args=-lc" cargo build --release --no-default-features --features=eip_196_c_api

FILE_LINUX=target/release/libeth_pairings.so
FILE_OSX=target/release/libeth_pairings.dylib
FILE_WINDOWS=target/release/libeth_pairings.dll
if [ -f "$FILE_LINUX" ]; then
    cp $FILE_LINUX ../../libs/libeth_196.so
elif [ -f "$FILE_OSX"]; then
    cp $FILE_OSX ../../libs/libeth_196.dylib
else
    cp $FILE_WINDOWS ../../libs/libeth_196.dll
fi