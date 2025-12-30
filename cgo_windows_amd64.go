//go:build windows && amd64

package axidevio

/*
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/lib/windows-amd64 -laxidev_io -luser32 -lstdc++ -lm -lkernel32
*/
import "C"