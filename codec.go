package main

import (
	"reflect"
	"unsafe"
)

func (v *Value) encode() []byte {
	return (*(*[]byte)(unsafe.Pointer(&v)))[:VALUE_SIZE]
}

func decode(b []byte) *Value {
	return ((*Value)(unsafe.Pointer(&b[0])))
}

func encodeArray(a []Value) []byte {
	sh := reflect.SliceHeader{
		Data: uintptr((unsafe.Pointer(&a[0]))),
		Len:  len(a) * VALUE_SIZE,
		Cap:  len(a) * VALUE_SIZE,
	}
	return (*(*[]byte)(unsafe.Pointer(&sh)))
}

func decodeArray(b []byte) []Value {
	sh := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&b[0])),
		Len:  len(b) / VALUE_SIZE,
		Cap:  len(b) / VALUE_SIZE,
	}
	return *(*[]Value)(unsafe.Pointer(&sh))
}
