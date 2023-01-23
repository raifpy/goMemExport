package main

import (
	"testing"
	"unsafe"
)

func TestCodec(t *testing.T) {
	var v = &Value{
		private:   1 << 31,
		Zamazingo: stob64("Hello World!"),
		Confirm:   true,
	}

	encoded := encode(v)
	v2 := decode(encoded)
	if unsafe.Pointer(v) != unsafe.Pointer(&encoded[0]) || unsafe.Pointer(v2) != unsafe.Pointer(&encoded[0]) {
		t.Fatalf("unexcepted conversion object: %p encoded: %p decoded: %p", v, v2, encoded)
	}
}
