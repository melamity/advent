package main

import (
	"advent/common"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	regex := regexp.MustCompile(`(do|don't|mul)\((?:(\d+),(\d+))?\)`)
	resultOne := 0
	resultTwo := 0
	mulEnable := true
	for line := range ch {
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch match[1] {
			case "do":
				mulEnable = true
				break
			case "don't":
				mulEnable = false
				break
			case "mul":
				arg1, err := strconv.Atoi(match[2])
				if err != nil {
					continue
				}
				arg2, err := strconv.Atoi(match[3])
				if err != nil {
					continue
				}

				resultOne += arg1 * arg2
				if mulEnable {
					resultTwo += arg1 * arg2
				}
				break
			}
		}
	}

	fmt.Println("Part 1 Result:", resultOne)
	fmt.Println("Part 2 Result:", resultTwo)
}
