package main

import (
	"bytes"
	"testing"
)

func TestConvert(t *testing.T) {
	var text = "Hello World!"
	textAsBytesZeroAlloc := stob(text)
	textAsBytesAlloc := []byte(text)

	//lint:ignore S1004 bytes.Equal is converting bytes to string and comparing. However, we don't want to re-allocate anything.
	if bytes.Compare(textAsBytesZeroAlloc[:], textAsBytesAlloc) != 0 {
		t.Fatalf("unexcept stob result: %s", textAsBytesZeroAlloc)
	}

}

func TestString(t *testing.T) {
	raw := [64]byte{65, 67, 66} // 65,67,66,0,0,0....0
	if btos(raw[:]) != "ACB" {
		t.Fatalf("unexcept btos result: %s", btos(raw[:]))
	}
}
