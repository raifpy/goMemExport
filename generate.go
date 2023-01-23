package main

import (
	"time"

	"github.com/brianvoe/gofakeit"
)

func init() {
	gofakeit.Seed(time.Now().Unix())
}

func generate(to int) []Value {
	list := make([]Value, to)
	for i := range list {
		list[i] = Value{
			private:   gofakeit.Uint32(),
			Zamazingo: stob64(gofakeit.Address().Address),
			Confirm:   gofakeit.Bool(),
		}
	}

	return list
}
