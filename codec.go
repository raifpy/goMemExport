package main

import (
	"reflect"
	"unsafe"
)

func encode(v *Value) []byte {
	return (*(*[]byte)(unsafe.Pointer(&v)))
}

func decode(b []byte) *Value {
	return ((*Value)(unsafe.Pointer(&b[0])))
}

func encodeArray(a []Value) []byte {
	sh := reflect.SliceHeader{
		Data: uintptr((unsafe.Pointer(&a[0]))),
		Len:  len(a) * valueSize,
		Cap:  len(a) * valueSize,
	}
	return (*(*[]byte)(unsafe.Pointer(&sh)))
}

func decodeArray(b []byte) []Value {

	sh := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&b[0])),
		Len:  len(b) / valueSize,
		Cap:  len(b) / valueSize,
	}
	return *(*[]Value)(unsafe.Pointer(&sh))
}

// Allocs pointer array of size 8* len(b)/valueSize
// Useless
func decodePointerArray(b []byte) []*Value {
	arr := make([]*Value, len(b)/valueSize)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = decode(b[i*valueSize:])
	}
	return arr
}
