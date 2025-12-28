//go:build windows && arm64

package typrio

/*
#cgo LDFLAGS: -L${SRCDIR}/../lib/windows-arm64 -ltypr_io
*/
import "C"
