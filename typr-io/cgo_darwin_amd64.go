//go:build darwin && amd64

package typrio

/*
#cgo LDFLAGS: -L${SRCDIR}/../lib/macos-x86_64 -ltypr_io -Wl,-rpath,${SRCDIR}/../lib/macos-x86_64
*/
import "C"
