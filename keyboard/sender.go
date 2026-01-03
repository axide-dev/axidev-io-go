package keyboard

/*
#include <axidev-io/c_api.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"sync"
	"unsafe"

	axidevio "github.com/axide-dev/axidev-io-go"
)

// Sender provides keyboard input injection capabilities.
type Sender struct {
	handle C.axidev_io_keyboard_sender_t
	mu     sync.Mutex
}

// NewSender creates a new keyboard Sender instance.
// Returns an error if allocation fails.
func NewSender() (*Sender, error) {
	handle := C.axidev_io_keyboard_sender_create()
	if handle == nil {
		return nil, errors.New("failed to create keyboard sender")
	}
	return &Sender{handle: handle}, nil
}

// Close destroys the sender and releases resources.
// Safe to call multiple times.
func (s *Sender) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle != nil {
		C.axidev_io_keyboard_sender_destroy(s.handle)
		s.handle = nil
	}
}

// IsReady returns true if the backend is ready to inject keyboard events.
func (s *Sender) IsReady() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return false
	}
	return bool(C.axidev_io_keyboard_sender_is_ready(s.handle))
}

// BackendType returns the active backend type as an integer.
func (s *Sender) BackendType() uint8 {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return 0
	}
	return uint8(C.axidev_io_keyboard_sender_type(s.handle))
}

// Capabilities returns the backend capabilities.
func (s *Sender) Capabilities() Capabilities {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return Capabilities{}
	}
	var caps C.axidev_io_keyboard_capabilities_t
	C.axidev_io_keyboard_sender_get_capabilities(s.handle, &caps)
	return Capabilities{
		CanInjectKeys:            bool(caps.can_inject_keys),
		CanInjectText:            bool(caps.can_inject_text),
		CanSimulateHID:           bool(caps.can_simulate_hid),
		SupportsKeyRepeat:        bool(caps.supports_key_repeat),
		NeedsAccessibilityPerm:   bool(caps.needs_accessibility_perm),
		NeedsInputMonitoringPerm: bool(caps.needs_input_monitoring_perm),
		NeedsUinputAccess:        bool(caps.needs_uinput_access),
	}
}

// RequestPermissions requests runtime permissions required by the backend.
// Returns true if the backend is ready after requesting permissions.
func (s *Sender) RequestPermissions() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return false
	}
	return bool(C.axidev_io_keyboard_sender_request_permissions(s.handle))
}

// KeyDown simulates a physical key press.
func (s *Sender) KeyDown(key Key) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return errors.New("sender is closed")
	}
	if !C.axidev_io_keyboard_sender_key_down(s.handle, C.axidev_io_keyboard_key_t(key)) {
		return axidevio.GetLastErrorOrDefault("key down failed")
	}
	return nil
}

// KeyUp simulates a physical key release.
func (s *Sender) KeyUp(key Key) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return errors.New("sender is closed")
	}
	if !C.axidev_io_keyboard_sender_key_up(s.handle, C.axidev_io_keyboard_key_t(key)) {
		return axidevio.GetLastErrorOrDefault("key up failed")
	}
	return nil
}

// Tap simulates a key tap (press then release).
func (s *Sender) Tap(key Key) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return errors.New("sender is closed")
	}
	if !C.axidev_io_keyboard_sender_tap(s.handle, C.axidev_io_keyboard_key_t(key)) {
		return axidevio.GetLastErrorOrDefault("tap failed")
	}
	return nil
}

// ActiveModifiers returns the currently active modifiers.
func (s *Sender) ActiveModifiers() Modifier {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return 0
	}
	return Modifier(C.axidev_io_keyboard_sender_active_modifiers(s.handle))
}

// HoldModifier presses the specified modifier keys.
func (s *Sender) HoldModifier(mods Modifier) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return errors.New("sender is closed")
	}
	if !C.axidev_io_keyboard_sender_hold_modifier(s.handle, C.axidev_io_keyboard_modifier_t(mods)) {
		return axidevio.GetLastErrorOrDefault("hold modifier failed")
	}
	return nil
}

// ReleaseModifier releases the specified modifier keys.
func (s *Sender) ReleaseModifier(mods Modifier) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return errors.New("sender is closed")
	}
	if !C.axidev_io_keyboard_sender_release_modifier(s.handle, C.axidev_io_keyboard_modifier_t(mods)) {
		return axidevio.GetLastErrorOrDefault("release modifier failed")
	}
	return nil
}

// ReleaseAllModifiers releases all currently held modifiers.
func (s *Sender) ReleaseAllModifiers() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return errors.New("sender is closed")
	}
	if !C.axidev_io_keyboard_sender_release_all_modifiers(s.handle) {
		return axidevio.GetLastErrorOrDefault("release all modifiers failed")
	}
	return nil
}

// Combo executes a key combo: press modifiers, tap key, release modifiers.
func (s *Sender) Combo(mods Modifier, key Key) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return errors.New("sender is closed")
	}
	if !C.axidev_io_keyboard_sender_combo(s.handle, C.axidev_io_keyboard_modifier_t(mods), C.axidev_io_keyboard_key_t(key)) {
		return axidevio.GetLastErrorOrDefault("combo failed")
	}
	return nil
}

// TypeText injects UTF-8 text directly (layout-independent on supporting backends).
func (s *Sender) TypeText(text string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return errors.New("sender is closed")
	}
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	if !C.axidev_io_keyboard_sender_type_text_utf8(s.handle, cText) {
		return axidevio.GetLastErrorOrDefault("type text failed")
	}
	return nil
}

// TypeCharacter injects a single Unicode codepoint.
func (s *Sender) TypeCharacter(codepoint rune) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle == nil {
		return errors.New("sender is closed")
	}
	if !C.axidev_io_keyboard_sender_type_character(s.handle, C.uint32_t(codepoint)) {
		return axidevio.GetLastErrorOrDefault("type character failed")
	}
	return nil
}

// Flush forces delivery of pending keyboard events.
func (s *Sender) Flush() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle != nil {
		C.axidev_io_keyboard_sender_flush(s.handle)
	}
}

// SetKeyDelay sets the delay (in microseconds) used by tap/combo operations.
func (s *Sender) SetKeyDelay(delayMicroseconds uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.handle != nil {
		C.axidev_io_keyboard_sender_set_key_delay(s.handle, C.uint32_t(delayMicroseconds))
	}
}
