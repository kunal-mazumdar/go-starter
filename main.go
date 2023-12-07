package main

import (
	"fmt"
	"github.com/otiai10/copy"
)

func main() {
	err := copy.Copy("/tmp/a.txt", "/tmp/b.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("Test 1")
}
