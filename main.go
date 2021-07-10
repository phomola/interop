package interop

import (
	"reflect"
	"unsafe"
)

func AllocBytes(size int) []byte {
	return unsafe.Slice((*byte)(Alloc(size)), size)
}

func FreeBytes(b []byte) {
	Free(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&b)).Data))
}

func AllocStruct(p interface{}) {
	v1 := reflect.ValueOf(p)
	if v1.Kind() == reflect.Ptr {
		v2 := reflect.Indirect(v1)
		if v2.Kind() == reflect.Ptr {
			t := v2.Type().Elem()
			p := Alloc(int(t.Size()))
			b := unsafe.Slice((*byte)(p), t.Size())
			for i := range b {
				b[i] = 0
			}
			v2.Set(reflect.NewAt(t, p))
			return
		}
	}
	panic("AllocStruct argument must be a pointer to a pointer")
}

