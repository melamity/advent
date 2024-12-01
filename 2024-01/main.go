package main

import (
	"advent/common"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "filename")
		return
	}

	ch, err := common.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	list := make([][]int, 2)
	for i := range list {
		list[i] = make([]int, 0)
	}

	for line := range ch {
		for index, input := range strings.Fields(line) {
			number, err := strconv.Atoi(input)
			if err != nil {
				continue
			}

			list[index] = append(list[index], number)
		}
	}

	for i := range list {
		slices.Sort(list[i])
	}

	partOneResult := 0
	for i := 0; i < len(list[0]); i++ {
		partOneResult += max(list[0][i], list[1][i]) - min(list[0][i], list[1][i])
	}
	fmt.Printf("Part 1 Result: %d\n", partOneResult)

	partTwoResult := 0
	for _, v := range list[0] {
		partTwoResult += v * common.Count(list[1], func(cmp int) bool {
			return cmp == v
		})
	}
	fmt.Printf("Part 2 Result: %d\n", partTwoResult)

	return
}
