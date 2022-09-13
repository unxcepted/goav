package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/log.h>
/*
#include <stdio.h>
extern void goAvLogCustomCallback(int level, char* msg, char* parent);

static inline void avLogCustomCallback(void *avcl, int level, const char *fmt, va_list vl)
{
	if (level > av_log_get_level()) return;
	AVClass* avc = avcl ? *(AVClass **) avcl : NULL;
	char parent[1024];
	if (avc) {
		sprintf(parent, "%p", avcl);
	}
	char msg[1024];
	vsprintf(msg, fmt, vl);
	goAvLogCustomCallback(level, msg, parent);
}
static inline void setAvLogCustomCallback()
{
	av_log_set_callback(avLogCustomCallback);
}
static inline void resetAvLogCallback()
{
	av_log_set_callback(av_log_default_callback);
}
*/
import "C"

// Logging constants
const (
	AV_LOG_QUIET   = C.AV_LOG_QUIET
	AV_LOG_PANIC   = C.AV_LOG_PANIC
	AV_LOG_FATAL   = C.AV_LOG_FATAL
	AV_LOG_ERROR   = C.AV_LOG_ERROR
	AV_LOG_WARNING = C.AV_LOG_WARNING
	AV_LOG_INFO    = C.AV_LOG_INFO
	AV_LOG_VERBOSE = C.AV_LOG_VERBOSE
	AV_LOG_DEBUG   = C.AV_LOG_DEBUG
)

// AvLogGetLevel returns the current log level.
func AvLogGetLevel() int {
	return int(C.av_log_get_level())
}

// AvLogSetLevel sets the log level.
func AvLogSetLevel(level int) {
	C.av_log_set_level(C.int(level))
}

// AvLogCallback represents a log callback
type AvLogCallback func(level int, msg, parent string)

var avLogCallback AvLogCallback

// AvLogSetCallback sets the log callback to a custom callback
func AvLogSetCallback(c AvLogCallback) {
	avLogCallback = c
	C.setAvLogCustomCallback()
}

//export goAvLogCustomCallback
func goAvLogCustomCallback(level C.int, msg, parent *C.char) {
	if avLogCallback == nil {
		return
	}
	avLogCallback(int(level), C.GoString(msg), C.GoString(parent))
}

// AvLogResetCallback resets the log callback to the default callback
func AvLogResetCallback() {
	C.resetAvLogCallback()
}
