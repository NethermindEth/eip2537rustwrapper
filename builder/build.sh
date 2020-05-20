#!/bin/sh
git submodule update —recursive —init
cd eip1962
cp eip2537_header.h ../../
RUSTFLAGS="-C link-args=-lc" cargo build --release --no-default-features --features=eip_2357_c_api
cp target/release/libeth_pairings.a ../../libs
