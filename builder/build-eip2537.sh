#!/bin/sh
git submodule update --recursive --init
cd eip1962
cp eip2537_header.h ../../
RUSTFLAGS="-C link-args=-lc" cargo build --release --no-default-features --features=eip_2357_c_api

FILE_LINUX=target/release/libeth_pairings.so
FILE_OSX=target/release/libeth_pairings.dylib
FILE_WINDOWS=target/release/libeth_pairings.dll
if [ -f "$FILE_LINUX" ]; then
    cp $FILE_LINUX ../../libs/libeth_2537.so
elif [ -f "$FILE_OSX"]; then
    cp $FILE_OSX ../../libs/libeth_2537.dylib
else
    cp $FILE_WINDOWS ../../libs/eth_2537.dll
fi