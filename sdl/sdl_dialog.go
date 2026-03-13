package sdl

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/jupiterrider/purego-sdl3/internal/convert"
)

const (
	PropFileDialogFiltersPointer = "SDL.filedialog.filters"
	PropFileDialogNfiltersNumber = "SDL.filedialog.nfilters"
	PropFileDialogWindowPointer  = "SDL.filedialog.window"
	PropFileDialogLocationString = "SDL.filedialog.location"
	PropFileDialogManyBoolean    = "SDL.filedialog.many"
	PropFileDialogTitleString    = "SDL.filedialog.title"
	PropFileDialogAcceptString   = "SDL.filedialog.accept"
	PropFileDialogCancelString   = "SDL.filedialog.cancel"
)

// [FileDialogType] defines the various types of file dialogs.
//
// [FileDialogType]: https://wiki.libsdl.org/SDL3/SDL_FileDialogType
type FileDialogType uint32

const (
	FileDialogOpenFile FileDialogType = iota
	FileDialogSaveFile
	FileDialogOpenFolder
)

// [DialogFileFilter] is an entry for filters for file dialogs.
//
// [DialogFileFilter]: https://wiki.libsdl.org/SDL3/SDL_DialogFileFilter
type DialogFileFilter struct {
	name, pattern *byte
}

func NewDialogFileFilter(name, pattern string) DialogFileFilter {
	return DialogFileFilter{convert.ToBytePtr(name), convert.ToBytePtr(pattern)}
}

func (d DialogFileFilter) Name() string {
	return convert.ToString(d.name)
}

func (d DialogFileFilter) Pattern() string {
	return convert.ToString(d.pattern)
}

// [DialogFileCallback] defines the callback used by file dialog functions.
//
// [DialogFileCallback]: https://wiki.libsdl.org/SDL3/SDL_DialogFileCallback
type DialogFileCallback uintptr

func NewDialogFileCallback(callback func(userdata unsafe.Pointer, filelist []string, filter int32)) DialogFileCallback {
	cb := purego.NewCallback(func(userdata unsafe.Pointer, filelist **byte, filter int32) uintptr {
		callback(userdata, convert.ToStringSlice(filelist), filter)
		return 0
	})

	return DialogFileCallback(cb)
}

// [ShowFileDialogWithProperties] creates and launches a file dialog with the specified properties.
//
// [ShowFileDialogWithProperties]: https://wiki.libsdl.org/SDL3/SDL_ShowFileDialogWithProperties
func ShowFileDialogWithProperties(dialogType FileDialogType, callback DialogFileCallback, userdata unsafe.Pointer, props PropertiesID) {
	sdlShowFileDialogWithProperties(dialogType, callback, userdata, props)
}

// [ShowOpenFileDialog] displays a dialog that lets the user select a file on their filesystem.
//
// [ShowOpenFileDialog]: https://wiki.libsdl.org/SDL3/SDL_ShowOpenFileDialog
func ShowOpenFileDialog(callback DialogFileCallback, userdata unsafe.Pointer, window *Window, filters []DialogFileFilter, defaultLocation string, allowMany bool) {
	sdlShowOpenFileDialog(callback, userdata, window, filters, int32(len(filters)), convert.ToBytePtrNullable(defaultLocation), allowMany)
}

// [ShowOpenFolderDialog] displays a dialog that lets the user select a folder on their filesystem.
//
// [ShowOpenFolderDialog]: https://wiki.libsdl.org/SDL3/SDL_ShowOpenFolderDialog
func ShowOpenFolderDialog(callback DialogFileCallback, userdata unsafe.Pointer, window *Window, defaultLocation string, allowMany bool) {
	sdlShowOpenFolderDialog(callback, userdata, window, defaultLocation, allowMany)
}

// [ShowSaveFileDialog] displays a dialog that lets the user choose a new or existing file on their filesystem.
//
// [ShowSaveFileDialog]: https://wiki.libsdl.org/SDL3/SDL_ShowSaveFileDialog
func ShowSaveFileDialog(callback DialogFileCallback, userdata unsafe.Pointer, window *Window, filters []DialogFileFilter, defaultLocation string) {
	sdlShowSaveFileDialog(callback, userdata, window, filters, int32(len(filters)), defaultLocation)
}
