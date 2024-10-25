package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"
)

func main() {

	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(homeDir)
}
