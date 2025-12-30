package keyboard

/*
#include <axidev-io/c_api.h>

// Callback bridge for Go
extern void goKeyboardListenerCallback(uint32_t codepoint, uint16_t key, uint8_t mods, _Bool pressed, void* user_data);

static void keyboard_listener_callback_bridge(uint32_t codepoint, axidev_io_keyboard_key_t key,
                                              axidev_io_keyboard_modifier_t mods, _Bool pressed,
                                              void* user_data) {
    goKeyboardListenerCallback(codepoint, key, mods, pressed, user_data);
}

static _Bool start_keyboard_listener_with_bridge(axidev_io_keyboard_listener_t listener, void* user_data) {
    return axidev_io_keyboard_listener_start(listener, keyboard_listener_callback_bridge, user_data);
}
*/
import "C"

import (
	"errors"
	"runtime/cgo"
	"sync"
	"unsafe"

	axidevio "github.com/ziedyousfi/axidev-io-go"
)

// ListenerCallback is the function signature for keyboard event callbacks.
// The callback may be invoked from an internal thread and must be thread-safe.
type ListenerCallback func(event KeyEvent)

// Listener provides global keyboard event monitoring.
type Listener struct {
	handle    C.axidev_io_keyboard_listener_t
	mu        sync.Mutex
	callback  ListenerCallback
	cgoHandle cgo.Handle
}

// NewListener creates a new keyboard Listener instance.
// Returns an error if allocation fails.
func NewListener() (*Listener, error) {
	handle := C.axidev_io_keyboard_listener_create()
	if handle == nil {
		return nil, errors.New("failed to create keyboard listener")
	}
	return &Listener{handle: handle}, nil
}

// Close destroys the listener and releases resources.
// Safe to call multiple times.
func (l *Listener) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.handle != nil {
		C.axidev_io_keyboard_listener_stop(l.handle)
		C.axidev_io_keyboard_listener_destroy(l.handle)
		l.handle = nil
	}
	if l.cgoHandle != 0 {
		l.cgoHandle.Delete()
		l.cgoHandle = 0
	}
}

// Start begins listening for keyboard events.
// The callback may be invoked from an internal thread.
// The callback must be thread-safe and avoid long-blocking work.
func (l *Listener) Start(callback ListenerCallback) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.handle == nil {
		return errors.New("listener is closed")
	}
	if callback == nil {
		return errors.New("callback cannot be nil")
	}

	l.callback = callback
	l.cgoHandle = cgo.NewHandle(l)

	if !C.start_keyboard_listener_with_bridge(l.handle, unsafe.Pointer(&l.cgoHandle)) {
		l.cgoHandle.Delete()
		l.cgoHandle = 0
		return axidevio.GetLastErrorOrDefault("failed to start keyboard listener")
	}
	return nil
}

// Stop stops listening for keyboard events.
// Safe to call from any thread; no-op if not running.
func (l *Listener) Stop() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.handle != nil {
		C.axidev_io_keyboard_listener_stop(l.handle)
	}
	if l.cgoHandle != 0 {
		l.cgoHandle.Delete()
		l.cgoHandle = 0
	}
}

// IsListening returns true if the listener is currently active.
func (l *Listener) IsListening() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.handle == nil {
		return false
	}
	return bool(C.axidev_io_keyboard_listener_is_listening(l.handle))
}
