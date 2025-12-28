// Example usage of the typrio package.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ziedyousfi/typr-io-go/typr-io"
)

func main() {
	fmt.Println("typr-io-go example")
	fmt.Println("Library version:", typrio.LibraryVersion())

	// --- Sender Example ---
	sender, err := typrio.NewSender()
	if err != nil {
		fmt.Println("Failed to create sender:", err)
		os.Exit(1)
	}
	defer sender.Close()

	fmt.Println("Sender ready:", sender.IsReady())

	caps := sender.Capabilities()
	fmt.Printf("Capabilities: %+v\n", caps)

	if caps.NeedsAccessibilityPerm {
		fmt.Println("Requesting accessibility permissions...")
		if !sender.RequestPermissions() {
			fmt.Println("Permissions not granted, some features may not work")
		}
	}

	// Convert key names to Key values
	keyA := typrio.StringToKey("A")
	keyEnter := typrio.StringToKey("Return")
	fmt.Printf("Key A = %d, Key Enter = %d\n", keyA, keyEnter)
	fmt.Printf("Key A name = %q\n", typrio.KeyToString(keyA))

	// Type some text (if supported)
	if caps.CanInjectText {
		fmt.Println("Typing text in 5 seconds...")
		time.Sleep(5	 * time.Second)
		if err := sender.TypeText("Hello from typr-io-go!"); err != nil {
			fmt.Println("TypeText error:", err)
		}
	}

	// --- Listener Example ---
	listener, err := typrio.NewListener()
	if err != nil {
		fmt.Println("Failed to create listener:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("\nStarting key listener (press Ctrl+C to exit)...")
	err = listener.Start(func(event typrio.KeyEvent) {
		action := "released"
		if event.Pressed {
			action = "pressed"
		}
		keyName := typrio.KeyToString(event.Key)
		if keyName == "" {
			keyName = fmt.Sprintf("0x%04X", event.Key)
		}
		fmt.Printf("Key %s: %s (codepoint=%d, mods=0x%02X)\n",
			action, keyName, event.Codepoint, event.Modifiers)
	})
	if err != nil {
		fmt.Println("Failed to start listener:", err)
		os.Exit(1)
	}

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nStopping listener...")
	listener.Stop()
	fmt.Println("Done.")
}
