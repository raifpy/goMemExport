package main

import "unsafe"

type Value struct {
	Zamazingo [64]byte // 64
	private   uint32   // 4 byte
	Confirm   bool     // 1-(4 with padding) byte

}

var valueSize int = int(unsafe.Sizeof(Value{}))
