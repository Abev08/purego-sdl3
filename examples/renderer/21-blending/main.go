package main

import (
	"github.com/abev08/purego-sdl3/sdl"
)

const (
	WINDOW_WIDTH  int32 = 640
	WINDOW_HEIGHT int32 = 480

	// UI constants
	ROWS         int32 = 2
	COLS         int32 = 3
	GRID_SIZE          = (WINDOW_WIDTH - 1) / 18
	PANEL_SIZE         = GRID_SIZE * 4
	COL_OFFSET         = GRID_SIZE * COLS
	ROW_OFFSET         = (WINDOW_HEIGHT - ROWS*PANEL_SIZE) / 4
	RECT_SIZE    int32 = 50
	RED_OFFSET         = GRID_SIZE
	GREEN_OFFSET       = RECT_SIZE/3 + GRID_SIZE
	BLUE_OFFSET        = RECT_SIZE*2/3 + GRID_SIZE
)

func main() {
	sdl.SetHint(sdl.HintRenderVSync, "1")

	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo) {
		panic(sdl.GetError())
	}

	var window *sdl.Window
	var renderer *sdl.Renderer
	if !sdl.CreateWindowAndRenderer("examples/renderer/blending", WINDOW_WIDTH, WINDOW_HEIGHT, 0, &window, &renderer) {
		panic(sdl.GetError())
	}
	defer sdl.DestroyWindow(window)
	defer sdl.DestroyRenderer(renderer)

	panels := make([]sdl.FRect, ROWS*COLS)
	for row := range ROWS {
		for col := range COLS {
			panels[col+row*COLS] = sdl.FRect{
				X: float32(col*PANEL_SIZE + col*COL_OFFSET),
				Y: float32(row*PANEL_SIZE + (row+1)*ROW_OFFSET),
				W: float32(PANEL_SIZE),
				H: float32(PANEL_SIZE),
			}
		}
	}

	blend_modes := []sdl.BlendMode{
		// The default no blending: dstRGB := srcRGB,
		//                          dstA  := srcA
		sdl.BlendModeNone,
		// Alpha blending: dstRGB := srcA * srcRGB + (1 - srcA) * dstRGB
		//                 dstA   := srcA          + (1 - srcA) * dstA
		sdl.BlendModeBlend,
		// Additive blending: dstRGB := srcRGB + dstRGB
		//                    dstA   := srcA   + dstA
		sdl.BlendModeAdd,
		// Modulate blending: dstRGB := srcRGB * dstRGB
		//                    dstA   := dstA
		sdl.BlendModeMod,
		// Multiply blending: dstRGB := srcRGB * dstRGB + (1 - srcA) * dstRGB
		//                    dstA   := dstA
		sdl.BlendModeMul,
		// Our custom blending 'Screen Blending': dstRGB := 1 - (1 - dstRGB) * (1 - srcRGB)
		//                                        dstA   := dstA
		// Create 'screen blend' mode
		sdl.ComposeCustomBlendMode(
			sdl.BlendFactorOneMinusDstColor, // srcRGB factor    := (1 - dstRGB)
			sdl.BlendFactorOne,              // dstRGB factor    := 1
			sdl.BlendOperationAdd,           // RGB    operation := +
			sdl.BlendFactorZero,             // srcA   factor    := 0
			sdl.BlendFactorOne,              // dstA   factor    := dstA
			sdl.BlendOperationAdd,           // A      operation := +
		),
	}
	blend_mode_names := []string{"NONE", "BLEND", "ADD", "MOD", "MUL", "SCREEN \"CUSTOM\""}

	surface := sdl.CreateSurface(RECT_SIZE, RECT_SIZE, sdl.PixelFormatRGBA8888)
	if surface == nil {
		panic("Couldn't create surface: " + sdl.GetError())
	}

	sdl.FillSurfaceRect(surface, nil, 0xFF0000FF) // Red
	red_rect_texture := sdl.CreateTextureFromSurface(renderer, surface)
	if red_rect_texture == nil {
		panic("Couldn't create texture: " + sdl.GetError())
	}
	defer sdl.DestroyTexture(red_rect_texture)

	sdl.FillSurfaceRect(surface, nil, 0x00FF00FF) // Green
	green_rect_texture := sdl.CreateTextureFromSurface(renderer, surface)
	if green_rect_texture == nil {
		panic("Couldn't create texture: " + sdl.GetError())
	}
	defer sdl.DestroyTexture(green_rect_texture)

	sdl.FillSurfaceRect(surface, nil, 0x0000FFFF) // Blue
	blue_rect_texture := sdl.CreateTextureFromSurface(renderer, surface)
	if blue_rect_texture == nil {
		panic("Couldn't create texture: " + sdl.GetError())
	}
	defer sdl.DestroyTexture(blue_rect_texture)

	sdl.DestroySurface(surface)

	alpha := uint8(255)
	var event sdl.Event

	running := true
	for running {
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit:
				running = false
			case sdl.EventKeyDown:
				e := event.Key()
				switch e.Scancode {
				case sdl.ScancodeEscape:
					running = false
				case sdl.ScancodeUp: // UP arrow increase alpha
					if int(alpha)+8 > 255 {
						alpha = 255
					} else {
						alpha += 8
					}
				case sdl.ScancodeDown: // DOWN arrow decrease alpha
					if int(alpha)-8 < 0 {
						alpha = 0
					} else {
						alpha -= 8
					}
				}
			}
		}

		sdl.SetRenderDrawColor(renderer, 0, 0, 0, sdl.AlphaOpaque)
		sdl.RenderClear(renderer)

		// Render checkerboard panels
		for i := range ROWS * COLS {
			// Loop through the panel pixels
			for y := panels[i].Y; y < float32(PANEL_SIZE)+panels[i].Y; y += float32(GRID_SIZE) {
				for x := panels[i].X; x < float32(PANEL_SIZE)+panels[i].X; x += float32(GRID_SIZE) {
					grid := sdl.FRect{X: x, Y: y, W: float32(GRID_SIZE), H: float32(GRID_SIZE)}
					dark := int32(x/float32(GRID_SIZE)+y/float32(GRID_SIZE)) % 2

					if dark != 0 {
						sdl.SetRenderDrawColor(renderer, 70, 70, 70, 255) // Darker color
					} else {
						sdl.SetRenderDrawColor(renderer, 110, 110, 110, 255) // Lighter color
					}

					sdl.RenderFillRect(renderer, &grid)
				}
			}

			// Label the blend mode
			sdl.SetRenderDrawColor(renderer, 255, 255, 255, sdl.AlphaOpaque)
			sdl.RenderDebugText(renderer, panels[i].X, panels[i].Y-15, blend_mode_names[i])
		}

		// Render panels
		sdl.RenderRects(renderer, panels)

		// Render UI text
		sdl.RenderDebugText(renderer, float32(WINDOW_WIDTH-176)/2, float32(WINDOW_HEIGHT-30), "UP/DOWN: CHANGE ALPHA")
		sdl.RenderDebugTextFormat(renderer, float32(WINDOW_WIDTH-80)/2, float32(WINDOW_HEIGHT-20), "ALPHA: %d", alpha)

		// Update textures alpha mod
		sdl.SetTextureAlphaMod(red_rect_texture, alpha)
		sdl.SetTextureAlphaMod(green_rect_texture, alpha)
		sdl.SetTextureAlphaMod(blue_rect_texture, alpha)

		// Render panels
		for i := range ROWS * COLS {
			// Update rects destination
			red_dst := sdl.FRect{X: panels[i].X + float32(RED_OFFSET), Y: panels[i].Y + float32(RED_OFFSET), W: float32(RECT_SIZE), H: float32(RECT_SIZE)}
			green_dst := sdl.FRect{X: panels[i].X + float32(GREEN_OFFSET), Y: panels[i].Y + float32(GREEN_OFFSET), W: float32(RECT_SIZE), H: float32(RECT_SIZE)}
			blue_dst := sdl.FRect{X: panels[i].X + float32(BLUE_OFFSET), Y: panels[i].Y + float32(BLUE_OFFSET), W: float32(RECT_SIZE), H: float32(RECT_SIZE)}

			// Apply the current blend mode
			supported := sdl.SetTextureBlendMode(red_rect_texture, blend_modes[i]) // just make sure the renderer supports this blend mode
			sdl.SetTextureBlendMode(green_rect_texture, blend_modes[i])
			sdl.SetTextureBlendMode(blue_rect_texture, blend_modes[i])

			// Render textures
			sdl.RenderTexture(renderer, red_rect_texture, nil, &red_dst)
			sdl.RenderTexture(renderer, green_rect_texture, nil, &green_dst)
			sdl.RenderTexture(renderer, blue_rect_texture, nil, &blue_dst)

			// Not all renderers support all blend modes. The renderer will try to pick something close in this case,
			// but it should be noted that the results might be unexpected, so we add "[UNSUPPORTED]" to this panel.
			if !supported {
				textWidth := float32(104.0)
				dst := sdl.FRect{X: panels[i].X + ((panels[i].W - textWidth) / 2.0), Y: panels[i].Y + (panels[i].H - 8), W: textWidth, H: 8}
				sdl.SetRenderDrawColor(renderer, 0, 0, 0, sdl.AlphaOpaque)
				sdl.RenderFillRect(renderer, &dst)
				sdl.SetRenderDrawColor(renderer, 255, 255, 255, sdl.AlphaOpaque)
				sdl.RenderDebugText(renderer, dst.X, dst.Y, "[UNSUPPORTED]")
			}
		}

		sdl.RenderPresent(renderer)
	}
}
