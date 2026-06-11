package main

import (
	"fmt"
	"io"
	"math"
	"math/rand/v2"
	"net/http"
	"os"
	"runtime"
	"time"
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/img"
	"github.com/jupiterrider/purego-sdl3/sdl"
	"github.com/jupiterrider/purego-sdl3/ttf"
)

// Asset files are read directly when running on Windows/Linux.
// When running on WASM the local http server should host them, symbolic link to assets directory in build dir is enough.

const WASM_TARGET_FPS float64 = 60

var WASM_FRAME_TIME time.Duration = time.Microsecond * time.Duration(math.Floor((1000.0/WASM_TARGET_FPS)*1000.0))

func main() {
	fmt.Println("Running on:", runtime.GOOS, runtime.GOARCH)
	runningOnWasm := runtime.GOOS == "js" && runtime.GOARCH == "wasm"
	fmt.Println("Compiled with Go:", runtime.Version())

	defer fmt.Println("SDL3 error:", sdl.GetError())

	major, minor, patch := sdl.GetVersion()
	fmt.Printf("SDL3 version: %d.%d.%d\n", major, minor, patch)
	fmt.Println("SDL3 revision:", sdl.GetRevision())
	major, minor, patch = img.Version()
	fmt.Printf("SDL3 Image version: %d.%d.%d\n", major, minor, patch)
	major, minor, patch = ttf.Version()
	fmt.Printf("SDL3 TTF version: %d.%d.%d\n", major, minor, patch)
	major, minor, patch = ttf.GetFreeTypeVersion()
	fmt.Printf("FreeType version: %d.%d.%d\n", major, minor, patch)
	major, minor, patch = ttf.GetHarfBuzzVersion()
	fmt.Printf("HarfBuzz version: %d.%d.%d\n", major, minor, patch)

	defer sdl.Quit()
	if !sdl.Init(sdl.InitVideo | sdl.InitAudio) {
		panic("SDL3 initialization failed, err: " + sdl.GetError())
	}

	defer ttf.Quit()
	if !ttf.Init() {
		panic("SDL3_TTF initialization failed, err: " + sdl.GetError())
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
	// gopherStream := sdl.IOFromConstMem(readFile("assets/gopher-happy.bmp"))
	// if gopherStream == nil {
	// 	panic(sdl.GetError())
	// }
	// gopherSurface := img.LoadIO(gopherStream, true)
	// if gopherSurface == nil {
	// 	panic(sdl.GetError())
	// }
	// defer sdl.DestroySurface(gopherSurface)
	// gopherTexture := sdl.CreateTextureFromSurface(renderer, gopherSurface)
	// Instead of loading sprite into surface and then into texture the img.LoadTextureIO() can be used
	gopherTexture := img.LoadTextureIO(renderer, sdl.IOFromConstMem(readFile("assets/gopher-happy.bmp")), true)
	if gopherTexture == nil {
		panic(gopherTexture)
	}
	defer sdl.DestroyTexture(gopherTexture)
	sdl.SetTextureScaleMode(gopherTexture, sdl.ScaleModeNearest)

	// Loading audio sample
	toneStream := sdl.IOFromConstMem(readFile("assets/tone.wav"))
	toneSample := AudioSample{}
	if !sdl.LoadWAVIO(toneStream, true, &toneSample.Spec, &toneSample.Data, &toneSample.Length) {
		panic(sdl.GetError())
	}
	defer sdl.Free(unsafe.Pointer(toneSample.Data))
	toneSample.Stream = sdl.OpenAudioDeviceStream(sdl.AudioDeviceDefaultPlayback, &toneSample.Spec, 0, nil)
	if toneSample.Stream == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyAudioStream(toneSample.Stream)
	sdl.SetAudioStreamGain(toneSample.Stream, 0.2) // Reduce the volume

	// TTF stuff, creating few text textures, TextEngine, Text objects, etc
	fontStream := sdl.IOFromMem(readFile("assets/OpenSans-Regular.ttf"))
	font := ttf.OpenFontIO(fontStream, true, 32)
	if font == nil {
		panic(sdl.GetError())
	}
	defer ttf.CloseFont(font)

	var fontVDPI, fontHDPI int32
	ttf.GetFontDPI(font, &fontVDPI, &fontHDPI)
	fmt.Printf("Loaded font, size: %v, height: %d, ascent: %d, descent: %d, line skip: %d, direction: %d, outline: %d, hinting: %d, kerning: %v, generation: %d, family name: %s, vertical DPI: %d, horizontal DPI: %d, style: %v, style name: %s, wrap alignment: %v\n",
		ttf.GetFontSize(font), ttf.GetFontHeight(font), ttf.GetFontAscent(font), ttf.GetFontDescent(font), ttf.GetFontLineSkip(font), ttf.GetFontDirection(font),
		ttf.GetFontOutline(font), ttf.GetFontHinting(font), ttf.GetFontKerning(font), ttf.GetFontGeneration(font), ttf.GetFontFamilyName(font),
		fontVDPI, fontHDPI, ttf.GetFontStyle(font), ttf.GetFontStyleName(font), ttf.GetFontWrapAlignment(font))

	text := "Hello from SDL3 TTF"
	textSurface := ttf.RenderTextBlended(font, text, uint64(len(text)), sdl.Color{R: 255, G: 0, B: 0, A: 0})
	if textSurface == nil {
		panic(sdl.GetError())
	}
	textTexture := sdl.CreateTextureFromSurface(renderer, textSurface)
	if textTexture == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyTexture(textTexture)
	sdl.DestroySurface(textSurface) // Surface is no longer needed
	text = "SDL3 Image\nalso works"
	text2Surface := ttf.RenderTextBlendedWrapped(font, text, uint64(len(text)), sdl.Color{R: 255, G: 255, B: 0, A: 0}, 0)
	if text2Surface == nil {
		panic(sdl.GetError())
	}
	text2Texture := sdl.CreateTextureFromSurface(renderer, text2Surface)
	if text2Texture == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyTexture(text2Texture)
	sdl.DestroySurface(text2Surface) // Surface is no longer needed
	runeSurface := ttf.RenderGlyphBlended(font, 'A', sdl.Color{R: 255, G: 0, B: 255, A: 0})
	if runeSurface == nil {
		panic(sdl.GetError())
	}
	runeTexture := sdl.CreateTextureFromSurface(renderer, runeSurface)
	if runeTexture == nil {
		panic(sdl.GetError())
	}
	defer sdl.DestroyTexture(runeTexture)
	sdl.DestroySurface(runeSurface) // Surface is no longer needed
	fmt.Println("Font has 'A' glyph rendered:", ttf.FontHasGlyph(font, 'A'))
	glyphMetrics := [5]int32{}
	ttf.GetGlyphMetrics(font, 'A', &glyphMetrics[0], &glyphMetrics[1], &glyphMetrics[2], &glyphMetrics[3], &glyphMetrics[4])
	fmt.Println("Glyph 'A' metrics:", glyphMetrics)
	engine := ttf.CreateRendererTextEngine(renderer)
	if engine == nil {
		panic(sdl.GetError())
	}
	defer ttf.DestroyRendererTextEngine(engine)
	text = "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	engineText := ttf.CreateText(engine, font, text, uint64(len(text)))
	if engineText == nil {
		panic(sdl.GetError())
	}
	defer ttf.DestroyText(engineText)
	engineTextColor := sdl.Color{R: 85, G: 170, B: 255, A: 255}
	text = " Sed tincidunt enim eget."
	ttf.AppendTextString(engineText, text, uint64(len(text)))
	var w, h int32
	ttf.GetStringSize(font, text, uint64(len(text)), &w, &h)
	fmt.Printf("Text '%s' has width: %v, height: %v\n", text, w, h)
	ttf.GetStringSizeWrapped(font, text, uint64(len(text)), 250, &w, &h)
	fmt.Printf("Text '%s' wrapped to 250px has width: %v, height: %v\n", text, w, h)
	ttf.GetTextSize(engineText, &w, &h)
	fmt.Printf("Text of the ttf.Text object has width: %v, height: %v\n", w, h)

	// FPS counter variables
	perfFreq, frameStart, frameCount := float32(sdl.GetPerformanceFrequency()), sdl.GetPerformanceCounter(), 0
	var fps, frameTime float32
	frameExecutionStart := time.Now()

	// Main loop start
	mousePos := sdl.FPoint{}
	running, mouseButtonDown, mouseLeftClick := true, false, false
	event := sdl.Event{}
	for running {
		mouseLeftClick = false

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
					mouseLeftClick = true
				}
			}
		}

		// Update audio stream before render part
		if toneSample.playing {
			toneSample.Update()
			if toneSample.finished {
				fmt.Println("Tone sample, finished playing")
				toneSample.playing = false
				toneSample.finished = false
				sdl.PauseAudioStreamDevice(toneSample.Stream)
			}
		}

		// Render start
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
		sdl.RenderTextureAffine(renderer, gopherTexture, nil, &sdl.FPoint{X: 820, Y: 220}, &sdl.FPoint{X: 900, Y: 260}, &sdl.FPoint{X: 780, Y: 290})

		// Geometry, an arrow
		sdl.RenderGeometry(renderer, nil, []sdl.Vertex{
			{Position: sdl.FPoint{X: 820, Y: 60}, Color: sdl.FColor{R: 1, G: 1, B: 1, A: 1}},  // Body top left
			{Position: sdl.FPoint{X: 920, Y: 60}, Color: sdl.FColor{R: 1, G: 0, B: 0, A: 1}},  // Body top right
			{Position: sdl.FPoint{X: 820, Y: 110}, Color: sdl.FColor{R: 0, G: 1, B: 0, A: 1}}, // Body bottom left
			{Position: sdl.FPoint{X: 920, Y: 110}, Color: sdl.FColor{R: 0, G: 0, B: 1, A: 1}}, // Body bottom right
			{Position: sdl.FPoint{X: 920, Y: 30}, Color: sdl.FColor{R: 1, G: 1, B: 0, A: 1}},  // Point top
			{Position: sdl.FPoint{X: 920, Y: 140}, Color: sdl.FColor{R: 0, G: 1, B: 1, A: 1}}, // Point bottom
			{Position: sdl.FPoint{X: 975, Y: 85}, Color: sdl.FColor{R: 1, G: 0, B: 1, A: 1}},  // Point right
		}, []int32{
			1, 0, 2,
			1, 2, 3,
			6, 4, 1,
			6, 5, 3,
			6, 1, 3,
		})

		// Button to play audio sample
		sdl.SetRenderDrawColor(renderer, 255, 255, 255, 255)
		sdl.RenderDebugText(renderer, 40, 530, "Click to play audio sample")
		r = sdl.FRect{X: 10, Y: 523, W: 20, H: 20}
		sdl.RenderRect(renderer, &r)
		if sdl.PointInRectFloat(mousePos, r) {
			sdl.RenderFillRect(renderer, &r)
			if mouseLeftClick {
				if toneSample.playing {
					fmt.Println("Tone sample, reset")
					sdl.ClearAudioStream(toneSample.Stream)
				} else {
					fmt.Println("Tone sample, starts playing")
				}
				toneSample.playing = true
				sdl.PutAudioStreamData(toneSample.Stream, toneSample.Data, int32(toneSample.Length)) // Queue data to be played
				if !sdl.ResumeAudioStreamDevice(toneSample.Stream) {
					panic(sdl.GetError())
				}
			}
		}

		// Simple debug text
		sdl.SetRenderDrawColorFloat(renderer, 1, 1, 1, 1)
		sdl.RenderDebugText(renderer, 10, 10, "Hello")

		// Mouse position text
		sdl.RenderDebugTextFormat(renderer, 10, 30, "Mouse position: {x: %.0f, y: %.0f}", mousePos.X, mousePos.Y)

		// Text texture created using SDL3 TTF
		sdl.RenderTexture(renderer, textTexture, nil, &sdl.FRect{X: 850, Y: 160, W: float32(textTexture.W), H: float32(textTexture.H)})
		sdl.RenderTexture(renderer, text2Texture, nil, &sdl.FRect{X: 910, Y: 230, W: float32(text2Texture.W), H: float32(text2Texture.H)})
		sdl.RenderTexture(renderer, runeTexture, nil, &sdl.FRect{X: 860, Y: 120, W: float32(runeTexture.W), H: float32(runeTexture.H)})
		engineTextColor.R, engineTextColor.G, engineTextColor.B = (engineTextColor.R+1)%255, (engineTextColor.G+1)%255, (engineTextColor.B+1)%255
		ttf.SetTextColor(engineText, engineTextColor.R, engineTextColor.G, engineTextColor.B, engineTextColor.A)
		ttf.DrawRendererText(engineText, 20, 560)

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

		// On Wasm the main loop has to be somehow slowed down?
		frameExecutionTime := time.Since(frameExecutionStart)
		if runningOnWasm {
			if sleepDur := WASM_FRAME_TIME - frameExecutionTime - time.Millisecond; sleepDur > 0 {
				time.Sleep(sleepDur)
			}
		}
		frameExecutionStart = time.Now()
	}
}

