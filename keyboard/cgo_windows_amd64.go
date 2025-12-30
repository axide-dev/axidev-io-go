//go:build windows && amd64

package keyboard

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/../lib/windows-x64 -laxidev_io -luser32 -lstdc++ -lm
*/
import "C"
