package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

func main() {
	fmt.Println("Running on:", runtime.GOOS, runtime.GOARCH)
	runningOnWasm := runtime.GOOS == "js" && runtime.GOARCH == "wasm"
	fmt.Println("Compiled with Go:", runtime.Version())

	defer fmt.Println("SDL3 error:", sdl.GetError())

	major, minor, patch := sdl.GetVersion()
	fmt.Printf("SDL3 version: %d.%d.%d\n", major, minor, patch)
	fmt.Println("SDL3 revision:", sdl.GetRevision())

	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic("SDL3 initialization failed, err: " + sdl.GetError())
	}

	fmt.Print("Available renderers: ")
	for i := int32(0); i < sdl.GetNumRenderDrivers(); i++ {
		fmt.Print(sdl.GetRenderDriver(i), ", ")
	}
	fmt.Println()

	var window *sdl.Window
	var renderer *sdl.Renderer
	// window = sdl.CreateWindow("Hello WASM", 1280, 720, sdl.WindowResizable)
	// renderer = sdl.CreateRenderer(window, "")
	if !sdl.CreateWindowAndRenderer("Hello WASM", 1280, 720, sdl.WindowResizable, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	fmt.Println("Window title:", sdl.GetWindowTitle(window))
	fmt.Println("Renderer name:", sdl.GetRendererName(renderer))

	// Setting VSync is not supported on WASM?
	if !runningOnWasm && !sdl.SetRenderVSync(renderer, 1) {
		fmt.Println(sdl.GetError())
	}

	perfFreq, frameStart, frameCount := float32(sdl.GetPerformanceFrequency()), sdl.GetPerformanceCounter(), 0
	var fps, frameTime float32

	mousePos := sdl.FPoint{}
	running, mouseButtonDown := true, false
	event := sdl.Event{}
	for running {
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit:
				running = false // Maybe this shouldn't be allowed on Wasm?

			case sdl.EventKeyUp:
				e := event.Key()
				switch e.Key {
				case sdl.KeycodeEscape:
					running = false // Maybe this shouldn't be allowed on Wasm?
				}

			case sdl.EventMouseMotion:
				e := event.Motion()
				mousePos.X, mousePos.Y = e.X, e.Y

			case sdl.EventMouseButtonDown:
				e := event.Button()
				switch sdl.MouseButtonFlags(e.Button) {
				case sdl.ButtonLeft:
					mouseButtonDown = true
				}

			case sdl.EventMouseButtonUp:
				e := event.Button()
				switch sdl.MouseButtonFlags(e.Button) {
				case sdl.ButtonLeft:
					mouseButtonDown = false
				}
			}
		}

		sdl.SetRenderDrawColor(renderer, 0, 0, 255, 255)
		sdl.RenderClear(renderer)

		// Line
		sdl.SetRenderDrawColor(renderer, 255, 0, 0, 255)
		sdl.RenderLine(renderer, 100, 100, 300, 200)

		// Rect
		sdl.RenderRect(renderer, &sdl.FRect{X: 350, Y: 100, W: 100, H: 100})

		// Left rect filled
		sdl.SetRenderDrawColor(renderer, 0, 255, 255, 255)
		sdl.RenderFillRect(renderer, &sdl.FRect{X: 100, Y: 300, W: 200, H: 200})

		// Right rect filled, clickable
		r := sdl.FRect{X: 500, Y: 50, W: 300, H: 150}
		if sdl.PointInRectFloat(mousePos, r) {
			if mouseButtonDown {
				sdl.SetRenderDrawColorFloat(renderer, 0.2, 0.2, 0.2, 1)
			} else {
				sdl.SetRenderDrawColorFloat(renderer, 0.8, 0.8, 0.8, 1)
			}
		} else {
			sdl.SetRenderDrawColorFloat(renderer, 0, 1, 0, 1)
		}
		sdl.RenderFillRect(renderer, &r)

		// Connected lines
		sdl.SetRenderDrawColorFloat(renderer, 1, 1, 0, 1)
		sdl.RenderLines(renderer, []sdl.FPoint{{X: 400, Y: 250}, {X: 450, Y: 350}, {X: 500, Y: 250}, {X: 550, Y: 350}, {X: 600, Y: 250}})

		// Multiple rectangles
		sdl.SetRenderDrawColorFloat(renderer, 1, 0, 1, 1)
		sdl.RenderRects(renderer, []sdl.FRect{
			{X: 350, Y: 400, W: 30, H: 30}, {X: 390, Y: 400, W: 30, H: 30}, {X: 430, Y: 400, W: 30, H: 30}, {X: 470, Y: 400, W: 30, H: 30},
			{X: 350, Y: 440, W: 30, H: 30}, {X: 390, Y: 440, W: 30, H: 30}, {X: 430, Y: 440, W: 30, H: 30}, {X: 470, Y: 440, W: 30, H: 30},
			{X: 350, Y: 480, W: 30, H: 30}, {X: 390, Y: 480, W: 30, H: 30}, {X: 430, Y: 480, W: 30, H: 30}, {X: 470, Y: 480, W: 30, H: 30},
		})

		// Multiple filled rectangles
		sdl.SetRenderDrawColorFloat(renderer, 1, 1, 0, 1)
		sdl.RenderFillRects(renderer, []sdl.FRect{
			{X: 550, Y: 400, W: 30, H: 30}, {X: 590, Y: 400, W: 30, H: 30}, {X: 630, Y: 400, W: 30, H: 30}, {X: 670, Y: 400, W: 30, H: 30},
			{X: 550, Y: 440, W: 30, H: 30}, {X: 590, Y: 440, W: 30, H: 30}, {X: 630, Y: 440, W: 30, H: 30}, {X: 670, Y: 440, W: 30, H: 30},
			{X: 550, Y: 480, W: 30, H: 30}, {X: 590, Y: 480, W: 30, H: 30}, {X: 630, Y: 480, W: 30, H: 30}, {X: 670, Y: 480, W: 30, H: 30},
		})

		// Simple debug text
		sdl.SetRenderDrawColorFloat(renderer, 1, 1, 1, 1)
		sdl.RenderDebugText(renderer, 10, 10, "Hello")

		// Mouse position text
		sdl.RenderDebugTextFormat(renderer, 10, 30, "Mouse position: {x: %.0f, y: %.0f}", mousePos.X, mousePos.Y)

		// FPS counter
		frameEnd := sdl.GetPerformanceCounter()
		frameCount++
		dt := float32(frameEnd-frameStart) / perfFreq
		if dt >= 1 {
			fps = float32(frameCount) / dt
			frameTime = (dt * 1000) / fps
			frameCount, frameStart = 0, frameEnd
		}
		sdl.RenderDebugTextFormat(renderer, 10, 50, "FPS: %.0f, frame time: %.4f ms", fps, frameTime)

		sdl.RenderPresent(renderer)

		if runningOnWasm {
			time.Sleep(time.Millisecond * 16) // On Wasm the main loop has to be somehow slown down?
		}
	}
}