type AudioSample struct {
	Spec   sdl.AudioSpec    // Audio format of the audio sample
	Data   *uint8           // Audio data of the audio sample
	Length uint32           // Audio data length of the audio sample
	Stream *sdl.AudioStream // Handle to audio stream to allow playing of the audio sample

	playing  bool // Is the audio sample currently playing?
	finished bool // Has the audio sample finished playing?

	lastQueuedDataLength   int32
	finishedPlayingCounter uint8
}

func (as *AudioSample) Update() {
	dataQueued := sdl.GetAudioStreamQueued(as.Stream)
	if as.lastQueuedDataLength == dataQueued {
		as.finishedPlayingCounter++
	} else {
		as.finishedPlayingCounter = 0
		as.lastQueuedDataLength = dataQueued
	}
	if dataQueued == 0 || as.finishedPlayingCounter >= 10 {
		as.finished = true
		as.finishedPlayingCounter = 0
	}
}

func readFile(path string) []byte {
	var data []byte
	var err error
	if runtime.GOOS == "js" && runtime.GOARCH == "wasm" {
		// When running on WASM request local http server for the file
		resp, err := http.DefaultClient.Get(path)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
	} else {
		// When not running on WASM read the file yourself
		data, err = os.ReadFile(path)
	}
	if err != nil {
		panic(err)
	}
	return data
}
