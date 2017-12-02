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

	// Setup a divs variable to store the evenly divided portions
	var divs []int

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

		evenQ := DivideEvenly(intSlice)
		divs = append(divs, evenQ)
	}

	// Sum each diff from the diffs slice
	dsum := 0
	for _, num := range diffs {
		dsum += num
	}

	// Sum each evenQ from the divs slice
	qsum := 0
	for _, num := range divs {
		qsum += num
	}

	// Print the answer to the screen
	fmt.Println(dsum)
	fmt.Println(qsum)
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

// Get the evenly divided result of integers
func DivideEvenly(array []int) (int) {
	de := 0
	first := array
	second := array
	for _, v1 := range first {
		for _, v2 := range second {
			if v1 != v2 {
				q, r := divmod(v1, v2)
				if r == 0 {
					return q
				}
			}
		}
	}
	return de
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
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
