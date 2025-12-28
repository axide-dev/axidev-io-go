//go:build linux && arm64

package typrio

/*
#cgo LDFLAGS: -L${SRCDIR}/../lib/linux-arm64 -ltypr_io -Wl,-rpath,${SRCDIR}/../lib/linux-arm64
*/
import "C"
