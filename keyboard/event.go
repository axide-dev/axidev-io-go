package keyboard

// KeyEvent represents a keyboard event received by the listener.
type KeyEvent struct {
	// Codepoint is the Unicode codepoint produced by the event (0 if none).
	Codepoint uint32

	// Key is the logical key ID (0 if unknown).
	Key Key

	// Modifiers is the current modifier bitmask.
	Modifiers Modifier

	// Pressed is true for key press, false for key release.
	Pressed bool
}

// IsPress returns true if this is a key press event.
func (e KeyEvent) IsPress() bool { return e.Pressed }

// IsRelease returns true if this is a key release event.
func (e KeyEvent) IsRelease() bool { return !e.Pressed }

// Rune returns the Unicode rune for this event, or 0 if none.
func (e KeyEvent) Rune() rune { return rune(e.Codepoint) }

// KeyName returns the canonical name of the key.
func (e KeyEvent) KeyName() string { return KeyToString(e.Key) }
