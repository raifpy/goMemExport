package main

import (
	"flag"
	"fmt"
	"os"
)

var path = flag.String("path", "", "Path to write-read object")
var size = flag.Int("size", 100_000, "write generate object size")

func exit() {
	fmt.Fprintf(os.Stderr, "usage: %s  --path file read|write != (%v)", os.Args[0], os.Args)
	os.Exit(1)
}

func main() {
	flag.Parse()

	fmt.Printf("path: %v\n", *path)

	if *path == "" {
		exit()
	}

	if len(flag.Args()) == 0 {
		exit()
	}

	var err error
	switch flag.Args()[0] {
	case "read":
		err = read()
	case "write":
		err = write()
	default:
		exit()
	}

	if err != nil {
		fmt.Printf("err: %v\n", err)
		exit()
	}
}
