package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Open csv file
	file, err := os.Open("input.txt")
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
	var k int
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

	// Multiply 3 numbers if sum = 2020

	for i = 0; i < len(list); i++ {
		for j = 0 + i; j < len(list); j++ {
			for k = 0 + j; k < len(list); k++ {
				if list[i]+list[j]+list[k] == 2020 {
					result = list[i] * list[j] * list[k]
					break
				}
			}
		}
	}

	fmt.Printf("Result part 2: %d\n", result)

}
