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

// [TrayCallback] defines a callback that is invoked when a tray entry is selected.
//
// [TrayCallback]: https://wiki.libsdl.org/SDL3/SDL_TrayCallback
type TrayCallback uintptr

// [CreateTray] creates an icon to be placed in the operating system's tray, or equivalent.
//
// [CreateTray]: https://wiki.libsdl.org/SDL3/SDL_CreateTray
// func CreateTray(icon *Surface, tooltip string) *Tray {
//	return sdlCreateTray(icon, tooltip)
// }

// [SetTrayIcon] updates the system tray icon's icon.
//
// [SetTrayIcon]: https://wiki.libsdl.org/SDL3/SDL_SetTrayIcon
// func SetTrayIcon(tray *Tray, icon *Surface) {
//	sdlSetTrayIcon(tray, icon)
// }

// [SetTrayTooltip] updates the system tray icon's tooltip.
//
// [SetTrayTooltip]: https://wiki.libsdl.org/SDL3/SDL_SetTrayTooltip
// func SetTrayTooltip(tray *Tray, tooltip string) {
//	sdlSetTrayTooltip(tray, tooltip)
// }

// [CreateTrayMenu] creates a menu for a system tray.
//
// [CreateTrayMenu]: https://wiki.libsdl.org/SDL3/SDL_CreateTrayMenu
// func CreateTrayMenu(tray *Tray) *TrayMenu {
//	return sdlCreateTrayMenu(tray)
// }

// [CreateTraySubmenu] creates a submenu for a system tray entry.
//
// [CreateTraySubmenu]: https://wiki.libsdl.org/SDL3/SDL_CreateTraySubmenu
// func CreateTraySubmenu(entry *TrayEntry) *TrayMenu {
//	return sdlCreateTraySubmenu(entry)
// }

// [GetTrayMenu] gets a previously created tray menu.
//
// [GetTrayMenu]: https://wiki.libsdl.org/SDL3/SDL_GetTrayMenu
// func GetTrayMenu(tray *Tray) *TrayMenu {
//	return sdlGetTrayMenu(tray)
// }

// [GetTraySubmenu] gets a previously created tray entry submenu.
//
// [GetTraySubmenu]: https://wiki.libsdl.org/SDL3/SDL_GetTraySubmenu
// func GetTraySubmenu(entry *TrayEntry) *TrayMenu {
//	return sdlGetTraySubmenu(entry)
// }

// [GetTrayEntries] returns a list of entries in the menu, in order.
//
// [GetTrayEntries]: https://wiki.libsdl.org/SDL3/SDL_GetTrayEntries
// func GetTrayEntries(menu *TrayMenu, size *int32) **TrayEntry {
//	return sdlGetTrayEntries(menu, size)
// }

// [RemoveTrayEntry] removes a tray entry.
//
// [RemoveTrayEntry]: https://wiki.libsdl.org/SDL3/SDL_RemoveTrayEntry
// func RemoveTrayEntry(entry *TrayEntry)  {
//	sdlRemoveTrayEntry(entry)
// }

// [InsertTrayEntryAt] inserts a tray entry at a given position.
//
// [InsertTrayEntryAt]: https://wiki.libsdl.org/SDL3/SDL_InsertTrayEntryAt
// func InsertTrayEntryAt(menu *TrayMenu, pos int32, label string, flags TrayEntryFlags) *TrayEntry {
//	return sdlInsertTrayEntryAt(menu, pos, label, flags)
// }

// [SetTrayEntryLabel] sets the label of an entry.
//
// [SetTrayEntryLabel]: https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryLabel
// func SetTrayEntryLabel(entry *TrayEntry, label string) {
//	sdlSetTrayEntryLabel(entry, label)
// }

// [GetTrayEntryLabel] gets the label of an entry.
//
// [GetTrayEntryLabel]: https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryLabel
// func GetTrayEntryLabel(entry *TrayEntry) string {
//	return sdlGetTrayEntryLabel(entry)
// }

// [SetTrayEntryChecked] sets whether or not an entry is checked.
//
// [SetTrayEntryChecked]: https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryChecked
// func SetTrayEntryChecked(entry *TrayEntry, checked bool) {
//	sdlSetTrayEntryChecked(entry, checked)
// }

// [GetTrayEntryChecked] gets whether or not an entry is checked.
//
// [GetTrayEntryChecked]: https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryChecked
// func GetTrayEntryChecked(entry *TrayEntry) bool {
//	return sdlGetTrayEntryChecked(entry)
// }

// [SetTrayEntryEnabled] sets whether or not an entry is enabled.
//
// [SetTrayEntryEnabled]: https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryEnabled
// func SetTrayEntryEnabled(entry *TrayEntry, enabled bool) {
//	sdlSetTrayEntryEnabled(entry, enabled)
// }

// [GetTrayEntryEnabled] gets whether or not an entry is enabled.
//
// [GetTrayEntryEnabled]: https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryEnabled
// func GetTrayEntryEnabled(entry *TrayEntry) bool {
//	return sdlGetTrayEntryEnabled(entry)
// }

// [SetTrayEntryCallback] sets a callback to be invoked when the entry is selected.
//
// [SetTrayEntryCallback]: https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryCallback
// func SetTrayEntryCallback(entry *TrayEntry, callback TrayCallback, userdata unsafe.Pointer) {
//	sdlSetTrayEntryCallback(entry, callback, userdata)
// }

// [ClickTrayEntry] simulates a click on a tray entry.
//
// [ClickTrayEntry]: https://wiki.libsdl.org/SDL3/SDL_ClickTrayEntry
// func ClickTrayEntry(entry *TrayEntry) {
//	sdlClickTrayEntry(entry)
// }

// [DestroyTray] destroys a tray object.
//
// [DestroyTray]: https://wiki.libsdl.org/SDL3/SDL_DestroyTray
// func DestroyTray(tray *Tray) {
//	sdlDestroyTray(tray)
// }

// [GetTrayEntryParent] gets the menu containing a certain tray entry.
//
// [GetTrayEntryParent]: https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryParent
// func GetTrayEntryParent(entry *TrayEntry) *TrayMenu {
//	return sdlGetTrayEntryParent(entry)
// }

// [GetTrayMenuParentEntry] gets the entry for which the menu is a submenu, if the current menu is a.
//
// [GetTrayMenuParentEntry]: https://wiki.libsdl.org/SDL3/SDL_GetTrayMenuParentEntry
// func GetTrayMenuParentEntry(menu *TrayMenu) *TrayEntry {
//	return sdlGetTrayMenuParentEntry(menu)
// }

// [GetTrayMenuParentTray] gets the tray for which this menu is the first-level menu, if the current.
//
// [GetTrayMenuParentTray]: https://wiki.libsdl.org/SDL3/SDL_GetTrayMenuParentTray
// func GetTrayMenuParentTray(menu *TrayMenu) *Tray {
//	return sdlGetTrayMenuParentTray(menu)
// }

// [UpdateTrays] updates the trays.
//
// [UpdateTrays]: https://wiki.libsdl.org/SDL3/SDL_UpdateTrays
// func UpdateTrays() {
// 	sdlUpdateTrays()
// }
