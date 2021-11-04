package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) > 1 {
		content, err := os.Open(args[1])
		if err != nil {
			fmt.Println(err)
		}

		io.Copy(os.Stdout, content)
	} else {
		fmt.Println("No file informed")
	}

}
