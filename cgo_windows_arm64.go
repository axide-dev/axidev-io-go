//go:build windows && arm64

package axidevio

/*
#cgo windows,arm64 LDFLAGS: -L${SRCDIR}/lib/windows-arm64 -laxidev_io -luser32 -lstdc++ -lm -lkernel32
*/
import "C"
