package main

import (
	"fmt"
	"os"
)

func XOR(input []byte, key []byte) (output []byte) {
	output = make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%len(key)]
	}
	return output
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: xor input key")
		os.Exit(1)
	}

	input := os.Args[1]
	key := os.Args[2]
	output := XOR([]byte(input), []byte(key))

	fmt.Printf("%s", string(output))
}
