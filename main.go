package main

import (
	"fmt"
	lib "github.com/padamstx/gomodlib"
)

const version string = "1.0.0"

func main() {
	fmt.Printf("modapp main, version %s\n", version)
	fmt.Printf("   message from gomodlib: %s\n", lib.HelloWorld())
}