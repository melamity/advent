package main

import (
	"advent/common"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func test(numbers []int) bool {
	isSortedAsc := sort.SliceIsSorted(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })
	isSortedDesc := sort.SliceIsSorted(numbers, func(i, j int) bool { return numbers[i] > numbers[j] })
	if !(isSortedAsc || isSortedDesc) {
		return false
	}

	for i := 0; i < len(numbers)-1; i++ {
		abs := common.Abs(numbers[i] - numbers[i+1])
		if abs < 1 || abs > 3 {
			return false
		}
	}

	return true
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "filename")
		return
	}

	ch, err := common.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	reportSafety := make([]bool, 0)
	reportSafetyDampening := make([]bool, 0)
	for line := range ch {
		fields := strings.Fields(line)
		numbers := make([]int, len(fields))

		for j, field := range fields {
			if field, err := strconv.Atoi(field); err == nil {
				numbers[j] = field
			}
		}

		isStrictlySafe := test(numbers)
		reportSafety = append(reportSafety, isStrictlySafe)

		// If a report is already safe, it will be even with dampening rules.
		if isStrictlySafe {
			reportSafetyDampening = append(reportSafetyDampening, true)
			continue
		}

		for i := range numbers {
			temp := make([]int, len(numbers))
			copy(temp, numbers)
			temp = temp[:i+copy(temp[i:], temp[i+1:])]

			if test(temp) {
				reportSafetyDampening = append(reportSafetyDampening, true)
				break
			}
		}
	}

	fmt.Println("Part 1 Result:", common.Count(reportSafety, func(i bool) bool {
		return i
	}))
	fmt.Println("Part 2 Result:", common.Count(reportSafetyDampening, func(i bool) bool {
		return i
	}))
}
