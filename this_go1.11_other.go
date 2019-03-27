// +build go1.11,!amd64
// +build go1.11,!386

package this

import _ "unsafe"

//go:linkname callers runtime.callers
//go:noescape
func callers(skip int, pcbuf []uintptr) int

// This returns the package/function name of the caller.
func This() string {
	var pcbuf [1]uintptr
	callers(1, pcbuf[:])
	return Name(pcbuf[0])
}

// ThisN returns the package/function n levels below the caller.
func ThisN(n int) string {
	var pcbuf [1]uintptr
	callers(1+n, pcbuf[:])
	return Name(pcbuf[0])
}
