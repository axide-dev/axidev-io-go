//go:build linux && arm64

package axidevio

/*
#cgo LDFLAGS: ${SRCDIR}/lib/linux-arm64/libaxidev_io.a -lstdc++ -linput -ludev -lxkbcommon -lpthread
*/
import "C"
