package typrio

/*
#include <typr-io/c_api.h>

// Callback bridge for Go
extern void goListenerCallback(uint32_t codepoint, uint16_t key, uint8_t mods, _Bool pressed, void* user_data);

static void listener_callback_bridge(uint32_t codepoint, typr_io_key_t key,
                                     typr_io_modifier_t mods, _Bool pressed,
                                     void* user_data) {
    goListenerCallback(codepoint, key, mods, pressed, user_data);
}

static _Bool start_listener_with_bridge(typr_io_listener_t listener, void* user_data) {
    return typr_io_listener_start(listener, listener_callback_bridge, user_data);
}
*/
import "C"

import (
	"errors"
	"runtime/cgo"
	"sync"
	"unsafe"
)

// ListenerCallback is the function signature for key event callbacks.
// The callback may be invoked from an internal thread and must be thread-safe.
type ListenerCallback func(event KeyEvent)

// Listener provides global key event monitoring.
type Listener struct {
	handle    C.typr_io_listener_t
	mu        sync.Mutex
	callback  ListenerCallback
	cgoHandle cgo.Handle
}

// NewListener creates a new Listener instance.
// Returns an error if allocation fails.
func NewListener() (*Listener, error) {
	handle := C.typr_io_listener_create()
	if handle == nil {
		return nil, errors.New("failed to create listener")
	}
	return &Listener{handle: handle}, nil
}

// Close destroys the listener and releases resources.
// Safe to call multiple times.
func (l *Listener) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.handle != nil {
		C.typr_io_listener_stop(l.handle)
		C.typr_io_listener_destroy(l.handle)
		l.handle = nil
	}
	if l.cgoHandle != 0 {
		l.cgoHandle.Delete()
		l.cgoHandle = 0
	}
}

// Start begins listening for key events.
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

	if !C.start_listener_with_bridge(l.handle, unsafe.Pointer(&l.cgoHandle)) {
		l.cgoHandle.Delete()
		l.cgoHandle = 0
		return getLastError("failed to start listener")
	}
	return nil
}

// Stop stops listening for key events.
// Safe to call from any thread; no-op if not running.
func (l *Listener) Stop() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.handle != nil {
		C.typr_io_listener_stop(l.handle)
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
	return bool(C.typr_io_listener_is_listening(l.handle))
}
