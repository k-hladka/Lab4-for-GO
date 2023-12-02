package main

import "github.com/andlabs/ui"

func main() {
	err := ui.Main(Travel)
	if err == nil {
		panic(err)
	}
}
