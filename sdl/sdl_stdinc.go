package sdl

import (
	"unsafe"
)

// [Time] is a structure specifying SDL time.
// SDL times are signed, 64-bit integers representing nanoseconds since the Unix epoch (Jan 1, 1970).
//
// [Time]: https://wiki.libsdl.org/SDL3/SDL_Time
type Time int64

const (
	MaxTime Time = 9223372036854775807
	MinTime Time = -9223372036854775808
)

const FltEpsilon = 0x0.000002p0

// [FourCC] defines a four character code as a Uint32.
//
// [FourCC]: https://wiki.libsdl.org/SDL3/SDL_FOURCC
func FourCC(a, b, c, d byte) uint32 {
	return uint32(a) | uint32(b)<<8 | uint32(c)<<16 | uint32(d)<<24
}

// [Malloc] allocates uninitialized memory.
//
// [Malloc]: https://wiki.libsdl.org/SDL3/SDL_malloc
// func Malloc(size uint64) unsafe.Pointer {
//	return sdlmalloc(size)
// }

// [Calloc] allocates a zero-initialized array.
//
// [Calloc]: https://wiki.libsdl.org/SDL3/SDL_calloc
// func Calloc(nmemb uint64, size uint64) unsafe.Pointer {
//	return sdlcalloc(nmemb, size)
// }

// [Realloc] changes the size of allocated memory.
//
// [Realloc]: https://wiki.libsdl.org/SDL3/SDL_realloc
// func realloc(mem unsafe.Pointer, size uint64) unsafe.Pointer {
//	return sdlrealloc(mem, size)
// }

// [Free] frees allocated memory.
//
// If mem is nil, this function does nothing.
//
// [Free]: https://wiki.libsdl.org/SDL3/SDL_free
func Free(mem unsafe.Pointer) {
	sdlfree(mem)
}

// [MallocFunc] defines a callback used to implement [Malloc].
//
// [MallocFunc]: https://wiki.libsdl.org/SDL3/SDL_malloc_func
type MallocFunc uintptr

// [CallocFunc] defines a callback used to implement [Calloc].
//
// [CallocFunc]: https://wiki.libsdl.org/SDL3/SDL_calloc_func
type CallocFunc uintptr

// [ReallocFunc] defines a callback used to implement [Realloc].
//
// [ReallocFunc]: https://wiki.libsdl.org/SDL3/SDL_realloc_func
type ReallocFunc uintptr

// [FreeFunc] defines a callback used to implement [Free].
//
// [FreeFunc]: https://wiki.libsdl.org/SDL3/SDL_free_func
type FreeFunc uintptr

// [GetOriginalMemoryFunctions] gets the original set of SDL memory functions.
//
// [GetOriginalMemoryFunctions]: https://wiki.libsdl.org/SDL3/SDL_GetOriginalMemoryFunctions
// func GetOriginalMemoryFunctions(malloc_func *malloc_func, calloc_func *calloc_func, realloc_func *realloc_func, free_func *free_func)  {
//	sdlGetOriginalMemoryFunctions(malloc_func, calloc_func, realloc_func, free_func)
// }

// [GetMemoryFunctions] gets the current set of SDL memory functions.
//
// [GetMemoryFunctions]: https://wiki.libsdl.org/SDL3/SDL_GetMemoryFunctions
// func GetMemoryFunctions(malloc_func *malloc_func, calloc_func *calloc_func, realloc_func *realloc_func, free_func *free_func)  {
//	sdlGetMemoryFunctions(malloc_func, calloc_func, realloc_func, free_func)
// }

// [SetMemoryFunctions] replaces SDL's memory allocation functions with a custom set.
//
// [SetMemoryFunctions]: https://wiki.libsdl.org/SDL3/SDL_SetMemoryFunctions
// func SetMemoryFunctions(malloc_func malloc_func, calloc_func calloc_func, realloc_func realloc_func, free_func free_func) bool {
//	return sdlSetMemoryFunctions(malloc_func, calloc_func, realloc_func, free_func)
// }

// [AlignedAlloc] allocates memory aligned to a specific alignment.
//
// [AlignedAlloc]: https://wiki.libsdl.org/SDL3/SDL_aligned_alloc
// func AlignedAlloc(alignment uint64, size uint64) unsafe.Pointer {
//	return sdlaligned_alloc(alignment, size)
// }

// [AlignedFree] frees memory allocated by [aligned_alloc].
//
// [AlignedFree]: https://wiki.libsdl.org/SDL3/SDL_aligned_free
// func AlignedFree(mem unsafe.Pointer)  {
//	sdlaligned_free(mem)
// }

// [GetNumAllocations] gets the number of outstanding (unfreed) allocations.
//
// [GetNumAllocations]: https://wiki.libsdl.org/SDL3/SDL_GetNumAllocations
// func GetNumAllocations() int32 {
//	return sdlGetNumAllocations()
// }

// [Environment] is a thread-safe set of environment variables.
//
// [Environment]: https://wiki.libsdl.org/SDL3/SDL_Environment
type Environment struct{}

// [GetEnvironment] gets the process environment.
//
// [GetEnvironment]: https://wiki.libsdl.org/SDL3/SDL_GetEnvironment
// func GetEnvironment() *Environment {
//	return sdlGetEnvironment()
// }

// [CreateEnvironment] creates a set of environment variables.
//
// [CreateEnvironment]: https://wiki.libsdl.org/SDL3/SDL_CreateEnvironment
// func CreateEnvironment(populated bool) *Environment {
//	return sdlCreateEnvironment(populated)
// }

// [GetEnvironmentVariable] gets the value of a variable in the environment.
//
// [GetEnvironmentVariable]: https://wiki.libsdl.org/SDL3/SDL_GetEnvironmentVariable
// func GetEnvironmentVariable(env *Environment, name string) string {
//	return sdlGetEnvironmentVariable(env, name)
// }

// [GetEnvironmentVariables] gets all variables in the environment.
//
// [GetEnvironmentVariables]: https://wiki.libsdl.org/SDL3/SDL_GetEnvironmentVariables
// func GetEnvironmentVariables(env *Environment) **byte {
//	return sdlGetEnvironmentVariables(env)
// }

