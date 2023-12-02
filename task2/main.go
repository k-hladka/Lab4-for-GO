package main

import (
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(Glass)
	if err == nil {
		panic(err)
	}
}
