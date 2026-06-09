#!/bin/bash

# Copy the Go WASM support file
cp -n "$(go env GOROOT)/lib/wasm/wasm_exec.js" build/.

# This command lists libraries officially supported by Emscripten
# /usr/lib/emsdk/upstream/emscripten/emcc --show-ports | grep "sdl3"

# Compile SDL3 directly into JS/WASM
# This can be done once and the generated files can be reused
if false; then
    # SDL3 shipped with Emscripten can be used, it may be outdated
    # Comment out not required SDL flags to reduce generated .wasm file size
    emccSDLFlags=""
    emccSDLFlags=$emccSDLFlags" -sUSE_SDL=3"
    # emccSDLFlags=$emccSDLFlags" -sUSE_SDL_IMAGE=3" # Not yet supported?
    emccSDLFlags=$emccSDLFlags" -sUSE_SDL_TTF=3"
    touch nothing.c
    echo 'void TTF_RenderUTF8_Solid() {}' >> nothing.c # Stub function is needed for this one or it will throw an error during runtime
    # The command assumes that you failed adding emcc to path so the absolute path to emcc is used
    /usr/lib/emsdk/upstream/emscripten/emcc nothing.c -o build/sdl.js \
        $emccSDLFlags \
        -sMODULARIZE=1 \
        -sEXPORT_NAME="createSDLModule" \
        -sMAIN_MODULE=1 \
        -sEXPORT_ALL=1 \
        -O3 || {
        echo "emcc build failed"
        rm nothing.c
        exit 1
    }
    rm nothing.c
else
    # SDL3 can be manually build into .js + .wasm from source, build_sdl.sh helps with that
    echo "Compile SDL3 into .js + .wasm manually using build_sdl.sh"
fi

# Compile Go into WASM
GOOS=js GOARCH=wasm go build -o build/go.wasm main.go || {
    echo "Go build failed"
    exit 1
}

# Host the app
python3 -m http.server 8080 --directory ./build/
