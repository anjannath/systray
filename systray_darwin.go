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
	_, err := C.setIcon(cstr, (C.int)(len(iconBytes)))
	return err
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
	_, err := C.setTitle(C.CString(title))
	return err
}

// SetTooltip sets the systray tooltip to display on mouse hover of the tray icon,
// only available on Mac and Windows.
func setTooltip(tooltip string) error {
	_, err := C.setTooltip(C.CString(tooltip))
	return err
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
	_, err := C.add_or_update_menu_item(
		C.int(item.id),
		C.CString(item.title),
		C.CString(item.tooltip),
		disabled,
		checked,
	)
	return err
}

func addSeparator(id int32) error {
	_, err := C.add_separator(C.int(id))
	return err
}

func hideMenuItem(item *MenuItem) error {
	_, err := C.hide_menu_item(
		C.int(item.id),
	)
	return err
}

func showMenuItem(item *MenuItem) error {
	_, err := C.show_menu_item(
		C.int(item.id),
	)
	return err
}

//export systray_ready
func systray_ready() {
	go systrayReady()
}

//export systray_on_exit
func systray_on_exit() {
	systrayExit()
}

//export systray_menu_item_selected
func systray_menu_item_selected(cID C.int) {
	systrayMenuItemSelected(int32(cID))
}
