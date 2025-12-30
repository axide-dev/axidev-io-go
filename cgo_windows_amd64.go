//go:build windows && amd64

package axidevio

/*
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/lib/windows-x64 -laxidev_io -luser32 -lc++ -lm
*/
import "C"
