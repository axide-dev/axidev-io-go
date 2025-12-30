// Example usage of the axidev-io-go keyboard package.
// This demonstrates keyboard input injection, event listening, and logging.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	axidevio "github.com/ziedyousfi/axidev-io-go"
	"github.com/ziedyousfi/axidev-io-go/keyboard"
)

func main() {
	fmt.Println("=== axidev-io-go Keyboard Example ===")
	fmt.Println()

	// --- Library Info ---
	fmt.Println("Library version:", axidevio.LibraryVersion())
	fmt.Println()

	// --- Logging Configuration ---
	fmt.Println("--- Logging ---")
	fmt.Printf("Current log level: %d\n", axidevio.GetLogLevel())
	fmt.Printf("Debug logging enabled: %v\n", axidevio.IsLogEnabled(axidevio.LogLevelDebug))
	fmt.Printf("Info logging enabled: %v\n", axidevio.IsLogEnabled(axidevio.LogLevelInfo))

	// Enable debug logging to see internal library messages
	axidevio.SetLogLevel(axidevio.LogLevelDebug)
	fmt.Println("Enabled debug logging")
	fmt.Println()

	// --- Keyboard Sender ---
	fmt.Println("--- Keyboard Sender ---")
	sender, err := keyboard.NewSender()
	if err != nil {
		fmt.Println("Failed to create keyboard sender:", err)
		os.Exit(1)
	}
	defer sender.Close()

	fmt.Println("Sender ready:", sender.IsReady())
	fmt.Println("Backend type:", sender.BackendType())

	// Display capabilities
	caps := sender.Capabilities()
	fmt.Println("Capabilities:")
	fmt.Printf("  Can inject keys:       %v\n", caps.CanInjectKeys)
	fmt.Printf("  Can inject text:       %v\n", caps.CanInjectText)
	fmt.Printf("  Can simulate HID:      %v\n", caps.CanSimulateHID)
	fmt.Printf("  Supports key repeat:   %v\n", caps.SupportsKeyRepeat)
	fmt.Printf("  Needs accessibility:   %v\n", caps.NeedsAccessibilityPerm)
	fmt.Printf("  Needs input monitor:   %v\n", caps.NeedsInputMonitoringPerm)
	fmt.Printf("  Needs uinput access:   %v\n", caps.NeedsUinputAccess)
	fmt.Println()

	// Request permissions if needed
	if caps.NeedsAccessibilityPerm || caps.NeedsInputMonitoringPerm {
		fmt.Println("Requesting permissions...")
		if !sender.RequestPermissions() {
			fmt.Println("Warning: Permissions not granted, some features may not work")
		} else {
			fmt.Println("Permissions granted!")
		}
		fmt.Println()
	}

	// --- Key Conversion ---
	fmt.Println("--- Key Conversion ---")
	testKeys := []string{"A", "Enter", "Return", "Space", "Escape", "F1", "Ctrl", "Shift"}
	for _, name := range testKeys {
		key := keyboard.StringToKey(name)
		backName := keyboard.KeyToString(key)
		fmt.Printf("  %q -> Key(%d) -> %q\n", name, key, backName)
	}
	fmt.Println()

	// --- Modifier Examples ---
	fmt.Println("--- Modifiers ---")
	mods := keyboard.ModCtrl | keyboard.ModShift
	fmt.Printf("Combined modifiers (Ctrl+Shift): 0x%02X\n", mods)
	fmt.Printf("  HasCtrl: %v, HasShift: %v, HasAlt: %v\n",
		mods.HasCtrl(), mods.HasShift(), mods.HasAlt())
	fmt.Println()

	// --- Typing Demo ---
	if caps.CanInjectText {
		fmt.Println("--- Text Injection Demo ---")
		fmt.Println("Will type text in 3 seconds... (switch to a text editor!)")
		time.Sleep(3 * time.Second)

		if err := sender.TypeText("Hello from axidev-io-go! 🎉"); err != nil {
			fmt.Println("TypeText error:", err)
		} else {
			fmt.Println("Text typed successfully!")
		}

		// Type a newline
		if err := sender.Tap(keyboard.StringToKey("Return")); err != nil {
			fmt.Println("Tap error:", err)
		}
		fmt.Println()
	}

	// --- Keyboard Listener ---
	fmt.Println("--- Keyboard Listener ---")
	listener, err := keyboard.NewListener()
	if err != nil {
		fmt.Println("Failed to create keyboard listener:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Starting keyboard listener (press Ctrl+C to exit)...")
	fmt.Println("Press keys to see events:")
	fmt.Println()

	err = listener.Start(func(event keyboard.KeyEvent) {
		action := "↑ released"
		if event.Pressed {
			action = "↓ pressed "
		}

		keyName := event.KeyName()
		if keyName == "" {
			keyName = fmt.Sprintf("Unknown(0x%04X)", event.Key)
		}

		// Show modifier state
		modStr := ""
		if event.Modifiers.HasCtrl() {
			modStr += "Ctrl+"
		}
		if event.Modifiers.HasShift() {
			modStr += "Shift+"
		}
		if event.Modifiers.HasAlt() {
			modStr += "Alt+"
		}
		if event.Modifiers.HasSuper() {
			modStr += "Super+"
		}

		charInfo := ""
		if event.Codepoint > 0 && event.Codepoint < 0x110000 {
			charInfo = fmt.Sprintf(" char='%c'", event.Rune())
		}

		fmt.Printf("  %s %s%s%s\n", action, modStr, keyName, charInfo)
	})
	if err != nil {
		fmt.Println("Failed to start listener:", err)
		os.Exit(1)
	}

	fmt.Printf("Listening: %v\n", listener.IsListening())

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\n--- Cleanup ---")
	fmt.Println("Stopping listener...")
	listener.Stop()
	fmt.Printf("Listening after stop: %v\n", listener.IsListening())

	// Check for any errors
	if lastErr := axidevio.GetLastError(); lastErr != "" {
		fmt.Println("Last error:", lastErr)
		axidevio.ClearLastError()
	}

	fmt.Println("Done!")
}
