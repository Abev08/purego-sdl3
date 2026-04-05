package main

import (
	"fmt"
	"time"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

func main() {
	defer fmt.Println("error:", sdl.GetError())

	major, minor, patch := sdl.GetVersion()
	fmt.Printf("Version: %d.%d.%d\n", major, minor, patch)
	fmt.Println("Revision:", sdl.GetRevision())

	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic("SDL3 initialization failed")
	}

	window := sdl.CreateWindow("Hello WASM", 1280, 720, sdl.WindowResizable)
	defer sdl.DestroyWindow(window)

	fmt.Println("Window title:", sdl.GetWindowTitle(window))

	renderer := sdl.CreateRenderer(window, "")
	defer sdl.DestroyRenderer(renderer)

	mousePos := sdl.FPoint{}
	running, mouseButtonDown := true, false
	for running {
		event := sdl.Event{}
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit:
				running = false

			case sdl.EventMouseMotion:
				e := event.Motion()
				mousePos.X, mousePos.Y = e.X, e.Y

			case sdl.EventMouseButtonDown:
				e := event.Button()
				if e.Button == uint8(sdl.ButtonLeft) {
					mouseButtonDown = true
				}

			case sdl.EventMouseButtonUp:
				e := event.Button()
				if e.Button == uint8(sdl.ButtonLeft) {
					mouseButtonDown = false
				}
			}
		}

		sdl.SetRenderDrawColor(renderer, 0, 0, 255, 255)
		sdl.RenderClear(renderer)

		sdl.SetRenderDrawColor(renderer, 255, 0, 0, 255)
		sdl.RenderLine(renderer, 100, 100, 300, 200)

		r := sdl.FRect{X: 100, Y: 300, W: 200, H: 200}
		if sdl.PointInRectFloat(mousePos, r) {
			sdl.SetRenderDrawColor(renderer, 0, 255, 255, 255)
		} else {
			sdl.SetRenderDrawColor(renderer, 255, 255, 0, 255)
		}
		sdl.RenderFillRect(renderer, &r)

		r = sdl.FRect{X: 500, Y: 50, W: 300, H: 150}
		if sdl.PointInRectFloat(mousePos, r) {
			if mouseButtonDown {
				sdl.SetRenderDrawColor(renderer, 0, 0, 0, 255)
			} else {
				sdl.SetRenderDrawColor(renderer, 255, 255, 255, 255)
			}
		} else {
			sdl.SetRenderDrawColor(renderer, 0, 255, 0, 255)
		}
		sdl.RenderFillRect(renderer, &r)

		sdl.SetRenderDrawColor(renderer, 255, 255, 255, 255)
		sdl.RenderDebugText(renderer, 10, 10, fmt.Sprintf("Mouse position: {x: %.0f, y: %.0f}", mousePos.X, mousePos.Y))

		sdl.RenderPresent(renderer)

		time.Sleep(time.Millisecond * 16)
	}
}
