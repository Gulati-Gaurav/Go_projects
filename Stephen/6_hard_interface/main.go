package main

import (
	"io"
	"os"
)

func main() {
	fileName := os.Args[1] // The first string is the location of compiled code when we run the file. Then onwards what we pass to the code.
	file, _ := os.Open(fileName)
	io.Copy(os.Stdout, file)
}
