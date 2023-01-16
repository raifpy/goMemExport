package main

type Value struct {
	Zamazingo [64]byte // 64
	private   uint32   // 4 byte
	Confirm   bool     // 1-(4 with padding) byte

}

const VALUE_SIZE = 72
