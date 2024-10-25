package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	score := 0
	for _, calc := range data {
		fmt.Print(calc[0], " ")
		var resp int
		fmt.Scanln(&resp)
		ans, err := strconv.Atoi(calc[1])
		if err == nil && ans == resp {
			score++
		}
	}
	fmt.Println("Your final score = ", score)
}
