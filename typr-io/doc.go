// Package typrio provides Go bindings for the typr-io input injection and listening library.
//
// This package wraps the typr-io C API, providing a safe and idiomatic Go interface
// for simulating keyboard input and monitoring global key events.
//
// # Quick Start
//
// Create a Sender to inject keyboard input:
//
//	sender, err := typrio.NewSender()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer sender.Close()
//
//	// Type some text
//	sender.TypeText("Hello, World!")
//
//	// Press a key combination
//	sender.Combo(typrio.ModCtrl, typrio.StringToKey("S"))
//
// Create a Listener to monitor keyboard events:
//
//	listener, err := typrio.NewListener()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer listener.Close()
//
//	listener.Start(func(event typrio.KeyEvent) {
//	    fmt.Printf("Key: %s, Pressed: %v\n", typrio.KeyToString(event.Key), event.Pressed)
//	})
//
// # Thread Safety
//
// All Sender and Listener methods are thread-safe. However, the listener callback
// is invoked from a background thread, so callbacks must be thread-safe and avoid
// blocking operations.
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
package typrio
