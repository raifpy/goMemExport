package main

import (
	"testing"
	"unsafe"
)

func TestSize(t *testing.T) {
	if unsafe.Sizeof(Value{}) != VALUE_SIZE {
		t.Fatalf("%d != %d", unsafe.Sizeof(Value{}), VALUE_SIZE)
	}
}