// [SetEnvironmentVariable] sets the value of a variable in the environment.
//
// [SetEnvironmentVariable]: https://wiki.libsdl.org/SDL3/SDL_SetEnvironmentVariable
// func SetEnvironmentVariable(env *Environment, name string, value string, overwrite bool) bool {
//	return sdlSetEnvironmentVariable(env, name, value, overwrite)
// }

// [UnsetEnvironmentVariable] clears a variable from the environment.
//
// [UnsetEnvironmentVariable]: https://wiki.libsdl.org/SDL3/SDL_UnsetEnvironmentVariable
// func UnsetEnvironmentVariable(env *Environment, name string) bool {
//	return sdlUnsetEnvironmentVariable(env, name)
// }

// [DestroyEnvironment] destroys a set of environment variables.
//
// [DestroyEnvironment]: https://wiki.libsdl.org/SDL3/SDL_DestroyEnvironment
// func DestroyEnvironment(env *Environment) {
//	sdlDestroyEnvironment(env)
// }

// [Getenv] gets the value of a variable in the environment.
//
// [Getenv]: https://wiki.libsdl.org/SDL3/SDL_getenv
// func Getenv(name string) string {
//	return sdlgetenv(name)
// }

// [GetenvUnsafe] gets the value of a variable in the environment.
//
// [GetenvUnsafe]: https://wiki.libsdl.org/SDL3/SDL_getenv_unsafe
// func GetenvUnsafe(name string) string {
//	return sdlgetenv_unsafe(name)
// }

// [SetenvUnsafe] sets the value of a variable in the environment.
//
// [SetenvUnsafe]: https://wiki.libsdl.org/SDL3/SDL_setenv_unsafe
// func SetenvUnsafe(name string, value string, overwrite int32) int32 {
//	return sdlsetenv_unsafe(name, value, overwrite)
// }

// [UnsetenvUnsafe] clears a variable from the environment.
//
// [UnsetenvUnsafe]: https://wiki.libsdl.org/SDL3/SDL_unsetenv_unsafe
// func UnsetenvUnsafe(name string) int32 {
//	return sdlunsetenv_unsafe(name)
// }

// [CompareCallback] defines a callback used with SDL sorting and binary search functions.
//
// [CompareCallback]: https://wiki.libsdl.org/SDL3/SDL_CompareCallback
type CompareCallback uintptr

// [Qsort] sorts an array.
//
// [Qsort]: https://wiki.libsdl.org/SDL3/SDL_qsort
// func Qsort(base unsafe.Pointer, nmemb uint64, size uint64, compare CompareCallback)  {
//	sdlqsort(base, nmemb, size, compare)
// }

// [Bsearch] performs a binary search on a previously sorted array.
//
// [Bsearch]: https://wiki.libsdl.org/SDL3/SDL_bsearch
// func Bsearch(key unsafe.Pointer, base unsafe.Pointer, nmemb uint64, size uint64, compare CompareCallback) unsafe.Pointer {
//	return sdlbsearch(key, base, nmemb, size, compare)
// }

// [CompareCallbackR] defines a callback used with SDL sorting and binary search functions.
//
// [CompareCallbackR]: https://wiki.libsdl.org/SDL3/SDL_CompareCallback_r
type CompareCallbackR uintptr

// [QsortR] sorts an array, passing a userdata pointer to the compare function.
//
// [QsortR]: https://wiki.libsdl.org/SDL3/SDL_qsort_r
// func QsortR(base unsafe.Pointer, nmemb uint64, size uint64, compare CompareCallback_r, userdata unsafe.Pointer)  {
//	sdlqsort_r(base, nmemb, size, compare, userdata)
// }

// [BsearchR] performs a binary search on a previously sorted array, passing a userdata.
//
// [BsearchR]: https://wiki.libsdl.org/SDL3/SDL_bsearch_r
// func BsearchR(key unsafe.Pointer, base unsafe.Pointer, nmemb uint64, size uint64, compare CompareCallback_r, userdata unsafe.Pointer) unsafe.Pointer {
//	return sdlbsearch_r(key, base, nmemb, size, compare, userdata)
// }

// [Abs] computes the absolute value of `x`.
//
// [Abs]: https://wiki.libsdl.org/SDL3/SDL_abs
// func Abs(x int32) int32 {
//	return sdlabs(x)
// }

// [Isalpha] queries if a character is alphabetic (a letter).
//
// [Isalpha]: https://wiki.libsdl.org/SDL3/SDL_isalpha
// func Isalpha(x int32) int32 {
//	return sdlisalpha(x)
// }

// [Isalnum] queries if a character is alphabetic (a letter) or a number.
//
// [Isalnum]: https://wiki.libsdl.org/SDL3/SDL_isalnum
// func Isalnum(x int32) int32 {
//	return sdlisalnum(x)
// }

// [Isblank] reports if a character is blank (a space or tab).
//
// [Isblank]: https://wiki.libsdl.org/SDL3/SDL_isblank
// func Isblank(x int32) int32 {
//	return sdlisblank(x)
// }

// [Iscntrl] reports if a character is a control character.
//
// [Iscntrl]: https://wiki.libsdl.org/SDL3/SDL_iscntrl
// func Iscntrl(x int32) int32 {
//	return sdliscntrl(x)
// }

// [Isdigit] reports if a character is a numeric digit.
//
// [Isdigit]: https://wiki.libsdl.org/SDL3/SDL_isdigit
// func Isdigit(x int32) int32 {
//	return sdlisdigit(x)
// }

// [Isxdigit] reports if a character is a hexadecimal digit.
//
// [Isxdigit]: https://wiki.libsdl.org/SDL3/SDL_isxdigit
// func Isxdigit(x int32) int32 {
//	return sdlisxdigit(x)
// }

// [Ispunct] reports if a character is a punctuation mark.
//
// [Ispunct]: https://wiki.libsdl.org/SDL3/SDL_ispunct
// func Ispunct(x int32) int32 {
//	return sdlispunct(x)
// }

