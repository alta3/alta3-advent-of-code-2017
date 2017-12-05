package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"fmt"
	"sort"
)

func main() {
	// Read the file into memory (catch errors)
	file, err := os.Open("C:\\Users\\Michael\\Documents\\d04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var num int

	// Scan the file for processing then get the total count of good passwords
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		strCol := strings.Fields(scanner.Text())
		if len(strCol) == ListToMapLength(strCol) {
			var wordArrays []string
			for _,v := range strCol {
				wordArrays = append(wordArrays, ToCharArray(v))
			}
			if len(strCol) == ListToMapLength(wordArrays) {
				num +=  1
			}
		}
	}

	fmt.Println(num)
}

func ListToMapLength(list []string) int {
	set := make(map[string]struct{})
	for _, v := range list {
		set[v] = struct{}{}
	}
	return len(set)
}

func ToCharArray(word string) string {
	order := strings.Split(word, "")
	sort.Strings(order)
	return strings.Join(order, "")
}
