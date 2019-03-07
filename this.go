package this

import (
	"runtime"
	"strings"
	"unsafe"
)

type funcInfo struct {
	*runtime.Func
	_ unsafe.Pointer
}

type inlinedCall struct {
	_  [12]byte
	fn int32
	_  [4]byte
}

//go:linkname callers runtime.callers
//go:noescape
func callers(n int, pcs []uintptr) int

//go:linkname findfunc runtime.findfunc
func findfunc(pc uintptr) funcInfo

//go:linkname funcname runtime.funcname
func funcname(f funcInfo) string

//go:linkname funcdata runtime.funcdata
func funcdata(f funcInfo, i uint8) unsafe.Pointer

//go:linkname pcdatavalue runtime.pcdatavalue
func pcdatavalue(f funcInfo, table int32, targetpc uintptr, cache unsafe.Pointer) int32

//go:linkname funcnameFromNameoff runtime.funcnameFromNameoff
func funcnameFromNameoff(f funcInfo, nameoff int32) string

// This returns the package/function name being called.
func This() string {
	// acquire caller info
	var pc [1]uintptr
	callers(1, pc[:])
	info := findfunc(pc[0])

	// adjust pc if necessary
	if pc[0] > info.Entry() {
		pc[0]--
	}

	// attempt to determine name, walking inlining data
	name := funcname(info)
	if inldata := funcdata(info, 2); inldata != nil {
		inltree := (*[1 << 20]inlinedCall)(inldata)
		ix := pcdatavalue(info, 1, pc[0], nil)
		if ix >= 0 {
			name = funcnameFromNameoff(info, inltree[ix].fn)
		}
	}

	return strings.TrimSuffix(name, ".init")
}

// ThisN returns the package/function of the caller n frames up, where ThisN(1) == This().
func ThisN(n int) string {
	// acquire caller info
	var pc [1]uintptr
	callers(n, pc[:])
	info := findfunc(pc[0])

	// adjust pc if necessary
	if pc[0] > info.Entry() {
		pc[0]--
	}

	// attempt to determine name, walking inlining data
	name := funcname(info)
	if inldata := funcdata(info, 2); inldata != nil {
		inltree := (*[1 << 20]inlinedCall)(inldata)
		ix := pcdatavalue(info, 1, pc[0], nil)
		if ix >= 0 {
			name = funcnameFromNameoff(info, inltree[ix].fn)
		}
	}

	return strings.TrimSuffix(name, ".init")
}