// [Isspace] reports if a character is whitespace.
//
// [Isspace]: https://wiki.libsdl.org/SDL3/SDL_isspace
// func Isspace(x int32) int32 {
//	return sdlisspace(x)
// }

// [Isupper] reports if a character is upper case.
//
// [Isupper]: https://wiki.libsdl.org/SDL3/SDL_isupper
// func Isupper(x int32) int32 {
//	return sdlisupper(x)
// }

// [Islower] reports if a character is lower case.
//
// [Islower]: https://wiki.libsdl.org/SDL3/SDL_islower
// func Islower(x int32) int32 {
//	return sdlislower(x)
// }

// [Isprint] reports if a character is "printable".
//
// [Isprint]: https://wiki.libsdl.org/SDL3/SDL_isprint
// func Isprint(x int32) int32 {
//	return sdlisprint(x)
// }

// [Isgraph] reports if a character is any "printable" except space.
//
// [Isgraph]: https://wiki.libsdl.org/SDL3/SDL_isgraph
// func Isgraph(x int32) int32 {
//	return sdlisgraph(x)
// }

// [Toupper] converts low-ASCII English letters to uppercase.
//
// [Toupper]: https://wiki.libsdl.org/SDL3/SDL_toupper
// func Toupper(x int32) int32 {
//	return sdltoupper(x)
// }

// [Tolower] converts low-ASCII English letters to lowercase.
//
// [Tolower]: https://wiki.libsdl.org/SDL3/SDL_tolower
// func Tolower(x int32) int32 {
//	return sdltolower(x)
// }

// [Crc16] calculates a CRC-16 value.
//
// [Crc16]: https://wiki.libsdl.org/SDL3/SDL_crc16
// func Crc16(crc uint16, data unsafe.Pointer, len uint64) uint16 {
//	return sdlcrc16(crc, data, len)
// }

// [Crc32] calculates a CRC-32 value.
//
// [Crc32]: https://wiki.libsdl.org/SDL3/SDL_crc32
// func Crc32(crc uint32, data unsafe.Pointer, len uint64) uint32 {
//	return sdlcrc32(crc, data, len)
// }

// [Murmur332] calculates a 32-bit MurmurHash3 value for a block of data.
//
// [Murmur332]: https://wiki.libsdl.org/SDL3/SDL_murmur3_32
// func Murmur332(data unsafe.Pointer, len uint64, seed uint32) uint32 {
//	return sdlmurmur3_32(data, len, seed)
// }

// [Memcpy] copies non-overlapping memory.
//
// [Memcpy]: https://wiki.libsdl.org/SDL3/SDL_memcpy
// func Memcpy(dst unsafe.Pointer, src unsafe.Pointer, len uint64) unsafe.Pointer {
//	return sdlmemcpy(dst, src, len)
// }

// [Memmove] copies memory ranges that might overlap.
//
// [Memmove]: https://wiki.libsdl.org/SDL3/SDL_memmove
// func Memmove(dst unsafe.Pointer, src unsafe.Pointer, len uint64) unsafe.Pointer {
//	return sdlmemmove(dst, src, len)
// }

// [Memset] initializes all bytes of buffer of memory to a specific value.
//
// [Memset]: https://wiki.libsdl.org/SDL3/SDL_memset
// func Memset(dst unsafe.Pointer, c int32, len uint64) unsafe.Pointer {
//	return sdlmemset(dst, c, len)
// }

// [Memset4] initializes all 32-bit words of buffer of memory to a specific value.
//
// [Memset4]: https://wiki.libsdl.org/SDL3/SDL_memset4
// func Memset4(dst unsafe.Pointer, val uint32, dwords uint64) unsafe.Pointer {
//	return sdlmemset4(dst, val, dwords)
// }

// [Memcmp] compares two buffers of memory.
//
// [Memcmp]: https://wiki.libsdl.org/SDL3/SDL_memcmp
// func Memcmp(s1 unsafe.Pointer, s2 unsafe.Pointer, len uint64) int32 {
//	return sdlmemcmp(s1, s2, len)
// }

// [Wcslen] this works exactly like wcslen but doesn't require access to a C runtime.
//
// [Wcslen]: https://wiki.libsdl.org/SDL3/SDL_wcslen
// func Wcslen(wstr *wchar_t) uint64 {
//	return sdlwcslen(wstr)
// }

// [Wcsnlen] this works exactly like wcsnlen but doesn't require access to a C.
//
// [Wcsnlen]: https://wiki.libsdl.org/SDL3/SDL_wcsnlen
// func Wcsnlen(wstr *wchar_t, maxlen uint64) uint64 {
//	return sdlwcsnlen(wstr, maxlen)
// }

// [Wcslcpy] copies a wide string.
//
// [Wcslcpy]: https://wiki.libsdl.org/SDL3/SDL_wcslcpy
// func Wcslcpy(dst *wchar_t, src *wchar_t, maxlen uint64) uint64 {
//	return sdlwcslcpy(dst, src, maxlen)
// }

// [Wcslcat] concatenates wide strings.
//
// [Wcslcat]: https://wiki.libsdl.org/SDL3/SDL_wcslcat
// func Wcslcat(dst *wchar_t, src *wchar_t, maxlen uint64) uint64 {
//	return sdlwcslcat(dst, src, maxlen)
// }

// [Wcsdup] allocates a copy of a wide string.
//
// [Wcsdup]: https://wiki.libsdl.org/SDL3/SDL_wcsdup
// func Wcsdup(wstr *wchar_t) *wchar_t {
//	return sdlwcsdup(wstr)
// }

// [Wcsstr] searchs a wide string for the first instance of a specific substring.
//
// [Wcsstr]: https://wiki.libsdl.org/SDL3/SDL_wcsstr
// func Wcsstr(haystack *wchar_t, needle *wchar_t) *wchar_t {
//	return sdlwcsstr(haystack, needle)
// }

// [Wcsnstr] searchs a wide string, up to n wide chars, for the first instance of a.
//
// [Wcsnstr]: https://wiki.libsdl.org/SDL3/SDL_wcsnstr
// func Wcsnstr(haystack *wchar_t, needle *wchar_t, maxlen uint64) *wchar_t {
//	return sdlwcsnstr(haystack, needle, maxlen)
// }

