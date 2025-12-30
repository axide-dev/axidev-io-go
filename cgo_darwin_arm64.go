//go:build darwin && arm64

package axidevio

/*
#cgo LDFLAGS: ${SRCDIR}/lib/macos-arm64/libaxidev_io.a -lstdc++ -framework ApplicationServices -framework Carbon -framework Foundation -framework CoreGraphics
*/
import "C"
