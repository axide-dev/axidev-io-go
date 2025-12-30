package keyboard

/*
#include <stdint.h>
#include <stdbool.h>
*/
import "C"

import (
	"runtime/cgo"
	"unsafe"
)

//export goKeyboardListenerCallback
func goKeyboardListenerCallback(codepoint C.uint32_t, key C.uint16_t, mods C.uint8_t, pressed C._Bool, userData unsafe.Pointer) {
	handle := *(*cgo.Handle)(userData)
	listener := handle.Value().(*Listener)

	if listener.callback != nil {
		event := KeyEvent{
			Codepoint: uint32(codepoint),
			Key:       Key(key),
			Modifiers: Modifier(mods),
			Pressed:   bool(pressed),
		}
		listener.callback(event)
	}
}
