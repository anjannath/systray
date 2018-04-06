// +build darwin

package systray

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c -fobjc-arc
#cgo darwin LDFLAGS: -framework Cocoa

#include "systray_darwin.h"
*/
import "C"

import (
	"unsafe"
	"io/ioutil"
)

func nativeLoop() (err error) {
	_, err = C.nativeLoop()
	return
}

func quit() {
	C.quit()
}

// Sets the systray icon.
// iconBytes should be the content of .ico/.jpg/.png
func setIcon(iconBytes []byte) error {
	cstr := (*C.char)(unsafe.Pointer(&iconBytes[0]))
	return C.setIcon(cstr, (C.int)(len(iconBytes)))
}

// Sets the systray icon by path to file.
// File should be one of .ico/.jpg/.png
func setIconPath(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return setIcon(b)
}

// SetTitle sets the systray title, only available on Mac.
func setTitle(title string) error {
	return C.setTitle(C.CString(title))
}

// SetTooltip sets the systray tooltip to display on mouse hover of the tray icon,
// only available on Mac and Windows.
func setTooltip(tooltip string) error {
	return C.setTooltip(C.CString(tooltip))
}

func addOrUpdateMenuItem(item *MenuItem) error {
	var disabled C.short
	if item.disabled {
		disabled = 1
	}
	var checked C.short
	if item.checked {
		checked = 1
	}
	return C.add_or_update_menu_item(
		C.int(item.id),
		C.CString(item.title),
		C.CString(item.tooltip),
		disabled,
		checked,
	)
}

func addSeparator(id int32) error {
	return C.add_separator(C.int(id))
}

func hideMenuItem(item *MenuItem) error {
	return C.hide_menu_item(
		C.int(item.id),
	)
}

func showMenuItem(item *MenuItem) error {
	return C.show_menu_item(
		C.int(item.id),
	)
}

//export systray_ready
func systray_ready() {
	systrayReady()
}

//export systray_on_exit
func systray_on_exit() {
	systrayExit()
}

//export systray_menu_item_selected
func systray_menu_item_selected(cID C.int) {
	systrayMenuItemSelected(int32(cID))
}
