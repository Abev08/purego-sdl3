#!/bin/bash

# Copy the Go WASM support file
cp -n "$(go env GOROOT)/lib/wasm/wasm_exec.js" .

# Compile SDL3 directly into JS/WASM
# The command assumes that you failed adding emcc to path so the absolute path to emcc is used
touch nothing.c
/usr/lib/emsdk/upstream/emscripten/emcc nothing.c -o sdl.js \
    -s USE_SDL=3 \
    -s MODULARIZE=1 \
    -s EXPORT_NAME='createSDLModule' \
    -s MAIN_MODULE=1 \
    -s EXPORT_ALL=1 \
    -O3 || {
    echo "emcc build failed";
    exit 1
}
rm nothing.c

# Compile Go into WASM
GOOS=js GOARCH=wasm go build -o go.wasm main.go || {
    echo "Go build failed";
    exit 1
}

# Host the app
python3 -m http.server 8080

