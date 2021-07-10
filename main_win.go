//go:build windows

package interop

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

func Alloc(size int) unsafe.Pointer {
	p, err := windows.LocalAlloc(windows.LMEM_FIXED, uint32(size))
	if err != nil {
		panic(err)
	}
	return unsafe.Pointer(p)
}

func Free(p unsafe.Pointer) {
	_, err := windows.LocalFree(windows.Handle(p))
	if err != nil {
		panic(err)
	}
}
