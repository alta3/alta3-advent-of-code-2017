package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	file, err := os.Open("C:\\Users\\Michael\\Documents\\d01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ints []int
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		captcha := scanner.Text()
		characters := strings.Split(captcha, "")
		ints = ToNum(characters)
	}

	vsum := 0
	i := 0
	for i < len(ints) {
		if i < len(ints) - 2{
			if ints[i] == ints[i + 1] {
				vsum += ints[i]
			}
		} else {
		    if ints[0] == ints[i] {
		    	vsum += ints[i]
			}
		}
		i += 1
	}

	hsum := 0
	hsize := len(ints) / 2
	i = 0
	for i < hsize {
		if ints[i] == ints[i + hsize] {
			hsum += ints[i]
		}
		i += 1
	}
	hsum = hsum * 2

	fmt.Println(vsum)
	fmt.Println(hsum)
}

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
