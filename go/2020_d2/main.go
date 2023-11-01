package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValidPasswordPolicyOne(policy string, password string) bool {
	parts := strings.Split(policy, " ")
	minMax := strings.Split(parts[0], "-")
	char := strings.Trim(parts[1], ":")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	charCount := strings.Count(password, char)
	return charCount >= min && charCount <= max // Boolean whether the condition is met
}

func isValidPasswordPolicyTwo(policy string, password string) bool {
	parts := strings.Split(policy, " ")
	pos := strings.Split(parts[0], "-")
	char := strings.Trim(parts[1], ":")
	pos1, _ := strconv.Atoi(pos[0])
	pos2, _ := strconv.Atoi(pos[1])

	return (password[pos1-1:pos1] == char) != (password[pos2-1:pos2] == char) // Boolean whether the condition is met
}

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("open complete.")

	// Use bufio (buffered input output library)
	// Creates a new scanner that wraps the provided file as its input source.
	scanner := bufio.NewScanner(file)
	fmt.Println("bufio complete.")

	validCountPolicyOne := 0
	validCountPolicyTwo := 0

	// Read the input file line by line
	for scanner.Scan() {
		line := scanner.Text()                  // Process this line
		parts := strings.Split(line, ":")       // Split into 2 parts (policy and password)
		policy := strings.TrimSpace(parts[0])   // Take the 1st part and store as policy
		password := strings.TrimSpace(parts[1]) // Take the 2nd part and store as password

		if isValidPasswordPolicyOne(policy, password) { // Simplification from if isValidPasswordPolicyOne(policy, password) == true
			validCountPolicyOne++
		}

		if isValidPasswordPolicyTwo(policy, password) {
			validCountPolicyTwo++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of valid passwords (Part 1): %d\n", validCountPolicyOne)
	fmt.Printf("Number of valid passwords (Part 2): %d\n", validCountPolicyTwo)
}
