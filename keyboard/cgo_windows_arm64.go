//go:build windows && arm64

package keyboard

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#cgo windows,arm64 LDFLAGS: -L${SRCDIR}/../lib/windows-arm64 -laxidev_io -luser32 -lstdc++ -lm
*/
import "C"
