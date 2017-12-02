package main

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	// Read the file into memory (catch errors)
	file, err := os.Open("C:\\Users\\Michael\\Documents\\test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Setup a diffs variable to store the max - min values of each line
	var diffs []int

	// Scan the file for processing then get the diff of max and min in each line and append to the diffs slice (catch error with scanner)
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		strCol := strings.Fields(scanner.Text())
		intSlice := ToNum(strCol)
		min, max := MinMax(intSlice)
		diff := max - min
		diffs = append(diffs, diff)
	}

	// Sum the each diff from the diffs slice
	sum := 0
	for _, num := range diffs {
		sum += num
	}

	// Print the answer to the screen
	fmt.Println(sum)
}

// Get the min and max of slice of integers
func MinMax(array []int) (int, int) {
	max := array[0]
	min := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

// Translate a slice of strings to a slice of integers
func ToNum(array []string) ([]int) {
	var ints []int
	for _, value := range array {
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, i)
	}
	return ints
}
