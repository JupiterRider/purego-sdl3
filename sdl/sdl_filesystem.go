package sdl

// [EnumerationResult] defines the possible results from an enumeration callback.
//
// [EnumerationResult]: https://wiki.libsdl.org/SDL3/SDL_EnumerationResult
type EnumerationResult uint32

const (
	EnumContinue EnumerationResult = iota // Value that requests that enumeration continue.
	EnumSuccess                           // Value that requests that enumeration stop, successfully.
	EnumFailure                           // Value that requests that enumeration stop, as a failure.
)

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

// [Folder] is a structure specifying the type of the OS-provided default folder for a specific purposes.
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

// func CopyFile(oldpath string, newpath string) bool {
//	return sdlCopyFile(oldpath, newpath)
// }

// func CreateDirectory(path string) bool {
//	return sdlCreateDirectory(path)
// }

// func EnumerateDirectory(path string, callback EnumerateDirectoryCallback, userdata unsafe.Pointer) bool {
//	return sdlEnumerateDirectory(path, callback, userdata)
// }

// [GetBasePath] gets the directory where the application was run from.
//
// [GetBasePath]: https://wiki.libsdl.org/SDL3/SDL_GetBasePath
func GetBasePath() string {
	return sdlGetBasePath()
}

// func GetCurrentDirectory() string {
//	return sdlGetCurrentDirectory()
// }

// func GetPathInfo(path string, info *PathInfo) bool {
//	return sdlGetPathInfo(path, info)
// }

// func GetPrefPath(org string, app string) string {
//	return sdlGetPrefPath(org, app)
// }

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
