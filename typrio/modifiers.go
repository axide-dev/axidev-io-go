package typrio

/*
#include <typr-io/c_api.h>
*/
import "C"

// Modifier represents a bitmask of modifier keys.
type Modifier uint8

// Modifier constants matching the C API.
// These can be combined using bitwise OR:
//
//	mods := typrio.ModCtrl | typrio.ModShift
//	sender.Combo(mods, typrio.StringToKey("S"))  // Ctrl+Shift+S
const (
	ModShift    Modifier = C.TYPR_IO_MOD_SHIFT
	ModCtrl     Modifier = C.TYPR_IO_MOD_CTRL
	ModAlt      Modifier = C.TYPR_IO_MOD_ALT
	ModSuper    Modifier = C.TYPR_IO_MOD_SUPER
	ModCapsLock Modifier = C.TYPR_IO_MOD_CAPSLOCK
	ModNumLock  Modifier = C.TYPR_IO_MOD_NUMLOCK
)

// HasShift returns true if the Shift modifier is set.
func (m Modifier) HasShift() bool { return m&ModShift != 0 }

// HasCtrl returns true if the Ctrl modifier is set.
func (m Modifier) HasCtrl() bool { return m&ModCtrl != 0 }

// HasAlt returns true if the Alt modifier is set.
func (m Modifier) HasAlt() bool { return m&ModAlt != 0 }

// HasSuper returns true if the Super (Windows/Command) modifier is set.
func (m Modifier) HasSuper() bool { return m&ModSuper != 0 }

// HasCapsLock returns true if CapsLock is active.
func (m Modifier) HasCapsLock() bool { return m&ModCapsLock != 0 }

// HasNumLock returns true if NumLock is active.
func (m Modifier) HasNumLock() bool { return m&ModNumLock != 0 }
