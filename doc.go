// Package axidevio provides Go bindings for the axidev-io library.
//
// This package provides common utilities (logging, error handling, version info)
// for the axidev-io library. Keyboard-specific functionality is in the keyboard subpackage.
//
// # Package Structure
//
//   - axidevio: Common utilities (logging, errors, version)
//   - axidevio/keyboard: Keyboard input injection and event monitoring
//
// # Logging
//
// Control library logging verbosity:
//
//	axidevio.SetLogLevel(axidevio.LogLevelDebug) // Enable debug logs
//	axidevio.SetLogLevel(axidevio.LogLevelError) // Only show errors
//
// # Keyboard Usage
//
// See the keyboard subpackage for keyboard input injection and monitoring:
//
//	import "github.com/axide-dev/axidev-io-go/keyboard"
//
//	sender, _ := keyboard.NewSender()
//	sender.TypeText("Hello!")
//	sender.Combo(keyboard.ModCtrl, keyboard.StringToKey("S"))
//
// # Runtime Path Configuration
//
// You may need to configure the library path at runtime:
//
// On macOS:
//
//	export DYLD_LIBRARY_PATH=/path/to/lib:$DYLD_LIBRARY_PATH
//
// On Linux:
//
//	export LD_LIBRARY_PATH=/path/to/lib:$LD_LIBRARY_PATH
package axidevio
