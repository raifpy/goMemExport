package main

import (
	"fmt"
	"os"
	"time"
)

func read() error {
	file, err := os.ReadFile(*path)
	if err != nil {
		return err
	}

	now := time.Now()
	objects := decodeArray(file)
	delay := time.Since(now)
	for i, v := range objects {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v.Zamazingo: %v\n", btos(v.Zamazingo))
		//fmt.Printf("i: %v z: %v p: %v c: %v\n", i, btos(v.Zamazingo), v.private, v.Confirm)
	}

	fmt.Printf("delay: %v\n", delay)

	return nil
}
