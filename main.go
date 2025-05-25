package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("./messages.txt")
	if err != nil {
		fmt.Printf("Error opening messages.txt: %s", err)
		return
	}
	defer f.Close()

	inStream := make([]byte, 8)

read:
	for {
		n, err := f.Read(inStream)
		if err != nil {
			break read
		}

		if n != 8 {
			fmt.Printf("read: %s\n", inStream[:(n-1)])
		} else {
			fmt.Printf("read: %s\n", inStream)
		}
	}
}
