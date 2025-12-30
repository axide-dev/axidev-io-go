package keyboard

/*
#include <axidev-io/c_api.h>
#include <stdlib.h>
*/
import "C"

import "unsafe"

// Key represents a logical keyboard key identifier.
type Key uint16

// KeyToString converts a Key to its canonical string name.
func KeyToString(key Key) string {
	cStr := C.axidev_io_keyboard_key_to_string(C.axidev_io_keyboard_key_t(key))
	if cStr == nil {
		return ""
	}
	defer C.axidev_io_free_string(cStr)
	return C.GoString(cStr)
}

// StringToKey parses a key name string to a Key value.
// Returns 0 (Key unknown) for unrecognized inputs.
// The parsing is case-insensitive and accepts common aliases like "esc", "space".
//
// Common key names:
//   - Letters: "A"-"Z"
//   - Digits: "0"-"9"
//   - Function keys: "F1"-"F24"
//   - Navigation: "Up", "Down", "Left", "Right", "Home", "End", "PageUp", "PageDown"
//   - Editing: "Backspace", "Delete", "Insert", "Tab", "Return", "Enter", "Space"
//   - Modifiers: "Shift", "Control", "Alt", "Super", "Meta", "CapsLock", "NumLock"
//   - Special: "Escape", "Esc", "PrintScreen", "ScrollLock", "Pause"
func StringToKey(name string) Key {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return Key(C.axidev_io_keyboard_string_to_key(cName))
}
