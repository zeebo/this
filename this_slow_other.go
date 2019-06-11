// This works on any go less than go 1.11

// +build !go1.11

package this

import "runtime"

// This returns the package/function name of the caller.
func This() string {
	pc, _, _, _ := runtime.Caller(1)
	return Name(pc)
}

// ThisN returns the package/function n levels below the caller.
func ThisN(n int) string {
	pc, _, _, _ := runtime.Caller(1 + n)
	return Name(pc)
}
