package sdl

import (
	"unsafe"

	"github.com/abev08/purego-sdl3/internal/convert"
)

// [GetBasePath] gets the directory where the application was run from.
//
// [GetBasePath]: https://wiki.libsdl.org/SDL3/SDL_GetBasePath
func GetBasePath() string {
	return sdlGetBasePath()
}

// [GetPrefPath] gets the user-and-app-specific path where files can be written.
//
// [GetPrefPath]: https://wiki.libsdl.org/SDL3/SDL_GetPrefPath
func GetPrefPath(org string, app string) string {
	ret := sdlGetPrefPath(convert.ToBytePtrNullable(org), convert.ToBytePtrNullable(app))
	defer Free(unsafe.Pointer(ret))
	return convert.ToString(ret)
}

// [Folder] defines the type of the OS-provided default folder for a specific purpose.
//
// [Folder]: https://wiki.libsdl.org/SDL3/SDL_Folder
type Folder uint32

const (
	FolderHome        Folder = iota // The folder which contains all of the current user's data, preferences, and documents. It usually contains most of the other folders. If a requested folder does not exist, the home folder can be considered a safe fallback to store a user's documents.
	FolderDesktop                   // The folder of files that are displayed on the desktop. Note that the existence of a desktop folder does not guarantee that the system does show icons on its desktop; certain GNU/Linux distros with a graphical environment may not have desktop icons.
	FolderDocuments                 // User document files, possibly application-specific. This is a good place to save a user's projects.
	FolderDownloads                 // Standard folder for user files downloaded from the internet.
	FolderMusic                     // Music files that can be played using a standard music player (mp3, ogg...).
	FolderPictures                  // Image files that can be displayed using a standard viewer (png, jpg...).
	FolderPublicShare               // Files that are meant to be shared with other users on the same computer.
	FolderSavedGames                // Save files for games.
	FolderScreenshots               // Application screenshots.
	FolderTemplates                 // Template files to be used when the user requests the desktop environment to create a new file in a certain folder, such as "New Text File.txt". Any file in the Templates folder can be used as a starting point for a new file.
	FolderVideos                    // Video files that can be played using a standard video player (mp4, webm...).
	FolderCount                     // Total number of types in this enum, not a folder type by itself.
)

// [GetUserFolder] finds the most suitable user folder for a specific purpose.
//
// [GetUserFolder]: https://wiki.libsdl.org/SDL3/SDL_GetUserFolder
// func GetUserFolder(folder Folder) string {
//	return sdlGetUserFolder(folder)
// }

// [PathType] defines the types of filesystem entries.
//
// [PathType]: https://wiki.libsdl.org/SDL3/SDL_PathType
type PathType uint32

const (
	PathTypeNone      PathType = iota // Path does not exist.
	PathTypeFile                      // A normal file.
	PathTypeDirectory                 // A directory.
	PathTypeOther                     // Something completely different like a device node (not a symlink, those are always followed).
)

// [PathInfo] defines the information about a path on the filesystem.
//
// [PathInfo]: https://wiki.libsdl.org/SDL3/SDL_PathInfo
type PathInfo struct {
	Type       PathType // The path type.
	Size       uint64   // The file size in bytes.
	CreateTime Time     // The time when the path was created.
	ModifyTime Time     // The last time the path was modified.
	AccessTime Time     // The last time the path was read.
}

// [GlobFlags] defines flags for path matching.
//
// [GlobFlags]: https://wiki.libsdl.org/SDL3/SDL_GlobFlags
type GlobFlags uint32

const GlobCaseinsensitive GlobFlags = 1 << 0

// [CreateDirectory] creates a directory, and any missing parent directories.
//
// [CreateDirectory]: https://wiki.libsdl.org/SDL3/SDL_CreateDirectory
// func CreateDirectory(path string) bool {
//	return sdlCreateDirectory(path)
// }

// [EnumerationResult] defines the possible results from an enumeration callback.
//
// [EnumerationResult]: https://wiki.libsdl.org/SDL3/SDL_EnumerationResult
type EnumerationResult uint32

const (
	EnumContinue EnumerationResult = iota // Value that requests that enumeration continue.
	EnumSuccess                           // Value that requests that enumeration stop, successfully.
	EnumFailure                           // Value that requests that enumeration stop, as a failure.
)

// [EnumerateDirectoryCallback] defines callback for directory enumeration.
//
// [EnumerateDirectoryCallback]: https://wiki.libsdl.org/SDL3/SDL_EnumerateDirectoryCallback
type EnumerateDirectoryCallback uintptr

// [EnumerateDirectory] enumerates a directory through a callback function.
//
// [EnumerateDirectory]: https://wiki.libsdl.org/SDL3/SDL_EnumerateDirectory
// func EnumerateDirectory(path string, callback EnumerateDirectoryCallback, userdata unsafe.Pointer) bool {
//	return sdlEnumerateDirectory(path, callback, userdata)
// }

// [RemovePath] removes a file or an empty directory.
//
// [RemovePath]: https://wiki.libsdl.org/SDL3/SDL_RemovePath
// func RemovePath(path string) bool {
//	return sdlRemovePath(path)
// }

// [RenamePath] renames a file or directory.
//
// [RenamePath]: https://wiki.libsdl.org/SDL3/SDL_RenamePath
// func RenamePath(oldpath string, newpath string) bool {
//	return sdlRenamePath(oldpath, newpath)
// }

// [CopyFile] copies a file.
//
// [CopyFile]: https://wiki.libsdl.org/SDL3/SDL_CopyFile
// func CopyFile(oldpath string, newpath string) bool {
//	return sdlCopyFile(oldpath, newpath)
// }

// [GetPathInfo] gets information about a filesystem path.
//
// [GetPathInfo]: https://wiki.libsdl.org/SDL3/SDL_GetPathInfo
// func GetPathInfo(path string, info *PathInfo) bool {
//	return sdlGetPathInfo(path, info)
// }

// [GlobDirectory] enumerates a directory tree, filtered by pattern, and return a list.
//
// [GlobDirectory]: https://wiki.libsdl.org/SDL3/SDL_GlobDirectory
// func GlobDirectory(path string, pattern string, flags GlobFlags, count *int32) **byte {
//	return sdlGlobDirectory(path, pattern, flags, count)
// }

// [GetCurrentDirectory] gets what the system believes is the "current working directory".
//
// [GetCurrentDirectory]: https://wiki.libsdl.org/SDL3/SDL_GetCurrentDirectory
func GetCurrentDirectory() string {
	ret := sdlGetCurrentDirectory()
	defer Free(unsafe.Pointer(ret))
	return convert.ToString(ret)
}
