package main

import (
	"reflect"
	"unsafe"
)

func stob(s string) [64]byte {
	if len(s) >= 64 {
		return *(*[64]byte)(unsafe.Pointer(&s))
	}
	var b [64]byte
	copy(b[:], s)
	return b
}

func btos(b [64]byte) string {
	var cut = 64
	for i := 63; i >= 0; i-- {
		if b[i] == 0 {
			cut = i
			break
		}
	}
	return (*(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(&b[0])),
		Len:  cut,
	})))
}
