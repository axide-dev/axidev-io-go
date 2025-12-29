//go:build windows && amd64

package typrio

/*
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/lib/windows-x64 -ltypr_io -luser32
*/
import "C"