// [Wcscmp] compares two null-terminated wide strings.
//
// [Wcscmp]: https://wiki.libsdl.org/SDL3/SDL_wcscmp
// func Wcscmp(str1 *wchar_t, str2 *wchar_t) int32 {
//	return sdlwcscmp(str1, str2)
// }

// [Wcsncmp] compares two wide strings up to a number of wchar_t values.
//
// [Wcsncmp]: https://wiki.libsdl.org/SDL3/SDL_wcsncmp
// func Wcsncmp(str1 *wchar_t, str2 *wchar_t, maxlen uint64) int32 {
//	return sdlwcsncmp(str1, str2, maxlen)
// }

// [Wcscasecmp] compares two null-terminated wide strings, case-insensitively.
//
// [Wcscasecmp]: https://wiki.libsdl.org/SDL3/SDL_wcscasecmp
// func Wcscasecmp(str1 *wchar_t, str2 *wchar_t) int32 {
//	return sdlwcscasecmp(str1, str2)
// }

// [Wcsncasecmp] compares two wide strings, case-insensitively, up to a number of wchar_t.
//
// [Wcsncasecmp]: https://wiki.libsdl.org/SDL3/SDL_wcsncasecmp
// func Wcsncasecmp(str1 *wchar_t, str2 *wchar_t, maxlen uint64) int32 {
//	return sdlwcsncasecmp(str1, str2, maxlen)
// }

// [Wcstol] parses a `long` from a wide string.
//
// [Wcstol]: https://wiki.libsdl.org/SDL3/SDL_wcstol
// func Wcstol(str *wchar_t, endp **wchar_t, base int32) int64 {
//	return sdlwcstol(str, endp, base)
// }

// [Strlen] this works exactly like strlen but doesn't require access to a C runtime.
//
// [Strlen]: https://wiki.libsdl.org/SDL3/SDL_strlen
// func Strlen(str string) uint64 {
//	return sdlstrlen(str)
// }

// [Strnlen] this works exactly like strnlen but doesn't require access to a C.
//
// [Strnlen]: https://wiki.libsdl.org/SDL3/SDL_strnlen
// func Strnlen(str string, maxlen uint64) uint64 {
//	return sdlstrnlen(str, maxlen)
// }

// [Strlcpy] copies a string.
//
// [Strlcpy]: https://wiki.libsdl.org/SDL3/SDL_strlcpy
// func Strlcpy(dst string, src string, maxlen uint64) uint64 {
//	return sdlstrlcpy(dst, src, maxlen)
// }

// [Utf8strlcpy] copies an UTF-8 string.
//
// [Utf8strlcpy]: https://wiki.libsdl.org/SDL3/SDL_utf8strlcpy
// func Utf8strlcpy(dst string, src string, dst_bytes uint64) uint64 {
//	return sdlutf8strlcpy(dst, src, dst_bytes)
// }

// [Strlcat] concatenates strings.
//
// [Strlcat]: https://wiki.libsdl.org/SDL3/SDL_strlcat
// func Strlcat(dst string, src string, maxlen uint64) uint64 {
//	return sdlstrlcat(dst, src, maxlen)
// }

// [Strdup] allocates a copy of a string.
//
// [Strdup]: https://wiki.libsdl.org/SDL3/SDL_strdup
// func Strdup(str string) string {
//	return sdlstrdup(str)
// }

// [Strndup] allocates a copy of a string, up to n characters.
//
// [Strndup]: https://wiki.libsdl.org/SDL3/SDL_strndup
// func Strndup(str string, maxlen uint64) string {
//	return sdlstrndup(str, maxlen)
// }

// [Strrev] reverses a string's contents.
//
// [Strrev]: https://wiki.libsdl.org/SDL3/SDL_strrev
// func Strrev(str string) string {
//	return sdlstrrev(str)
// }

// [Strupr] converts a string to uppercase.
//
// [Strupr]: https://wiki.libsdl.org/SDL3/SDL_strupr
// func Strupr(str string) string {
//	return sdlstrupr(str)
// }

// [Strlwr] converts a string to lowercase.
//
// [Strlwr]: https://wiki.libsdl.org/SDL3/SDL_strlwr
// func Strlwr(str string) string {
//	return sdlstrlwr(str)
// }

// [Strchr] searchs a string for the first instance of a specific byte.
//
// [Strchr]: https://wiki.libsdl.org/SDL3/SDL_strchr
// func Strchr(str string, c int32) string {
//	return sdlstrchr(str, c)
// }

// [Strrchr] searchs a string for the last instance of a specific byte.
//
// [Strrchr]: https://wiki.libsdl.org/SDL3/SDL_strrchr
// func Strrchr(str string, c int32) string {
//	return sdlstrrchr(str, c)
// }

// [Strstr] searchs a string for the first instance of a specific substring.
//
// [Strstr]: https://wiki.libsdl.org/SDL3/SDL_strstr
func Strstr(haystack, needle string) string {
	return sdlstrstr(haystack, needle)
}

// [Strnstr] searchs a string, up to n bytes, for the first instance of a specific.
//
// [Strnstr]: https://wiki.libsdl.org/SDL3/SDL_strnstr
// func Strnstr(haystack string, needle string, maxlen uint64) string {
//	return sdlstrnstr(haystack, needle, maxlen)
// }

// [Strcasestr] searchs a UTF-8 string for the first instance of a specific substring,.
//
// [Strcasestr]: https://wiki.libsdl.org/SDL3/SDL_strcasestr
// func Strcasestr(haystack string, needle string) string {
//	return sdlstrcasestr(haystack, needle)
// }

// [StrtokR] this works exactly like strtok_r but doesn't require access to a C.
//
// [StrtokR]: https://wiki.libsdl.org/SDL3/SDL_strtok_r
// func StrtokR(str string, delim string, saveptr **byte) string {
//	return sdlstrtok_r(str, delim, saveptr)
// }

// [Utf8strlen] counts the number of codepoints in a UTF-8 string.
//
// [Utf8strlen]: https://wiki.libsdl.org/SDL3/SDL_utf8strlen
// func Utf8strlen(str string) uint64 {
//	return sdlutf8strlen(str)
// }

