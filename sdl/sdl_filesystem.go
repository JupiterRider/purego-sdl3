package sdl

import (
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

type EnumerationResult uint32

const (
	EnumContinue EnumerationResult = iota
	EnumSuccess
	EnumFailure
)

type PathType uint32

const (
	PathTypeNone PathType = iota
	PathTypeFile
	PathTypeDirectory
	PathTypeOther
)

type Folder uint32

const (
	FolderHome Folder = iota
	FolderDesktop
	FolderDocuments
	FolderDownloads
	FolderMusic
	FolderPictures
	FolderPublicShare
	FolderSavedGames
	FolderScreenshots
	FolderTemplates
	FolderVideos
	FolderCount
)

// func CopyFile(oldpath string, newpath string) bool {
//	return sdlCopyFile(oldpath, newpath)
// }

// func CreateDirectory(path string) bool {
//	return sdlCreateDirectory(path)
// }

// func EnumerateDirectory(path string, callback EnumerateDirectoryCallback, userdata unsafe.Pointer) bool {
//	return sdlEnumerateDirectory(path, callback, userdata)
// }

// GetBasePath return the directory where the application was run from.
func GetBasePath() string {
	return sdlGetBasePath()
}

func GetCurrentDirectory() string {
	ret := sdlGetCurrentDirectory()
	defer Free(unsafe.Pointer(ret))
	return convert.ToString(ret)
}

// func GetPathInfo(path string, info *PathInfo) bool {
//	return sdlGetPathInfo(path, info)
// }

func GetPrefPath(org string, app string) string {
	ret := sdlGetPrefPath(convert.ToBytePtrNullable(org), convert.ToBytePtrNullable(app))
	defer Free(unsafe.Pointer(ret))
	return convert.ToString(ret)
}

// func GetUserFolder(folder Folder) string {
//	return sdlGetUserFolder(folder)
// }

// func GlobDirectory(path string, pattern string, flags GlobFlags, count *int32) **byte {
//	return sdlGlobDirectory(path, pattern, flags, count)
// }

// func RemovePath(path string) bool {
//	return sdlRemovePath(path)
// }

// func RenamePath(oldpath string, newpath string) bool {
//	return sdlRenamePath(oldpath, newpath)
// }
