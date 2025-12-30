package axidevio

/*
#include <axidev-io/c_api.h>
*/
import "C"

// LogLevel represents the logging verbosity level.
type LogLevel uint8

// Log level constants matching the C API.
// Levels are ordered from most verbose (LogLevelDebug) to least (LogLevelError).
const (
	LogLevelDebug LogLevel = C.AXIDEV_IO_LOG_LEVEL_DEBUG
	LogLevelInfo  LogLevel = C.AXIDEV_IO_LOG_LEVEL_INFO
	LogLevelWarn  LogLevel = C.AXIDEV_IO_LOG_LEVEL_WARN
	LogLevelError LogLevel = C.AXIDEV_IO_LOG_LEVEL_ERROR
)

// SetLogLevel sets the global logging level.
// Messages at or above the specified level will be displayed;
// lower-priority messages are suppressed.
func SetLogLevel(level LogLevel) {
	C.axidev_io_log_set_level(C.axidev_io_log_level_t(level))
}

// GetLogLevel returns the current global logging level.
func GetLogLevel() LogLevel {
	return LogLevel(C.axidev_io_log_get_level())
}

// IsLogEnabled checks whether messages at a specific level are currently enabled.
func IsLogEnabled(level LogLevel) bool {
	return bool(C.axidev_io_log_is_enabled(C.axidev_io_log_level_t(level)))
}
