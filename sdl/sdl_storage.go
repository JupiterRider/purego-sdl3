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

// func CloseStorage(storage *Storage) bool {
//	return sdlCloseStorage(storage)
// }

// func CopyStorageFile(storage *Storage, oldpath string, newpath string) bool {
//	return sdlCopyStorageFile(storage, oldpath, newpath)
// }

// func CreateStorageDirectory(storage *Storage, path string) bool {
//	return sdlCreateStorageDirectory(storage, path)
// }

// func EnumerateStorageDirectory(storage *Storage, path string, callback EnumerateDirectoryCallback, userdata unsafe.Pointer) bool {
//	return sdlEnumerateStorageDirectory(storage, path, callback, userdata)
// }

// func GetStorageFileSize(storage *Storage, path string, length *uint64) bool {
//	return sdlGetStorageFileSize(storage, path, length)
// }

// func GetStoragePathInfo(storage *Storage, path string, info *PathInfo) bool {
//	return sdlGetStoragePathInfo(storage, path, info)
// }

// func GetStorageSpaceRemaining(storage *Storage) uint64 {
//	return sdlGetStorageSpaceRemaining(storage)
// }

// func GlobStorageDirectory(storage *Storage, path string, pattern string, flags GlobFlags, count *int32) **byte {
//	return sdlGlobStorageDirectory(storage, path, pattern, flags, count)
// }

// func OpenFileStorage(path string) *Storage {
//	return sdlOpenFileStorage(path)
// }

// func OpenStorage(iface *StorageInterface, userdata unsafe.Pointer) *Storage {
//	return sdlOpenStorage(iface, userdata)
// }

// func OpenTitleStorage(override string, props PropertiesID) *Storage {
//	return sdlOpenTitleStorage(override, props)
// }

// func OpenUserStorage(org string, app string, props PropertiesID) *Storage {
//	return sdlOpenUserStorage(org, app, props)
// }

// func ReadStorageFile(storage *Storage, path string, destination unsafe.Pointer, length uint64) bool {
//	return sdlReadStorageFile(storage, path, destination, length)
// }

// func RemoveStoragePath(storage *Storage, path string) bool {
//	return sdlRemoveStoragePath(storage, path)
// }

// func RenameStoragePath(storage *Storage, oldpath string, newpath string) bool {
//	return sdlRenameStoragePath(storage, oldpath, newpath)
// }

// func StorageReady(storage *Storage) bool {
//	return sdlStorageReady(storage)
// }

// func WriteStorageFile(storage *Storage, path string, source unsafe.Pointer, length uint64) bool {
//	return sdlWriteStorageFile(storage, path, source, length)
// }
