package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Open csv file
	file, err := os.Open("input.csv")
	if err != nil {
		panic("open error")
	}
	defer file.Close()
	fmt.Println("open complete.")

	// Create a CSV reader
	reader := csv.NewReader(file)
	fmt.Println("read complete.")

	// Define empty slice
	var list []int

	// Infinite loop
	for {
		//  Parse content of CSV
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

	// Multiply 2 numbers if sum = 2020
	var i int
	var j int
	var result int

	for i = 0; i < len(list); i++ {
		for j = 0 + i; j < len(list); j++ {
			if list[i]+list[j] == 2020 {
				result = list[i] * list[j]
				break
			}
		}
	}

	fmt.Printf("Result part 1: %d\n", result)

}
