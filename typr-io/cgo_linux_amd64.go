//go:build linux && amd64

package typrio

/*
#cgo LDFLAGS: -L${SRCDIR}/../lib/linux-x86_64 -ltypr_io -Wl,-rpath,${SRCDIR}/../lib/linux-x86_64
*/
import "C"
