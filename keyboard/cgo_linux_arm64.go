//go:build linux && arm64

package keyboard

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#cgo LDFLAGS: ${SRCDIR}/../lib/linux-arm64/libaxidev_io.a -lstdc++ -linput -ludev -lxkbcommon -lpthread
*/
import "C"