// [Utf8strnlen] counts the number of codepoints in a UTF-8 string, up to n bytes.
//
// [Utf8strnlen]: https://wiki.libsdl.org/SDL3/SDL_utf8strnlen
// func Utf8strnlen(str string, bytes uint64) uint64 {
//	return sdlutf8strnlen(str, bytes)
// }

// [Itoa] converts an integer into a string.
//
// [Itoa]: https://wiki.libsdl.org/SDL3/SDL_itoa
// func Itoa(value int32, str string, radix int32) string {
//	return sdlitoa(value, str, radix)
// }

// [Uitoa] converts an unsigned integer into a string.
//
// [Uitoa]: https://wiki.libsdl.org/SDL3/SDL_uitoa
// func Uitoa(value uint32, str string, radix int32) string {
//	return sdluitoa(value, str, radix)
// }

// [Ltoa] converts a long integer into a string.
//
// [Ltoa]: https://wiki.libsdl.org/SDL3/SDL_ltoa
// func Ltoa(value int64, str string, radix int32) string {
//	return sdlltoa(value, str, radix)
// }

// [Ultoa] converts an unsigned long integer into a string.
//
// [Ultoa]: https://wiki.libsdl.org/SDL3/SDL_ultoa
// func Ultoa(value uint64, str string, radix int32) string {
//	return sdlultoa(value, str, radix)
// }

// [Lltoa] converts a long long integer into a string.
//
// [Lltoa]: https://wiki.libsdl.org/SDL3/SDL_lltoa
// func Lltoa(value int64, str string, radix int32) string {
//	return sdllltoa(value, str, radix)
// }

// [Ulltoa] converts an unsigned long long integer into a string.
//
// [Ulltoa]: https://wiki.libsdl.org/SDL3/SDL_ulltoa
// func Ulltoa(value uint64, str string, radix int32) string {
//	return sdlulltoa(value, str, radix)
// }

// [Atoi] parses an `int` from a string.
//
// [Atoi]: https://wiki.libsdl.org/SDL3/SDL_atoi
// func Atoi(str string) int32 {
//	return sdlatoi(str)
// }

// [Atof] parses a `double` from a string.
//
// [Atof]: https://wiki.libsdl.org/SDL3/SDL_atof
// func Atof(str string) float64 {
//	return sdlatof(str)
// }

// [Strtol] parses a `long` from a string.
//
// [Strtol]: https://wiki.libsdl.org/SDL3/SDL_strtol
// func Strtol(str string, endp **byte, base int32) int64 {
//	return sdlstrtol(str, endp, base)
// }

// [Strtoul] parses an `unsigned long` from a string.
//
// [Strtoul]: https://wiki.libsdl.org/SDL3/SDL_strtoul
// func Strtoul(str string, endp **byte, base int32) uint64 {
//	return sdlstrtoul(str, endp, base)
// }

// [Strtoll] parses a `long long` from a string.
//
// [Strtoll]: https://wiki.libsdl.org/SDL3/SDL_strtoll
// func Strtoll(str string, endp **byte, base int32) int64 {
//	return sdlstrtoll(str, endp, base)
// }

// [Strtoull] parses an `unsigned long long` from a string.
//
// [Strtoull]: https://wiki.libsdl.org/SDL3/SDL_strtoull
// func Strtoull(str string, endp **byte, base int32) uint64 {
//	return sdlstrtoull(str, endp, base)
// }

// [Strtod] parses a `double` from a string.
//
// [Strtod]: https://wiki.libsdl.org/SDL3/SDL_strtod
// func Strtod(str string, endp **byte) float64 {
//	return sdlstrtod(str, endp)
// }

// [Strcmp] compares two null-terminated UTF-8 strings.
//
// [Strcmp]: https://wiki.libsdl.org/SDL3/SDL_strcmp
// func Strcmp(str1 string, str2 string) int32 {
//	return sdlstrcmp(str1, str2)
// }

// [Strncmp] compares two UTF-8 strings up to a number of bytes.
//
// [Strncmp]: https://wiki.libsdl.org/SDL3/SDL_strncmp
// func Strncmp(str1 string, str2 string, maxlen uint64) int32 {
//	return sdlstrncmp(str1, str2, maxlen)
// }

// [Strcasecmp] compares two null-terminated UTF-8 strings, case-insensitively.
//
// [Strcasecmp]: https://wiki.libsdl.org/SDL3/SDL_strcasecmp
// func Strcasecmp(str1 string, str2 string) int32 {
//	return sdlstrcasecmp(str1, str2)
// }

// [Strncasecmp] compares two UTF-8 strings, case-insensitively, up to a number of bytes.
//
// [Strncasecmp]: https://wiki.libsdl.org/SDL3/SDL_strncasecmp
// func Strncasecmp(str1 string, str2 string, maxlen uint64) int32 {
//	return sdlstrncasecmp(str1, str2, maxlen)
// }

// [Strpbrk] searches a string for the first occurrence of any character contained in a.
//
// [Strpbrk]: https://wiki.libsdl.org/SDL3/SDL_strpbrk
// func Strpbrk(str string, breakset string) string {
//	return sdlstrpbrk(str, breakset)
// }

// [StepUTF8] decodes a UTF-8 string, one Unicode codepoint at a time.
//
// [StepUTF8]: https://wiki.libsdl.org/SDL3/SDL_StepUTF8
// func StepUTF8(pstr **byte, pslen *uint64) uint32 {
//	return sdlStepUTF8(pstr, pslen)
// }

// [StepBackUTF8] decodes a UTF-8 string in reverse, one Unicode codepoint at a time.
//
// [StepBackUTF8]: https://wiki.libsdl.org/SDL3/SDL_StepBackUTF8
// func StepBackUTF8(start string, pstr **byte) uint32 {
//	return sdlStepBackUTF8(start, pstr)
// }

// [UCS4ToUTF8] converts a single Unicode codepoint to UTF-8.
//
// [UCS4ToUTF8]: https://wiki.libsdl.org/SDL3/SDL_UCS4ToUTF8
// func UCS4ToUTF8(codepoint uint32, dst string) string {
//	return sdlUCS4ToUTF8(codepoint, dst)
// }

