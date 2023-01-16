package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed from_mem
var embedin []byte

//go:embed file_mem_list
var embedinlist []byte

func TestDecode(t *testing.T) {
	v := decode(embedin)
	fmt.Printf("v.private: %v\n", v.private)
	fmt.Printf("v.Confirm: %v\n", v.Confirm)
	fmt.Printf("v.Zamazingo: %s\n", v.Zamazingo)
}

func TestDecodeList(t *testing.T) {
	vs := decodeArray(embedinlist)
	fmt.Printf("len(vs): %v\n", len(vs))
	for i, v := range vs {
		fmt.Printf("i: %v priv: %v bool: %v v: %s\n\n", i, v.private, v.Confirm, v.Zamazingo)
	}
}
