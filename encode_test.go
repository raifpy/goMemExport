package main

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

var defaultEncodeTest = &Value{
	private:   34,
	Confirm:   true,
	Zamazingo: [64]byte{65, 66, 67, 68, 0},
}

func TestEncode(t *testing.T) {
	vbytes := defaultEncodeTest.encode()
	fmt.Printf("vbytes: %s\n", vbytes)
	// IGNORE BROKE STRING

	list := make([]Value, 12)
	for i := range list {
		list[i] = Value{
			private:   uint32(10*i + rand.Intn(10)),
			Zamazingo: [64]byte{byte(i)},
			Confirm:   i%2 == 0,
		}
	}
	v := encodeArray(list)
	fmt.Printf("len(v): %v\n", len(v))
	fmt.Printf("v: %v\n", v)
	fmt.Printf("list p: %p v 0 p: %p", list, &v[0])

	os.WriteFile("file_mem_list", v, 0666)
}

func TestEncodeAndWrite(t *testing.T) {
	if err := os.WriteFile("file_mem", defaultEncodeTest.encode(), 0666); err != nil {
		t.Fatal(err)
	}
}
