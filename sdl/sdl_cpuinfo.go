package sdl

// [GetNumLogicalCPUCores] gets the number of logical CPU cores available.
//
// [GetNumLogicalCPUCores]: https://wiki.libsdl.org/SDL3/SDL_GetNumLogicalCPUCores
func GetNumLogicalCPUCores() int32 {
	return sdlGetNumLogicalCPUCores()
}

// [GetCPUCacheLineSize] determines the L1 cache line size of the CPU.
//
// [GetCPUCacheLineSize]: https://wiki.libsdl.org/SDL3/SDL_GetCPUCacheLineSize
func GetCPUCacheLineSize() int32 {
	return sdlGetCPUCacheLineSize()
}

// [HasAltiVec] determines whether the CPU has AltiVec features.
//
// [HasAltiVec]: https://wiki.libsdl.org/SDL3/SDL_HasAltiVec
func HasAltiVec() bool {
	return sdlHasAltiVec()
}

// [HasMMX] determines whether the CPU has MMX features.
//
// [HasMMX]: https://wiki.libsdl.org/SDL3/SDL_HasMMX
func HasMMX() bool {
	return sdlHasMMX()
}

// [HasSSE] determines whether the CPU has SSE features.
//
// [HasSSE]: https://wiki.libsdl.org/SDL3/SDL_HasSSE
func HasSSE() bool {
	return sdlHasSSE()
}

// [HasSSE2] determines whether the CPU has SSE2 features.
//
// [HasSSE2]: https://wiki.libsdl.org/SDL3/SDL_HasSSE2
func HasSSE2() bool {
	return sdlHasSSE2()
}

// [HasSSE3] determines whether the CPU has SSE3 features.
//
// [HasSSE3]: https://wiki.libsdl.org/SDL3/SDL_HasSSE3
func HasSSE3() bool {
	return sdlHasSSE3()
}

// [HasSSE41] determines whether the CPU has SSE4.1 features.
//
// [HasSSE41]: https://wiki.libsdl.org/SDL3/SDL_HasSSE41
func HasSSE41() bool {
	return sdlHasSSE41()
}

// [HasSSE42] determines whether the CPU has SSE4.2 features.
//
// [HasSSE42]: https://wiki.libsdl.org/SDL3/SDL_HasSSE42
func HasSSE42() bool {
	return sdlHasSSE42()
}

// [HasAVX] determines whether the CPU has AVX features.
//
// [HasAVX]: https://wiki.libsdl.org/SDL3/SDL_HasAVX
func HasAVX() bool {
	return sdlHasAVX()
}

// [HasAVX2] determines whether the CPU has AVX2 features.
//
// [HasAVX2]: https://wiki.libsdl.org/SDL3/SDL_HasAVX2
func HasAVX2() bool {
	return sdlHasAVX2()
}

// [HasAVX512F] determines whether the CPU has AVX-512F (foundation) features.
//
// [HasAVX512F]: https://wiki.libsdl.org/SDL3/SDL_HasAVX512F
func HasAVX512F() bool {
	return sdlHasAVX512F()
}

// [HasARMSIMD] determines whether the CPU has ARM SIMD (ARMv6) features.
//
// [HasARMSIMD]: https://wiki.libsdl.org/SDL3/SDL_HasARMSIMD
func HasARMSIMD() bool {
	return sdlHasARMSIMD()
}

// [HasNEON] determines whether the CPU has NEON (ARM SIMD) features.
//
// [HasNEON]: https://wiki.libsdl.org/SDL3/SDL_HasNEON
func HasNEON() bool {
	return sdlHasNEON()
}

// [HasLSX] determines whether the CPU has LSX (LOONGARCH SIMD) features.
//
// [HasLSX]: https://wiki.libsdl.org/SDL3/SDL_HasLSX
func HasLSX() bool {
	return sdlHasLSX()
}

// [HasLASX] determines whether the CPU has LASX (LOONGARCH SIMD) features.
//
// [HasLASX]: https://wiki.libsdl.org/SDL3/SDL_HasLASX
func HasLASX() bool {
	return sdlHasLASX()
}

// [GetSystemRAM] gets the amount of RAM configured in the system.
//
// [GetSystemRAM]: https://wiki.libsdl.org/SDL3/SDL_GetSystemRAM
func GetSystemRAM() int32 {
	return sdlGetSystemRAM()
}

// [GetSIMDAlignment] reports the alignment this system needs for SIMD allocations.
//
// [GetSIMDAlignment]: https://wiki.libsdl.org/SDL3/SDL_GetSIMDAlignment
func GetSIMDAlignment() uint64 {
	return sdlGetSIMDAlignment()
}

// [GetSystemPageSize] reports the size of a page of memory.
//
// Available since SDL 3.4.0.
//
// [GetSystemPageSize]: https://wiki.libsdl.org/SDL3/SDL_GetSystemPageSize
func GetSystemPageSize() int32 {
	return sdlGetSystemPageSize()
}
