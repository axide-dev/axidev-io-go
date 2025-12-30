//go:build windows && amd64

package axidevio

/*
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/lib/windows-x64 -ltypr_io -luser32 -lstdc++ -lm
*/
import "C"
