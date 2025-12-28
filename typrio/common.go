package typrio

/*
#cgo CFLAGS: -I${SRCDIR}/../include
#cgo LDFLAGS: -L${SRCDIR}/../lib -ltypr_io

#include <typr-io/c_api.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

// LibraryVersion returns the typr-io library version string.
func LibraryVersion() string {
	return C.GoString(C.typr_io_library_version())
}

// GetLastError returns the last error string from the library, if any.
func GetLastError() string {
	cStr := C.typr_io_get_last_error()
	if cStr == nil {
		return ""
	}
	defer C.typr_io_free_string(cStr)
	return C.GoString(cStr)
}

// ClearLastError clears the last error string.
func ClearLastError() {
	C.typr_io_clear_last_error()
}

// getLastError is an internal helper that returns the last error or a fallback message.
func getLastError(fallback string) error {
	if errStr := GetLastError(); errStr != "" {
		return errors.New(errStr)
	}
	return errors.New(fallback)
}

// freeString is a helper to free C strings.
func freeString(s *C.char) {
	C.typr_io_free_string(s)
}

// cString allocates a C string that must be freed with C.free.
func cString(s string) *C.char {
	return C.CString(s)
}

// freeCString frees a C string allocated by cString.
func freeCString(s *C.char) {
	C.free(unsafe.Pointer(s))
}
