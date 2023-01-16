package main

import "os"

func write() error {
	return os.WriteFile(*path, encodeArray(generate(*size)), 0666)
}
