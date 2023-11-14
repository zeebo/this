package this

import "runtime"

// This returns the package/function name of the caller.
func This() string {
	var pcbuf [1]uintptr
	runtime.Callers(2, pcbuf[:])
	frame, _ := runtime.CallersFrames(pcbuf[:]).Next()
	return frame.Function
}

// ThisN returns the package/function n levels below the caller.
func ThisN(n int) string {
	var pcbuf [1]uintptr
	runtime.Callers(n+2, pcbuf[:])
	frame, _ := runtime.CallersFrames(pcbuf[:]).Next()
	return frame.Function
}
