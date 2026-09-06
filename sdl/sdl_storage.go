package sdl

// [StorageInterface] defines function interface for [Storage].
//
// [StorageInterface]: https://wiki.libsdl.org/SDL3/SDL_StorageInterface
type StorageInterface struct {
	Version        uint32                                                                                                  // The version of this interface.
	Close          *func(userdata uintptr) bool                                                                            // Called when the storage is closed.
	Ready          *func(userdata uintptr) bool                                                                            // Optional, returns whether the storage is currently ready for access.
	Enumerate      *func(userdata uintptr, path *byte, callback EnumerateDirectoryCallback, callbackUserdata uintptr) bool // Enumerate a directory, optional for write-only storage.
	Info           *func(userdata uintptr, path *byte, info *PathInfo) bool                                                // Get path information, optional for write-only storage.
	ReadFile       *func(userdata uintptr, path *byte, destination uintptr, length uint64) bool                            // Read a file from storage, optional for write-only storage.
	WriteFile      *func(userdata uintptr, path *byte, source uintptr, length uint64) bool                                 // Write a file to storage, optional for read-only storage.
	Mkdir          *func(userdata uintptr, path *byte) bool                                                                // Create a directory, optional for read-only storage.
	Remove         *func(userdata uintptr, path *byte) bool                                                                // Remove a file or empty directory, optional for read-only storage.
	Rename         *func(userdata uintptr, oldpath, newpath *byte) bool                                                    // Rename a path, optional for read-only storage.
	Copy           *func(userdata uintptr, oldpath, newpath *byte) bool                                                    // Copy a file, optional for read-only storage.
	SpaceRemaining *func(userdata uintptr) uint64                                                                          // Get the space remaining, optional for read-only storage.
}

// [Storage] is an abstract interface for filesystem access.
//
// [Storage]: https://wiki.libsdl.org/SDL3/SDL_Storage
type Storage struct{}

// [OpenTitleStorage] opens up a read-only container for the application's filesystem.
//
// [OpenTitleStorage]: https://wiki.libsdl.org/SDL3/SDL_OpenTitleStorage
// func OpenTitleStorage(override string, props PropertiesID) *Storage {
//	return sdlOpenTitleStorage(override, props)
// }

// [OpenUserStorage] opens up a container for a user's unique read/write filesystem.
//
// [OpenUserStorage]: https://wiki.libsdl.org/SDL3/SDL_OpenUserStorage
// func OpenUserStorage(org string, app string, props PropertiesID) *Storage {
//	return sdlOpenUserStorage(org, app, props)
// }

// [OpenFileStorage] opens up a container for local filesystem storage.
//
// [OpenFileStorage]: https://wiki.libsdl.org/SDL3/SDL_OpenFileStorage
// func OpenFileStorage(path string) *Storage {
//	return sdlOpenFileStorage(path)
// }

// [OpenStorage] opens up a container using a client-provided storage interface.
//
// [OpenStorage]: https://wiki.libsdl.org/SDL3/SDL_OpenStorage
// func OpenStorage(iface *StorageInterface, userdata unsafe.Pointer) *Storage {
//	return sdlOpenStorage(iface, userdata)
// }

// [CloseStorage] closes and frees a storage container.
//
// [CloseStorage]: https://wiki.libsdl.org/SDL3/SDL_CloseStorage
// func CloseStorage(storage *Storage) bool {
//	return sdlCloseStorage(storage)
// }

// [StorageReady] checks if the storage container is ready to use.
//
// [StorageReady]: https://wiki.libsdl.org/SDL3/SDL_StorageReady
// func StorageReady(storage *Storage) bool {
//	return sdlStorageReady(storage)
// }

// [GetStorageFileSize] queries the size of a file within a storage container.
//
// [GetStorageFileSize]: https://wiki.libsdl.org/SDL3/SDL_GetStorageFileSize
// func GetStorageFileSize(storage *Storage, path string, length *uint64) bool {
//	return sdlGetStorageFileSize(storage, path, length)
// }

// [ReadStorageFile] synchronouslys read a file from a storage container into a client-provided.
//
// [ReadStorageFile]: https://wiki.libsdl.org/SDL3/SDL_ReadStorageFile
// func ReadStorageFile(storage *Storage, path string, destination unsafe.Pointer, length uint64) bool {
//	return sdlReadStorageFile(storage, path, destination, length)
// }

// [WriteStorageFile] synchronouslys write a file from client memory into a storage container.
//
// [WriteStorageFile]: https://wiki.libsdl.org/SDL3/SDL_WriteStorageFile
// func WriteStorageFile(storage *Storage, path string, source unsafe.Pointer, length uint64) bool {
//	return sdlWriteStorageFile(storage, path, source, length)
// }

// [CreateStorageDirectory] creates a directory in a writable storage container.
//
// [CreateStorageDirectory]: https://wiki.libsdl.org/SDL3/SDL_CreateStorageDirectory
// func CreateStorageDirectory(storage *Storage, path string) bool {
//	return sdlCreateStorageDirectory(storage, path)
// }

// [EnumerateStorageDirectory] enumerates a directory in a storage container through a callback function.
//
// [EnumerateStorageDirectory]: https://wiki.libsdl.org/SDL3/SDL_EnumerateStorageDirectory
// func EnumerateStorageDirectory(storage *Storage, path string, callback EnumerateDirectoryCallback, userdata unsafe.Pointer) bool {
//	return sdlEnumerateStorageDirectory(storage, path, callback, userdata)
// }

// [RemoveStoragePath] removes a file or an empty directory in a writable storage container.
//
// [RemoveStoragePath]: https://wiki.libsdl.org/SDL3/SDL_RemoveStoragePath
// func RemoveStoragePath(storage *Storage, path string) bool {
//	return sdlRemoveStoragePath(storage, path)
// }

// [RenameStoragePath] renames a file or directory in a writable storage container.
//
// [RenameStoragePath]: https://wiki.libsdl.org/SDL3/SDL_RenameStoragePath
// func RenameStoragePath(storage *Storage, oldpath string, newpath string) bool {
//	return sdlRenameStoragePath(storage, oldpath, newpath)
// }

// [CopyStorageFile] copies a file in a writable storage container.
//
// [CopyStorageFile]: https://wiki.libsdl.org/SDL3/SDL_CopyStorageFile
// func CopyStorageFile(storage *Storage, oldpath string, newpath string) bool {
//	return sdlCopyStorageFile(storage, oldpath, newpath)
// }

// [GetStoragePathInfo] gets information about a filesystem path in a storage container.
//
// [GetStoragePathInfo]: https://wiki.libsdl.org/SDL3/SDL_GetStoragePathInfo
// func GetStoragePathInfo(storage *Storage, path string, info *PathInfo) bool {
//	return sdlGetStoragePathInfo(storage, path, info)
// }

// [GetStorageSpaceRemaining] queries the remaining space in a storage container.
//
// [GetStorageSpaceRemaining]: https://wiki.libsdl.org/SDL3/SDL_GetStorageSpaceRemaining
// func GetStorageSpaceRemaining(storage *Storage) uint64 {
//	return sdlGetStorageSpaceRemaining(storage)
// }

// [GlobStorageDirectory] enumerates a directory tree, filtered by pattern, and returns a list.
//
// [GlobStorageDirectory]: https://wiki.libsdl.org/SDL3/SDL_GlobStorageDirectory
// func GlobStorageDirectory(storage *Storage, path string, pattern string, flags GlobFlags, count *int32) **byte {
//	return sdlGlobStorageDirectory(storage, path, pattern, flags, count)
// }
