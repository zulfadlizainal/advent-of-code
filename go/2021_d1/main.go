package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Open txt file
	file, err := os.Open("input.txt")
	if err != nil {
		panic("open error")
	}
	defer file.Close()
	fmt.Println("open complete.")

	// Create a txt reader
	reader := csv.NewReader(file)
	fmt.Println("read complete.")

	// Define empty slice
	var list []int

	// Infinite loop
	for {
		//  Parse content of txt
		record, err := reader.Read()

		if err != nil {
			fmt.Println("parse complete.")
			break
		}

		// Convert string content to integer
		num, err := strconv.Atoi(record[0])

		if err != nil {
			panic("convert error")
		}

		// Save integer in a slice (array)
		list = append(list, num)
	}

	var i int = 0
	var result int = 0

	for i = 0; i < len(list)-1; i++ {
		if list[i+1] > list[i] {
			result++
		}
	}

	fmt.Printf("Result part 1: %d\n", result)

}
