package main

import "testing"

var testArrayEncoded []byte
var testArray []Value

func init() {
	testArrayEncoded = make([]byte, valueSize*10_000)
	for i := 0; i < 10_000; i++ {
		copy(testArrayEncoded[i*valueSize:], testArrayEncoded)
	}

	testArray = decodeArray(testArrayEncoded) // :))
}

func BenchmarkArrayEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = encodeArray(testArray)[0]
	}
}

func BenchmarkArrayDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = decodeArray(testArrayEncoded)[0]
	}
}

func BenchmarkPointerArraypDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = decodePointerArray(testArrayEncoded)[0]
	}
}
