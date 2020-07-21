package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var input = `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,19,9,23,1,5,23,27,1,27,9,31,1,6,31,35,2,35,9,39,1,39,6,43,2,9,43,47,1,47,6,51,2,51,9,55,1,5,55,59,2,59,6,63,1,9,63,67,1,67,10,71,1,71,13,75,2,13,75,79,1,6,79,83,2,9,83,87,1,87,6,91,2,10,91,95,2,13,95,99,1,9,99,103,1,5,103,107,2,9,107,111,1,111,5,115,1,115,5,119,1,10,119,123,1,13,123,127,1,2,127,131,1,131,13,0,99,2,14,0,0`

func splitIntoSubArs(ar []int, sliceLen int) [][]int {
	var retval [][]int
	var currentSub []int

	for k := 0; k < len(ar); k++ {
		currentSub = append(currentSub, ar[k])
		if len(currentSub) == sliceLen {
			retval = append(retval, currentSub)
			currentSub = nil
		}
	}
	return retval
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func makeProgram(originalProgram []int, noun, verb int) []int {
	retval := make([]int, len(originalProgram))
	copy(retval, originalProgram)
	retval[1] = noun
	retval[2] = verb
	return retval
}

func nounVerbGenerator() func() (int, int, error) {
	noun, verb := 0, 0
	return func() (int, int, error) {
		if verb == 99 && noun == 99 {
			return 0, 0, errors.New("no more values")
		}
		if verb > 99 {
			verb = 0
			noun++
			return noun, verb, nil
		} else {
			verb++
			return noun, verb - 1, nil
		}
	}
}

func main() {
	var mainProgram []int
	for _, st := range strings.Split(input, ",") {
		parsed, er := strconv.Atoi(st)
		if er != nil {
			panic(er)
		}
		mainProgram = append(mainProgram, parsed)
	}

	generator := nounVerbGenerator()
	for true {
		noun, verb, error := generator()
		fmt.Println(noun, verb)

		if error != nil {
			panic("Run out of attempts")
		}

		programAttempt := makeProgram(mainProgram, noun, verb)

		k := 0
		for true {
			if k >= len(programAttempt) {
				break
			}

			subProg := programAttempt[k:min(k+4, len(programAttempt))]
			// subProg might consist of a single instr

			instruction := subProg[0]
			if instruction == 99 {
				k++
				break
			} else {
				firstPtr, sndPtr, storageLoc := subProg[1], subProg[2], subProg[3]
				if instruction == 1 {
					// add
					programAttempt[storageLoc] = programAttempt[firstPtr] + programAttempt[sndPtr]
				}
				if instruction == 2 {
					// multiply
					programAttempt[storageLoc] = programAttempt[firstPtr] * programAttempt[sndPtr]
				}
				k = k + 4
			}
		}
		fmt.Println(programAttempt[0])
		if programAttempt[0] == 19690720 {
			fmt.Println(100*noun + verb)
			return
		}
	}

}
