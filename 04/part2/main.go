package main

import (
	"fmt"
	"strconv"
)

var pwdLower = 136760
var pwdUpper = 595730

func allTrue(rules ...func(int) bool) func(int) bool {
	return func(y int) bool {
		for _, rule := range rules {
			if !rule(y) {
				return false
			}
		}
		return true
	}
}

func adjacentDigitsMeetCriterion(x int) bool {
	intStr := strconv.Itoa(x)
	byteMap := make(map[byte][]int)
	for ind, b := range []byte(intStr) {
		byteMap[b] = append(byteMap[b], ind)
	}
	// check if rules are met
	for _, val := range byteMap {
		if len(val) == 2 {
			// check if adjacent
			if (val[1] - val[0]) == 1 {
				return true
			}

		}
	}
	return false
}

func allDigitsRising(x int) bool {
	intStr := strconv.Itoa(x)
	strByte := []byte(intStr)
	for k := 0; k < len(strByte)-1; k++ {
		if strByte[k] > strByte[k+1] {
			return false
		}
	}
	return true
}

func main() {
	matching := 0
	rulesChecker := allTrue(adjacentDigitsMeetCriterion, allDigitsRising)
	for k := pwdLower; k <= pwdUpper; k++ {
		if rulesChecker(k) {
			matching++
		}

	}

	fmt.Print(matching)

}
