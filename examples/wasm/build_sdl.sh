# This script builds SDL3 + SDL3_TTF + SDL3_Image from source into sdl.js + sdl.wasm files.
#
# The build script assumes following directory structure:
# - build_sdl.sh - this file
# - build/*      - directory where generated files are created, already existing files are overwritten
# - SDL/*        - main SDL3 library directory with required commit checked out
# - SDL_image/*  - SDL3_Image library directory with required commit checked out
# - SDL_ttf/*    - SDL3_TTF library directory with required commit checked out

dir=$PWD
mkdir build

## Building SDL3
rm -r SDL/build_wasm
mkdir SDL/build_wasm
cd SDL/build_wasm
# Configure build
/usr/lib/emsdk/upstream/emscripten/emcmake cmake .. -DCMAKE_BUILD_TYPE=Release -DSDL_STATIC=ON -DSDL_SHARED=OFF -DSDL_TESTS=OFF
# Build into libSDL3.a
/usr/lib/emsdk/upstream/emscripten/emmake make -j4
cd $dir

## Building SDL3_TTF
rm -r SDL_ttf/build_wasm
mkdir SDL_ttf/build_wasm
cd SDL_ttf/build_wasm
# Configure build
/usr/lib/emsdk/upstream/emscripten/emcmake cmake .. -DCMAKE_BUILD_TYPE=Release -DSDLTTF_VENDORED=ON -DSDLTTF_SAMPLES=OFF -DSDL3_DIR=$dir/SDL/build_wasm
# Build into libSDL3_ttf.a
/usr/lib/emsdk/upstream/emscripten/emmake make -j4
cd $dir

## Building SDL3_Image
rm -r SDL_image/build_wasm
mkdir SDL_image/build_wasm
cd SDL_image/build_wasm
# Configure build
/usr/lib/emsdk/upstream/emscripten/emcmake cmake .. -DCMAKE_BUILD_TYPE=Release -DSDL_IMAGE_VENDORED=ON -DSDL_IMAGE_SAMPLES=OFF -DSDL3_DIR=$dir/SDL/build_wasm
# Build into libSDL3_image.a
/usr/lib/emsdk/upstream/emscripten/emmake make -j4
cd $dir

## Combining created .a files into single .js + .wasm package
touch nothing.c
echo 'void TTF_RenderUTF8_Solid() {}' >> nothing.c # Stub function is needed for this one or it will throw an error during runtime
/usr/lib/emsdk/upstream/emscripten/emcc nothing.c -o $dir/build/sdl.js \
    -I$dir/SDL_ttf/include \
    -I$dir/SDL_image/include \
    -I$dir/SDL/include \
    $dir/SDL_ttf/build_wasm/libSDL3_ttf.a \
    $dir/SDL_image/build_wasm/libSDL3_image.a \
    $dir/SDL/build_wasm/libSDL3.a \
    -sMODULARIZE=1 \
    -sEXPORT_NAME="createSDLModule" \
    -sMAIN_MODULE=1 \
    -sEXPORT_ALL=1 \
    -O3 \
    --use-preload-plugins
rm nothing.c
