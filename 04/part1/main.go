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

func adjacentDigitsTheSame(x int) bool {
	intStr := strconv.Itoa(x)
	for k := 0; k < len(intStr)-1; k++ {
		if intStr[k] == intStr[k+1] {
			return true
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
	rulesChecker := allTrue(adjacentDigitsTheSame, allDigitsRising)
	for k := pwdLower; k <= pwdUpper; k++ {
		if rulesChecker(k) {
			matching++
		}

	}

	fmt.Print(matching)

}
