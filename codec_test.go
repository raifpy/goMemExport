package main

import (
	"fmt"
	"testing"
)

func TestCODEC(t *testing.T) {
	var v = &Value{
		private:   18,
		Zamazingo: [64]byte{112, 113, 65, 66},
		Confirm:   true,
	}

	encoded := v.encode()
	v2 := decode(encoded)
	fmt.Printf("v2.private: %v\n", v2.private)
	fmt.Printf("pointer of v: %p\npointer of v2: %p", v, v2)
}
