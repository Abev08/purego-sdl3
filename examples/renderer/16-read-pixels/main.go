package main

import (
	"unsafe"

	"github.com/abev08/purego-sdl3/sdl"
)

const (
	WINDOW_WIDTH  int32 = 640
	WINDOW_HEIGHT int32 = 480
)

func main() {
	sdl.SetHint(sdl.HintRenderVSync, "1")

	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("examples/renderer/read-pixels", WINDOW_WIDTH, WINDOW_HEIGHT, 0, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	surface := sdl.LoadPNG(sdl.GetBasePath() + "../sample.png")
	if surface == nil {
		panic(sdl.GetError())
	}
	texture := sdl.CreateTextureFromSurface(renderer, surface)
	if texture == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyTexture(texture)
	sdl.DestroySurface(surface)

	var event sdl.Event
	var converted_texture *sdl.Texture

	running := true
	for running {
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit:
				running = false
			case sdl.EventKeyDown:
				if event.Key().Scancode == sdl.ScancodeEscape {
					running = false
				}
			}
		}

		// We'll have a texture rotate around over 2 seconds (2000 milliseconds). 360 degrees in a circle!
		now := sdl.GetTicks()
		rotation := (float32(now%2000) / 2000.0) * 360.0

		sdl.SetRenderDrawColor(renderer, 0, 0, 0, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		// Center this one, and draw it with some rotation so it spins!
		dst_rect := sdl.FRect{
			X: float32(WINDOW_WIDTH-texture.W) / 2,
			Y: float32(WINDOW_HEIGHT-texture.H) / 2,
			W: float32(texture.W),
			H: float32(texture.H),
		}
		// Rotate it around the center of the texture; you can rotate it from a different point, too!
		center := sdl.FPoint{X: float32(texture.W / 2), Y: float32(texture.H / 2)}
		sdl.RenderTextureRotated(renderer, texture, nil, &dst_rect, float64(rotation), &center, sdl.FlipNone)

		// Download the pixels of what has just been rendered. This has to wait for the GPU to finish rendering it and everything before it,
		// and then make an expensive copy from the GPU to system RAM!
		surface = sdl.RenderReadPixels(renderer, nil)
		if surface != nil {
			// This is also expensive, but easier: convert the pixels to a format we want.
			if (surface.Format != sdl.PixelFormatRGBA8888) && (surface.Format != sdl.PixelFormatBGRA8888) {
				converted := sdl.ConvertSurface(surface, sdl.PixelFormatRGBA8888)
				sdl.DestroySurface(surface)
				surface = converted
			}

			// Rebuild converted_texture if the dimensions have changed (window resized, etc).
			if (converted_texture == nil) || (surface.W != converted_texture.W) || (surface.H != converted_texture.H) {
				if converted_texture != nil {
					sdl.DestroyTexture(converted_texture)
				}
				converted_texture = sdl.CreateTexture(renderer, sdl.PixelFormatRGBA8888, sdl.TextureAccessStreaming, surface.W, surface.H)
				if converted_texture == nil {
					panic("Couldn't (re)create conversion texture: " + sdl.GetError())
				}
			}

			// Turn each pixel into either black or white. This is a lousy technique but it works here.
			// In real life, something like Floyd-Steinberg dithering might work
			// better: https://en.wikipedia.org/wiki/Floyd%E2%80%93Steinberg_dithering*/
			pixels := unsafe.Slice((*uint32)(surface.Pixels), surface.H*surface.W)
			for y := range surface.H {
				for x := range surface.W {
					p := unsafe.Slice((*uint8)(unsafe.Pointer(&pixels[(y*surface.W)+x])), unsafe.Sizeof(uint32(0)))
					average := (p[1] + p[2] + p[3]) / 3
					if average == 0 {
						p[0], p[3] = 0xFF, 0xFF // make pure black pixels red
						p[1], p[2] = 0, 0
					} else if average > 50 { // make everything else either black or white
						p[1], p[2], p[3] = 0xFF, 0xFF, 0xFF
					} else { // make everything else either black or white
						p[1], p[2], p[3] = 0, 0, 0
					}
				}
			}

			// Upload the processed pixels back into a texture.
			sdl.UpdateTexture(converted_texture, nil, surface.Pixels, surface.Pitch)
			sdl.DestroySurface(surface)

			// Draw the texture to the top-left of the screen.
			dst_rect.X, dst_rect.Y = 0, 0
			dst_rect.W = float32(WINDOW_WIDTH) / 4
			dst_rect.H = float32(WINDOW_HEIGHT) / 4
			sdl.RenderTexture(renderer, converted_texture, nil, &dst_rect)
		}

		sdl.RenderPresent(renderer)
	}
}
