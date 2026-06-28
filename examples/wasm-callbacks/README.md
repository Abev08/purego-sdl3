This example demonstrates how to use SDL callback system to run the app.\
Using [SDL_EnterAppMainCallbacks]() is the preferred way of running SDL apps on WASM.\
The example opens a window and clears it with different color each frame. The window is opened 3 times, each time using different approach.\
To build the example for WASM use build script and index.html file from "wasm" example. The example can be build for Windows/Linux without changing anything.