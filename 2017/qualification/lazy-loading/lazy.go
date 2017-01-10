package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

const input string = "lazy_loading.txt"

func main() {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		lines := strings.Split(string(data), "\n")
		itemCountOffset := 0
		lastLength := 0
		caseNum := 1
		for i := 1; i < len(lines); i++ {
			if itemCountOffset == 0 {
				itemCountOffset, _ = strconv.Atoi(lines[i])
				if lastLength > 0 {
					result := processSlice(lines[i-lastLength : i])
					fmt.Printf("Case #%d: %s\n", caseNum, result)
					caseNum++
				}
				lastLength = itemCountOffset
			} else {
				itemCountOffset--
			}
		}
	}
}

func processSlice(slice []string) string {
	// Implement solution here
	s := toInts(slice)
	sort.Ints(s)

	boxesUsed := 0
	trips := 0
	size := len(s)

	for i := size; i > 0 && boxesUsed < size; i-- {
		top := s[i-1]
		numRequired := int(math.Ceil(float64(50) / float64(top)))
		boxesUsed += numRequired
		// If we would have used more boxes than are available then they would have to go on the previous trip
		if boxesUsed <= size {
			trips++
		}
	}
	return strconv.Itoa(trips)
}

func toInts(strSlice []string) []int {
	size := len(strSlice)
	intSlice := make([]int, size)
	for i := 0; i < size; i++ {
		intSlice[i], _ = strconv.Atoi(strSlice[i])
	}
	return intSlice
}
