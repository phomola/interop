//go:build !windows

package interop

import (
	"unsafe"
)

// #include <stdlib.h>
import "C"

func Alloc(size int) unsafe.Pointer {
	return C.malloc(C.ulong(size))
}

func Free(p unsafe.Pointer) {
	C.free(p)
}
