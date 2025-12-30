//go:build darwin && amd64

package axidevio

/*
#cgo LDFLAGS: ${SRCDIR}/lib/macos-x86_64/libaxidev_io.a -lstdc++ -framework ApplicationServices -framework Carbon -framework Foundation -framework CoreGraphics
*/
import "C"
