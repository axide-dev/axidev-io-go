// Package keyboard provides keyboard input injection and global event monitoring.
//
// This package wraps the axidev-io keyboard C API, providing a safe and idiomatic
// Go interface for simulating keyboard input and monitoring global keyboard events.
//
// # Sender - Keyboard Input Injection
//
// Create a Sender to inject keyboard input:
//
//	sender, err := keyboard.NewSender()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer sender.Close()
//
//	// Type some text
//	sender.TypeText("Hello, World!")
//
//	// Press a key combination (Ctrl+S)
//	sender.Combo(keyboard.ModCtrl, keyboard.StringToKey("S"))
//
//	// Tap a single key
//	sender.Tap(keyboard.StringToKey("Enter"))
//
// # Listener - Global Keyboard Event Monitoring
//
// Create a Listener to monitor keyboard events:
//
//	listener, err := keyboard.NewListener()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer listener.Close()
//
//	err = listener.Start(func(event keyboard.KeyEvent) {
//	    if event.Pressed {
//	        fmt.Printf("Key pressed: %s\n", keyboard.KeyToString(event.Key))
//	    }
//	})
//
// # Thread Safety
//
// All Sender and Listener methods are thread-safe. However, the listener callback
// is invoked from a background thread, so callbacks must be thread-safe and avoid
// blocking operations.
//
// # Permissions
//
// On some platforms, accessibility or input monitoring permissions may be required.
// Check the Capabilities struct returned by sender.Capabilities() to determine
// what permissions are needed.
package keyboard