// [Sscanf] this works exactly like sscanf but doesn't require access to a C runtime.
//
// [Sscanf]: https://wiki.libsdl.org/SDL3/SDL_sscanf
// func Sscanf(text string, fmt string) int32 {
//	return sdlsscanf(text, fmt)
// }

// [Vsscanf] this works exactly like vsscanf but doesn't require access to a C.
//
// [Vsscanf]: https://wiki.libsdl.org/SDL3/SDL_vsscanf
// func Vsscanf(text string, fmt string, ap va_list) int32 {
//	return sdlvsscanf(text, fmt, ap)
// }

// [Snprintf] this works exactly like snprintf but doesn't require access to a C.
//
// [Snprintf]: https://wiki.libsdl.org/SDL3/SDL_snprintf
// func Snprintf(text string, maxlen uint64, fmt string) int32 {
//	return sdlsnprintf(text, maxlen, fmt)
// }

// [Swprintf] this works exactly like swprintf but doesn't require access to a C.
//
// [Swprintf]: https://wiki.libsdl.org/SDL3/SDL_swprintf
// func Swprintf(text *wchar_t, maxlen uint64, fmt *wchar_t) int32 {
//	return sdlswprintf(text, maxlen, fmt)
// }

// [Vsnprintf] this works exactly like vsnprintf but doesn't require access to a C.
//
// [Vsnprintf]: https://wiki.libsdl.org/SDL3/SDL_vsnprintf
// func Vsnprintf(text string, maxlen uint64, fmt string, ap va_list) int32 {
//	return sdlvsnprintf(text, maxlen, fmt, ap)
// }

// [Vswprintf] this works exactly like vswprintf but doesn't require access to a C.
//
// [Vswprintf]: https://wiki.libsdl.org/SDL3/SDL_vswprintf
// func Vswprintf(text *wchar_t, maxlen uint64, fmt *wchar_t, ap va_list) int32 {
//	return sdlvswprintf(text, maxlen, fmt, ap)
// }

// [Asprintf] this works exactly like asprintf but doesn't require access to a C.
//
// [Asprintf]: https://wiki.libsdl.org/SDL3/SDL_asprintf
// func Asprintf(strp **byte, fmt string) int32 {
//	return sdlasprintf(strp, fmt)
// }

// [Vasprintf] this works exactly like vasprintf but doesn't require access to a C.
//
// [Vasprintf]: https://wiki.libsdl.org/SDL3/SDL_vasprintf
// func Vasprintf(strp **byte, fmt string, ap va_list) int32 {
//	return sdlvasprintf(strp, fmt, ap)
// }

// [Srand] seeds the pseudo-random number generator.
//
// [Srand]: https://wiki.libsdl.org/SDL3/SDL_srand
// func Srand(seed uint64)  {
//	sdlsrand(seed)
// }

// [Rand] generates a pseudo-random number less than n for positive n.
//
// [Rand]: https://wiki.libsdl.org/SDL3/SDL_rand
// func Rand(n int32) int32 {
//	return sdlrand(n)
// }

// [Randf] generates a uniform pseudo-random floating point number less than 1.0.
//
// [Randf]: https://wiki.libsdl.org/SDL3/SDL_randf
// func Randf() float32 {
//	return sdlrandf()
// }

// [RandBits] generates 32 pseudo-random bits.
//
// [RandBits]: https://wiki.libsdl.org/SDL3/SDL_rand_bits
// func RandBits() uint32 {
//	return sdlrand_bits()
// }

// [RandR] generates a pseudo-random number less than n for positive n.
//
// [RandR]: https://wiki.libsdl.org/SDL3/SDL_rand_r
// func RandR(state *uint64, n int32) int32 {
//	return sdlrand_r(state, n)
// }

// [RandfR] generates a uniform pseudo-random floating point number less than 1.0.
//
// [RandfR]: https://wiki.libsdl.org/SDL3/SDL_randf_r
// func RandfR(state *uint64) float32 {
//	return sdlrandf_r(state)
// }

// [RandBitsR] generates 32 pseudo-random bits.
//
// [RandBitsR]: https://wiki.libsdl.org/SDL3/SDL_rand_bits_r
// func RandBitsR(state *uint64) uint32 {
//	return sdlrand_bits_r(state)
// }

// [Acos] computes the arc cosine of `x`.
//
// [Acos]: https://wiki.libsdl.org/SDL3/SDL_acos
// func Acos(x float64) float64 {
//	return sdlacos(x)
// }

// [Acosf] computes the arc cosine of `x`.
//
// [Acosf]: https://wiki.libsdl.org/SDL3/SDL_acosf
// func Acosf(x float32) float32 {
//	return sdlacosf(x)
// }

// [Asin] computes the arc sine of `x`.
//
// [Asin]: https://wiki.libsdl.org/SDL3/SDL_asin
// func Asin(x float64) float64 {
//	return sdlasin(x)
// }

// [Asinf] computes the arc sine of `x`.
//
// [Asinf]: https://wiki.libsdl.org/SDL3/SDL_asinf
// func Asinf(x float32) float32 {
//	return sdlasinf(x)
// }

// [Atan] computes the arc tangent of `x`.
//
// [Atan]: https://wiki.libsdl.org/SDL3/SDL_atan
// func Atan(x float64) float64 {
//	return sdlatan(x)
// }

// [Atanf] computes the arc tangent of `x`.
//
// [Atanf]: https://wiki.libsdl.org/SDL3/SDL_atanf
// func Atanf(x float32) float32 {
//	return sdlatanf(x)
// }

// [Atan2] computes the arc tangent of `y / x`, using the signs of x and y to adjust.
//
// [Atan2]: https://wiki.libsdl.org/SDL3/SDL_atan2
// func Atan2(y float64, x float64) float64 {
//	return sdlatan2(y, x)
// }

// [Atan2f] computes the arc tangent of `y / x`, using the signs of x and y to adjust.
//
// [Atan2f]: https://wiki.libsdl.org/SDL3/SDL_atan2f
// func Atan2f(y float32, x float32) float32 {
//	return sdlatan2f(y, x)
// }

