package axidevio

/*
#cgo CFLAGS: -I${SRCDIR}/include

#include <axidev-io/c_api.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

// LibraryVersion returns the axidev-io library version string.
func LibraryVersion() string {
	return C.GoString(C.axidev_io_library_version())
}

// GetLastError returns the last error string from the library, if any.
func GetLastError() string {
	cStr := C.axidev_io_get_last_error()
	if cStr == nil {
		return ""
	}
	defer C.axidev_io_free_string(cStr)
	return C.GoString(cStr)
}

// ClearLastError clears the last error string.
func ClearLastError() {
	C.axidev_io_clear_last_error()
}

// GetLastErrorOrDefault returns the last error or a fallback message as an error.
// This is useful for wrapping C API calls that may set a global error.
func GetLastErrorOrDefault(fallback string) error {
	if errStr := GetLastError(); errStr != "" {
		return errors.New(errStr)
	}
	return errors.New(fallback)
}

// freeString is a helper to free C strings.
func freeString(s *C.char) {
	C.axidev_io_free_string(s)
}

// cString allocates a C string that must be freed with C.free.
func cString(s string) *C.char {
	return C.CString(s)
}

// freeCString frees a C string allocated by cString.
func freeCString(s *C.char) {
	C.free(unsafe.Pointer(s))
}
