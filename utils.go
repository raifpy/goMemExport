package main

import (
	"bytes"
	"reflect"
	"unsafe"
)

// zero allocation convert
func stob(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// allocates 64byte and copies string into
func stob64(s string) (b [64]byte) {
	copy(b[:], stob(s))
	return
}

func btos(b []byte) string {
	var cut = bytes.IndexByte(b, 0) // Go Strings are not like C strings. 0 (\0) is not just end of string, it is also breaking.
	if cut == -1 {
		cut = 64
	}
	return (*(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(&b[0])),
		Len:  cut,
	})))
}
