package typrio

// Capabilities describes the features supported by the active backend.
type Capabilities struct {
	// CanInjectKeys is true if the backend can send physical key events.
	CanInjectKeys bool

	// CanInjectText is true if the backend can inject arbitrary Unicode text.
	CanInjectText bool

	// CanSimulateHID is true if the backend simulates low-level HID events (e.g., uinput).
	CanSimulateHID bool

	// SupportsKeyRepeat is true if key repeat is supported by the backend.
	SupportsKeyRepeat bool

	// NeedsAccessibilityPerm is true if accessibility permission is required (platform-dependent).
	NeedsAccessibilityPerm bool

	// NeedsInputMonitoringPerm is true if input monitoring permission is required (platform-dependent).
	NeedsInputMonitoringPerm bool

	// NeedsUinputAccess is true if uinput or similar device access is required.
	NeedsUinputAccess bool
}
