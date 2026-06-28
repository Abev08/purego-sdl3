package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

// To build this example for WASM use build script and index.html file from "wasm" example.

var running = make(chan bool, 1) // Required with sdl.EnterAppMainCallbacks on WASM

func main() {
	fmt.Println("Running sdl.RunApp with app function directly in the mainFunction func. This doesn't work on WASM but works on Windows/Linux.")
	sdl.RunApp(0, nil, func(argc int32, argv **byte) int32 {
		appMain()
		return 0
	}, nil)
	fmt.Printf("Finished running the app with sdl.RunApp with app function directly in the mainFunction func.\n\n")

	//
	//
	//

	// Run the sdl.RunApp in goroutine so it won't block. Use channel to synchronize mainFunction call with main thread.
	fmt.Println("Running sdl.RunApp with channel sync. This works on WASM and Windows/Linux.")
	ch := make(chan int32, 1)
	go func() { sdl.RunApp(0, nil, func(argc int32, argv **byte) int32 { ch <- 1; return 0 }, nil) }()
	<-ch
	appMain()
	fmt.Printf("Finished running the app with sdl.RunApp with channel sync.\n\n")

	//
	//
	//

	// On WASM the appIterate function is synchronized with JS window.requestAnimationFrame which should match VSync refresh rate.
	// On Windows/Linux the appIterate function is synchronized with configured refresh rate - uncapped, VSync, VSync/2, etc.
	// sdl.EnterAppMainCallbacks() is the preferred way of running SDL apps on WASM.
	fmt.Println("Running sdl.EnterAppMainCallbacks this runs the application using callbacks. This works on WASM and Windows/Linux.")
	sdl.EnterAppMainCallbacks(0, nil, appInit, appIterate, appEvent, appQuit)
	<-running // Wait for SDL to finish, it's required on WASM
	fmt.Println("Finished running the app with sdl.EnterAppMainCallbacks.")
}

func appMain() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("WASM with callbacks", 1280, 720, sdl.WindowResizable, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	var event sdl.Event
	startTime := time.Now()
	redColor := float32(0.0)
	// Close the app after 5 seconds of running
	for time.Since(startTime) < (time.Second * 5) {
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventMouseButtonUp:
				e := event.Button()
				if e.Button&uint8(sdl.ButtonLMask) > 0 {
					fmt.Println("Mouse button left clicked")
				} else if e.Button&uint8(sdl.ButtonMMask) > 0 {
					fmt.Println("Mouse button middle clicked")
				}
			}
		}

		redColor += 0.01
		if redColor > 1 {
			redColor = 0
		}

		sdl.SetRenderDrawColorFloat(renderer, redColor, 0.2, 0.2, 1.0)
		sdl.RenderClear(renderer)
		sdl.RenderPresent(renderer)

		// Slow down the main loop, sdl.RunApp doesn't solve the problem that VSync doesn't work on WASM and it has to be somehow slowed down
		time.Sleep(time.Millisecond * 15)
	}
}

// AppState can be anything, SDL holds pointer to it assigned in init callback and passes it to other callbacks.
type AppState struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer

	iterationCount int32
	eventCount     int32
	redColor       float32
	startTime      time.Time
}

func appInit(appState *unsafe.Pointer, argc int32, argv **byte) sdl.AppResult {
	fmt.Println("SDL3 Init callback called")

	// Assign application state variable
	// Instead of using SDL's appState you can use global variable if you want
	app := AppState{}
	if appState != nil {
		*appState = unsafe.Pointer(&app)
	}

	if !sdl.CreateWindowAndRenderer("WASM with callbacks", 1280, 720, sdl.WindowResizable, &app.Window, &app.Renderer) {
		panic(sdl.GetError())
	}

	sdl.SetRenderVSync(app.Renderer, 1)

	app.startTime = time.Now()
	return sdl.AppContinue
}

func appIterate(appState unsafe.Pointer) sdl.AppResult {
	app := (*AppState)(appState) // Get application state

	if app.iterationCount == 0 {
		fmt.Println("SDL3 Iterate callback called")
	}
	app.iterationCount += 1

	app.redColor += 0.01
	if app.redColor > 1 {
		app.redColor = 0
	}

	sdl.SetRenderDrawColorFloat(app.Renderer, app.redColor, 0.2, 0.2, 1.0)
	sdl.RenderClear(app.Renderer)
	sdl.RenderPresent(app.Renderer)

	// Close the app after 5 seconds of running
	if time.Since(app.startTime) < (time.Second * 5) {
		return sdl.AppContinue
	}
	return sdl.AppSuccess
}

func appEvent(appState unsafe.Pointer, event *sdl.Event) sdl.AppResult {
	app := (*AppState)(appState) // Get application state

	if app.eventCount == 0 {
		fmt.Println("SDL3 Event callback called")
	}
	app.eventCount += 1

	switch event.Type() {
	case sdl.EventMouseButtonUp:
		e := event.Button()
		if e.Button&uint8(sdl.ButtonLMask) > 0 {
			fmt.Println("Mouse button left clicked")
		} else if e.Button&uint8(sdl.ButtonMMask) > 0 {
			fmt.Println("Mouse button middle clicked")
		}
	}

	return sdl.AppContinue
}

func appQuit(appState unsafe.Pointer, result int32) {
	app := (*AppState)(appState) // Get application state

	fmt.Println("SDL3 Quit callback called")

	fmt.Printf("AppState after running the app: %+v\nThe app was running for: %v\n", app, time.Since(app.startTime))

	sdl.DestroyRenderer(app.Renderer)
	sdl.DestroyWindow(app.Window)
	running <- true
}
