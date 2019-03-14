#include "textflag.h"

// This code taken from inserting the following into
// the runtime package, compiling it, and inspecting it
// with go tool objdump -s foo
//
//    //go:noinline
//    func bar(pc uintptr) string {
//           return ""
//    }
//    
//    //go:noinline
//    func foo() string {
//           return bar(getcallerpc())
//    }

TEXT ·This(SB),0,$24-16
	MOVQ addr-8(FP), AX
	MOVQ AX, 0(SP)
	CALL ·Name(SB)
	MOVQ 16(SP), AX
	MOVQ 8(SP), CX
	MOVQ CX, ret_base+0(FP)
	MOVQ AX, ret_len+8(FP)
	RET
