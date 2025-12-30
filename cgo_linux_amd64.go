//go:build linux && amd64

package axidevio

/*
#cgo LDFLAGS: ${SRCDIR}/lib/linux-x64/libaxidev_io.a -lstdc++ -linput -ludev -lxkbcommon -lpthread
*/
import "C"