// [Ceil] computes the ceiling of `x`.
//
// [Ceil]: https://wiki.libsdl.org/SDL3/SDL_ceil
// func Ceil(x float64) float64 {
//	return sdlceil(x)
// }

// [Ceilf] computes the ceiling of `x`.
//
// [Ceilf]: https://wiki.libsdl.org/SDL3/SDL_ceilf
// func Ceilf(x float32) float32 {
//	return sdlceilf(x)
// }

// [Copysign] copies the sign of one floating-point value to another.
//
// [Copysign]: https://wiki.libsdl.org/SDL3/SDL_copysign
// func Copysign(x float64, y float64) float64 {
//	return sdlcopysign(x, y)
// }

// [Copysignf] copies the sign of one floating-point value to another.
//
// [Copysignf]: https://wiki.libsdl.org/SDL3/SDL_copysignf
// func Copysignf(x float32, y float32) float32 {
//	return sdlcopysignf(x, y)
// }

// [Cos] computes the cosine of `x`.
//
// [Cos]: https://wiki.libsdl.org/SDL3/SDL_cos
// func Cos(x float64) float64 {
//	return sdlcos(x)
// }

// [Cosf] computes the cosine of `x`.
//
// [Cosf]: https://wiki.libsdl.org/SDL3/SDL_cosf
// func Cosf(x float32) float32 {
//	return sdlcosf(x)
// }

// [Exp] computes the exponential of `x`.
//
// [Exp]: https://wiki.libsdl.org/SDL3/SDL_exp
// func Exp(x float64) float64 {
//	return sdlexp(x)
// }

// [Expf] computes the exponential of `x`.
//
// [Expf]: https://wiki.libsdl.org/SDL3/SDL_expf
// func Expf(x float32) float32 {
//	return sdlexpf(x)
// }

// [Fabs] computes the absolute value of `x`.
//
// [Fabs]: https://wiki.libsdl.org/SDL3/SDL_fabs
// func Fabs(x float64) float64 {
//	return sdlfabs(x)
// }

// [Fabsf] computes the absolute value of `x`.
//
// [Fabsf]: https://wiki.libsdl.org/SDL3/SDL_fabsf
// func Fabsf(x float32) float32 {
//	return sdlfabsf(x)
// }

// [Floor] computes the floor of `x`.
//
// [Floor]: https://wiki.libsdl.org/SDL3/SDL_floor
// func Floor(x float64) float64 {
//	return sdlfloor(x)
// }

// [Floorf] computes the floor of `x`.
//
// [Floorf]: https://wiki.libsdl.org/SDL3/SDL_floorf
// func Floorf(x float32) float32 {
//	return sdlfloorf(x)
// }

// [Trunc] truncates `x` to an integer.
//
// [Trunc]: https://wiki.libsdl.org/SDL3/SDL_trunc
// func Trunc(x float64) float64 {
//	return sdltrunc(x)
// }

// [Truncf] truncates `x` to an integer.
//
// [Truncf]: https://wiki.libsdl.org/SDL3/SDL_truncf
// func Truncf(x float32) float32 {
//	return sdltruncf(x)
// }

// [Fmod] returns the floating-point remainder of `x / y`.
//
// [Fmod]: https://wiki.libsdl.org/SDL3/SDL_fmod
// func Fmod(x float64, y float64) float64 {
//	return sdlfmod(x, y)
// }

// [Fmodf] returns the floating-point remainder of `x / y`.
//
// [Fmodf]: https://wiki.libsdl.org/SDL3/SDL_fmodf
// func Fmodf(x float32, y float32) float32 {
//	return sdlfmodf(x, y)
// }

// [Isinf] returns whether the value is infinity.
//
// [Isinf]: https://wiki.libsdl.org/SDL3/SDL_isinf
// func Isinf(x float64) int32 {
//	return sdlisinf(x)
// }

// [Isinff] returns whether the value is infinity.
//
// [Isinff]: https://wiki.libsdl.org/SDL3/SDL_isinff
// func Isinff(x float32) int32 {
//	return sdlisinff(x)
// }

// [Isnan] returns whether the value is NaN.
//
// [Isnan]: https://wiki.libsdl.org/SDL3/SDL_isnan
// func Isnan(x float64) int32 {
//	return sdlisnan(x)
// }

// [Isnanf] returns whether the value is NaN.
//
// [Isnanf]: https://wiki.libsdl.org/SDL3/SDL_isnanf
// func Isnanf(x float32) int32 {
//	return sdlisnanf(x)
// }

// [Log] computes the natural logarithm of `x`.
//
// [Log]: https://wiki.libsdl.org/SDL3/SDL_log
// func Log(x float64) float64 {
//	return sdllog(x)
// }

// [Logf] computes the natural logarithm of `x`.
//
// [Logf]: https://wiki.libsdl.org/SDL3/SDL_logf
// func Logf(x float32) float32 {
//	return sdllogf(x)
// }

// [Log10] computes the base-10 logarithm of `x`.
//
// [Log10]: https://wiki.libsdl.org/SDL3/SDL_log10
// func Log10(x float64) float64 {
//	return sdllog10(x)
// }

// [Log10f] computes the base-10 logarithm of `x`.
//
// [Log10f]: https://wiki.libsdl.org/SDL3/SDL_log10f
// func Log10f(x float32) float32 {
//	return sdllog10f(x)
// }

// [Modf] splits `x` into integer and fractional parts.
//
// [Modf]: https://wiki.libsdl.org/SDL3/SDL_modf
// func Modf(x float64, y *:double) float64 {
//	return sdlmodf(x, y)
// }

// [Modff] splits `x` into integer and fractional parts.
//
// [Modff]: https://wiki.libsdl.org/SDL3/SDL_modff
// func Modff(x float32, y *float32) float32 {
//	return sdlmodff(x, y)
// }

// [Pow] raises `x` to the power `y`.
//
// [Pow]: https://wiki.libsdl.org/SDL3/SDL_pow
// func Pow(x float64, y float64) float64 {
//	return sdlpow(x, y)
// }

