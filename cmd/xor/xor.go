package main

import (
	"fmt"
	"os"

	"github.com/s77rt/xor"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: xor input key")
		os.Exit(1)
	}

	input := os.Args[1]
	key := os.Args[2]
	output := xor.XOR([]byte(input), []byte(key))

	fmt.Printf("%s", string(output))
}
