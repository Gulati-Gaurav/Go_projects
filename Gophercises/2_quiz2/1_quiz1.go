package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// timer package only triggers once. ticker package triggers regularly

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
	c := make(chan string)
	timer := time.NewTimer(time.Duration(3 * time.Second))

	for _, calc := range data {
		fmt.Print(calc[0], " ")
		go func() {
			var resp int
			fmt.Scanln(&resp)
			ans, err := strconv.Atoi(calc[1])
			if err == nil && ans == resp {
				score++
			}
			c <- "received"
		}()

		select {
		case <-timer.C:
			fmt.Println("Time Over. Your final score = ", score)
			return
		case <-c:
			continue
		}
	}
	fmt.Println("Your final score = ", score)
}