// [Powf] raises `x` to the power `y`.
//
// [Powf]: https://wiki.libsdl.org/SDL3/SDL_powf
// func Powf(x float32, y float32) float32 {
//	return sdlpowf(x, y)
// }

// [Round] rounds `x` to the nearest integer.
//
// [Round]: https://wiki.libsdl.org/SDL3/SDL_round
// func Round(x float64) float64 {
//	return sdlround(x)
// }

// [Roundf] rounds `x` to the nearest integer.
//
// [Roundf]: https://wiki.libsdl.org/SDL3/SDL_roundf
// func Roundf(x float32) float32 {
//	return sdlroundf(x)
// }

// [Lround] rounds `x` to the nearest integer representable as a long.
//
// [Lround]: https://wiki.libsdl.org/SDL3/SDL_lround
// func Lround(x float64) int64 {
//	return sdllround(x)
// }

// [Lroundf] rounds `x` to the nearest integer representable as a long.
//
// [Lroundf]: https://wiki.libsdl.org/SDL3/SDL_lroundf
// func Lroundf(x float32) int64 {
//	return sdllroundf(x)
// }

// [Scalbn] scales `x` by an integer power of two.
//
// [Scalbn]: https://wiki.libsdl.org/SDL3/SDL_scalbn
// func Scalbn(x float64, n int32) float64 {
//	return sdlscalbn(x, n)
// }

// [Scalbnf] scales `x` by an integer power of two.
//
// [Scalbnf]: https://wiki.libsdl.org/SDL3/SDL_scalbnf
// func Scalbnf(x float32, n int32) float32 {
//	return sdlscalbnf(x, n)
// }

// [Sin] computes the sine of `x`.
//
// [Sin]: https://wiki.libsdl.org/SDL3/SDL_sin
// func Sin(x float64) float64 {
//	return sdlsin(x)
// }

// [Sinf] computes the sine of `x`.
//
// [Sinf]: https://wiki.libsdl.org/SDL3/SDL_sinf
// func Sinf(x float32) float32 {
//	return sdlsinf(x)
// }

// [Sqrt] computes the square root of `x`.
//
// [Sqrt]: https://wiki.libsdl.org/SDL3/SDL_sqrt
// func Sqrt(x float64) float64 {
//	return sdlsqrt(x)
// }

// [Sqrtf] computes the square root of `x`.
//
// [Sqrtf]: https://wiki.libsdl.org/SDL3/SDL_sqrtf
// func Sqrtf(x float32) float32 {
//	return sdlsqrtf(x)
// }

// [Tan] computes the tangent of `x`.
//
// [Tan]: https://wiki.libsdl.org/SDL3/SDL_tan
// func Tan(x float64) float64 {
//	return sdltan(x)
// }

// [Tanf] computes the tangent of `x`.
//
// [Tanf]: https://wiki.libsdl.org/SDL3/SDL_tanf
// func Tanf(x float32) float32 {
//	return sdltanf(x)
// }

// [IconvT] is an opaque handle representing string encoding conversion state.
//
// [IconvT]: https://wiki.libsdl.org/SDL3/SDL_iconv_t
type IconvT struct{}

// [IconvOpen] this function allocates a context for the specified character set.
//
// [IconvOpen]: https://wiki.libsdl.org/SDL3/SDL_iconv_open
// func IconvOpen(tocode string, fromcode string) iconv_t {
//	return sdliconv_open(tocode, fromcode)
// }

// [IconvClose] this function frees a context used for character set conversion.
//
// [IconvClose]: https://wiki.libsdl.org/SDL3/SDL_iconv_close
// func IconvClose(cd iconv_t) int32 {
//	return sdliconv_close(cd)
// }

// [Iconv] this function converts text between encodings, reading from and writing to.
//
// [Iconv]: https://wiki.libsdl.org/SDL3/SDL_iconv
// func Iconv(cd iconv_t, inbuf **byte, inbytesleft *uint64, outbuf **byte, outbytesleft *uint64) uint64 {
//	return sdliconv(cd, inbuf, inbytesleft, outbuf, outbytesleft)
// }

// [IconvString] helpers function to convert a string's encoding in one call.
//
// [IconvString]: https://wiki.libsdl.org/SDL3/SDL_iconv_string
// func IconvString(tocode string, fromcode string, inbuf string, inbytesleft uint64) string {
//	return sdliconv_string(tocode, fromcode, inbuf, inbytesleft)
// }

// [SizeMulCheckOverflow] multiplys two integers, checking for overflow.
//
// [SizeMulCheckOverflow]: https://wiki.libsdl.org/SDL3/SDL_size_mul_check_overflow
// func SizeMulCheckOverflow(a uint64, b uint64, ret *uint64) bool {
//	return sdlsize_mul_check_overflow(a, b, ret)
// }

// [SizeMulCheckOverflowBuiltin]
//
// [SizeMulCheckOverflowBuiltin]: https://wiki.libsdl.org/SDL3/SDL_size_mul_check_overflow_builtin
// func SizeMulCheckOverflowBuiltin(a uint64, b uint64, ret *uint64) bool {
//	return sdlsize_mul_check_overflow_builtin(a, b, ret)
// }

// [SizeAddCheckOverflow] adds two integers, checking for overflow.
//
// [SizeAddCheckOverflow]: https://wiki.libsdl.org/SDL3/SDL_size_add_check_overflow
// func SizeAddCheckOverflow(a uint64, b uint64, ret *uint64) bool {
//	return sdlsize_add_check_overflow(a, b, ret)
// }

// [SizeAddCheckOverflowBuiltin]
//
// [SizeAddCheckOverflowBuiltin]: https://wiki.libsdl.org/SDL3/SDL_size_add_check_overflow_builtin
// func SizeAddCheckOverflowBuiltin(a uint64, b uint64, ret *uint64) bool {
//	return sdlsize_add_check_overflow_builtin(a, b, ret)
// }

// [FunctionPointer] defines a generic function pointer.
//
// [FunctionPointer]: https://wiki.libsdl.org/SDL3/SDL_FunctionPointer
type FunctionPointer uintptr
