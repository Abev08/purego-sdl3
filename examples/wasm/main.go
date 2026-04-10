package main

import (
	_ "embed"
	"fmt"
	"math/rand/v2"
	"runtime"
	"time"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

// For now the image gets embedded to allow texture creation on WASM, maybe there is other way?
//go:embed gopher-happy.bmp
var gopherHappy []byte

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
	// Window and renderer in one call
	if !sdl.CreateWindowAndRenderer("Hello WASM", 1280, 720, sdl.WindowResizable, &window, &renderer) {
		panic(sdl.GetError())
	}

	// Separate calls for window and renderer
	// window = sdl.CreateWindow("Hello WASM", 1280, 720, sdl.WindowResizable)
	// renderer = sdl.CreateRenderer(window, "")

	// Renderer with properties
	// rendererProperties := sdl.CreateProperties()
	// if rendererProperties == 0 {
	// 	panic(sdl.GetError())
	// }
	// if !sdl.SetPointerProperty(rendererProperties, sdl.PropRendererCreateWindowPointer, unsafe.Pointer(window)) {
	// 	panic(sdl.GetError())
	// }
	// if !sdl.SetNumberProperty(rendererProperties, sdl.PropRendererCreatePresentVSyncNumber, 1) {
	// 	panic(sdl.GetError())
	// }
	// renderer = sdl.CreateRendererWithProperties(rendererProperties)

	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	fmt.Println("Window title:", sdl.GetWindowTitle(window))
	fmt.Println("Renderer name:", sdl.GetRendererName(renderer))

	// Setting VSync is not supported on WASM?
	if !runningOnWasm && !sdl.SetRenderVSync(renderer, 1) {
		fmt.Println(sdl.GetError())
	}

	// Loading texture
	gopherStream := sdl.IOFromConstMem(gopherHappy)
	if gopherStream == nil {
		panic(sdl.GetError())
	}
	gopherSurface := sdl.LoadBMPIO(gopherStream, true)
	if gopherSurface == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroySurface(gopherSurface)
	gopherTexture := sdl.CreateTextureFromSurface(renderer, gopherSurface)
	if gopherTexture == nil {
		panic(gopherTexture)
	}
	defer sdl.DestroyTexture(gopherTexture)
	sdl.SetTextureScaleMode(gopherTexture, sdl.ScaleModeNearest)

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

		// Points
		sdl.SetRenderDrawColor(renderer, 255, 255, 255, 255)
		sdl.RenderPoint(renderer, 150, 150)
		points := make([]sdl.FPoint, 1000)
		for i := range points {
			points[i].X, points[i].Y = 100+100*rand.Float32(), 175+100*rand.Float32()
		}
		sdl.SetRenderDrawColor(renderer, 255, 0, 0, 255)
		sdl.RenderPoints(renderer, points)

		// Texture
		sdl.RenderTexture(renderer, gopherTexture, nil, &sdl.FRect{X: 600, Y: 275, W: 96, H: 96})
		sdl.RenderTextureRotated(renderer, gopherTexture, &sdl.FRect{X: 0, Y: 0, W: 32, H: 16}, &sdl.FRect{X: 700, Y: 250, W: 64, H: 32}, -45, nil, sdl.FlipNone)
		sdl.RenderTexture9Grid(renderer, gopherTexture, nil, 8, 8, 8, 8, 1, &sdl.FRect{X: 620, Y: 210, W: 64, H: 64})
		sdl.RenderTextureTiled(renderer, gopherTexture, nil, 1, &sdl.FRect{X: 720, Y: 320, W: 96, H: 160})

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
