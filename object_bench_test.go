package main

import "testing"

var testObject = Value{
	Zamazingo: stob64("Greetings from Go!"),
}
var testObjectEncoded = encode(&testObject)

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = encode(&testObject)[0]
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = decode(testObjectEncoded).Zamazingo[0]
	}
}
