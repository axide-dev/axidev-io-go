//go:build darwin && arm64

package typrio

/*
#cgo LDFLAGS: -L${SRCDIR}/../lib/macos-arm64 -ltypr_io -Wl,-rpath,${SRCDIR}/../lib/macos-arm64
*/
import "C"
