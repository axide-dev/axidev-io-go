//go:build windows && amd64

package typrio

/*
#cgo LDFLAGS: -L${SRCDIR}/../lib/windows-x64 -ltypr_io
*/
import "C"
