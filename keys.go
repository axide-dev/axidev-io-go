package axidevio

/*
#include <axidev-io/c_api.h>
*/
import "C"

// Key represents a logical key identifier.
type Key uint16

// KeyToString converts a Key to its canonical string name.
func KeyToString(key Key) string {
	cStr := C.typr_io_key_to_string(C.typr_io_key_t(key))
	if cStr == nil {
		return ""
	}
	defer freeString(cStr)
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
	cName := cString(name)
	defer freeCString(cName)
	return Key(C.typr_io_string_to_key(cName))
}
