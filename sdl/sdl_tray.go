package sdl

// [Tray] is an opaque handle representing a toplevel system tray object.
//
// [Tray]: https://wiki.libsdl.org/SDL3/SDL_Tray
type Tray struct{}

// [TrayMenu] is an opaque handle representing a menu/submenu on a system tray object.
//
// [TrayMenu]: https://wiki.libsdl.org/SDL3/SDL_TrayMenu
type TrayMenu struct{}

// [TrayEntry] is an opaque handle representing an entry on a system tray object.
//
// [TrayEntry]: https://wiki.libsdl.org/SDL3/SDL_TrayEntry
type TrayEntry struct{}

// [TrayEntryFlags] defines flags that control the creation of system tray entries.
//
// [TrayEntryFlags]: https://wiki.libsdl.org/SDL3/SDL_TrayEntryFlags
type TrayEntryFlags uint32

const (
	TrayentryButton   TrayEntryFlags = 0x00000001 // Make the entry a simple button. Required.
	TrayentryCheckbox TrayEntryFlags = 0x00000002 // Make the entry a checkbox. Required.
	TrayentrySubmenu  TrayEntryFlags = 0x00000004 // Prepare the entry to have a submenu. Required
	TrayentryDisabled TrayEntryFlags = 0x80000000 // Make the entry disabled. Optional.
	TrayentryChecked  TrayEntryFlags = 0x40000000 // Make the entry checked. This is valid only for checkboxes. Optional.
)

// func ClickTrayEntry(entry *TrayEntry)  {
//	sdlClickTrayEntry(entry)
// }

// func CreateTray(icon *Surface, tooltip string) *Tray {
//	return sdlCreateTray(icon, tooltip)
// }

// func CreateTrayMenu(tray *Tray) *TrayMenu {
//	return sdlCreateTrayMenu(tray)
// }

// func CreateTraySubmenu(entry *TrayEntry) *TrayMenu {
//	return sdlCreateTraySubmenu(entry)
// }

// func DestroyTray(tray *Tray)  {
//	sdlDestroyTray(tray)
// }

// func GetTrayEntries(menu *TrayMenu, size *int32) **TrayEntry {
//	return sdlGetTrayEntries(menu, size)
// }

// func GetTrayEntryChecked(entry *TrayEntry) bool {
//	return sdlGetTrayEntryChecked(entry)
// }

// func GetTrayEntryEnabled(entry *TrayEntry) bool {
//	return sdlGetTrayEntryEnabled(entry)
// }

// func GetTrayEntryLabel(entry *TrayEntry) string {
//	return sdlGetTrayEntryLabel(entry)
// }

// func GetTrayEntryParent(entry *TrayEntry) *TrayMenu {
//	return sdlGetTrayEntryParent(entry)
// }

// func GetTrayMenu(tray *Tray) *TrayMenu {
//	return sdlGetTrayMenu(tray)
// }

// func GetTrayMenuParentEntry(menu *TrayMenu) *TrayEntry {
//	return sdlGetTrayMenuParentEntry(menu)
// }

// func GetTrayMenuParentTray(menu *TrayMenu) *Tray {
//	return sdlGetTrayMenuParentTray(menu)
// }

// func GetTraySubmenu(entry *TrayEntry) *TrayMenu {
//	return sdlGetTraySubmenu(entry)
// }

// func InsertTrayEntryAt(menu *TrayMenu, pos int32, label string, flags TrayEntryFlags) *TrayEntry {
//	return sdlInsertTrayEntryAt(menu, pos, label, flags)
// }

// func RemoveTrayEntry(entry *TrayEntry)  {
//	sdlRemoveTrayEntry(entry)
// }

// func SetTrayEntryCallback(entry *TrayEntry, callback TrayCallback, userdata unsafe.Pointer)  {
//	sdlSetTrayEntryCallback(entry, callback, userdata)
// }

// func SetTrayEntryChecked(entry *TrayEntry, checked bool)  {
//	sdlSetTrayEntryChecked(entry, checked)
// }

// func SetTrayEntryEnabled(entry *TrayEntry, enabled bool)  {
//	sdlSetTrayEntryEnabled(entry, enabled)
// }

// func SetTrayEntryLabel(entry *TrayEntry, label string)  {
//	sdlSetTrayEntryLabel(entry, label)
// }

// func SetTrayIcon(tray *Tray, icon *Surface)  {
//	sdlSetTrayIcon(tray, icon)
// }

// func SetTrayTooltip(tray *Tray, tooltip string)  {
//	sdlSetTrayTooltip(tray, tooltip)
// }

// [UpdateTrays] updates the trays.
//
// [UpdateTrays]: https://wiki.libsdl.org/SDL3/SDL_UpdateTrays
// func UpdateTrays() {
// 	sdlUpdateTrays()
// }
