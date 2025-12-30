//go:build windows && arm64

package axidevio

/*
#cgo windows,arm64 LDFLAGS: -L${SRCDIR}/lib/windows-arm64 -ltypr_io -luser32 -lstdc++ -lm
*/
import "C"
