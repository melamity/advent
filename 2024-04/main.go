package main

import (
	"advent/common"
	"fmt"
	"os"
	"strings"
)

func chk_e(ws [][]string, x, y int, target string) bool {
	if x > len(ws[y])-len(target) {
		return false
	}

	results := make([]bool, len(target))
	for i := 0; i < len(target); i++ {
		results[i] = ws[y][x+i] == string(target[i])
	}

	return common.Count(results, func(x bool) bool { return x }) == len(target)
}

func chk_n(ws [][]string, x, y int, target string) bool {
	if y < len(target)-1 {
		return false
	}

	results := make([]bool, len(target))
	for i := 0; i < len(target); i++ {
		results[i] = ws[y-i][x] == string(target[i])
	}

	return common.Count(results, func(x bool) bool { return x }) == len(target)
}

func chk_ne(ws [][]string, x, y int, target string) bool {
	if y < len(target)-1 || x > len(ws[y])-len(target) {
		return false
	}

	results := make([]bool, len(target))
	for i := 0; i < len(target); i++ {
		results[i] = ws[y-i][x+i] == string(target[i])
	}

	return common.Count(results, func(x bool) bool { return x }) == len(target)
}

func chk_nw(ws [][]string, x, y int, target string) bool {
	if y < len(target)-1 || x < len(target)-1 {
		return false
	}

	results := make([]bool, len(target))
	for i := 0; i < len(target); i++ {
		results[i] = ws[y-i][x-i] == string(target[i])
	}

	return common.Count(results, func(x bool) bool { return x }) == len(target)
}

func chk_w(ws [][]string, x, y int, target string) bool {
	if x < len(target)-1 {
		return false
	}

	results := make([]bool, len(target))
	for i := 0; i < len(target); i++ {
		results[i] = ws[y][x-i] == string(target[i])
	}

	return common.Count(results, func(x bool) bool { return x }) == len(target)
}

func chk_s(ws [][]string, x, y int, target string) bool {
	if y > len(ws)-len(target) {
		return false
	}

	results := make([]bool, len(target))
	for i := 0; i < len(target); i++ {
		results[i] = ws[y+i][x] == string(target[i])
	}

	return common.Count(results, func(x bool) bool { return x }) == len(target)
}

func chk_sw(ws [][]string, x, y int, target string) bool {
	if y > len(ws)-len(target) || x < len(target)-1 {
		return false
	}

	results := make([]bool, len(target))
	for i := 0; i < len(target); i++ {
		results[i] = ws[y+i][x-i] == string(target[i])
	}

	return common.Count(results, func(x bool) bool { return x }) == len(target)
}

func chk_se(ws [][]string, x, y int, target string) bool {
	if y > len(ws)-len(target) || x > len(ws[y])-len(target) {
		return false
	}

	results := make([]bool, len(target))
	for i := 0; i < len(target); i++ {
		results[i] = ws[y+i][x+i] == string(target[i])
	}

	return common.Count(results, func(x bool) bool { return x }) == len(target)
}

func chk_x_mas(ws [][]string, x, y int) bool {
	if (x == 0 || y == 0) || (x == len(ws[y])-1 || y == len(ws)-1) {
		return false
	}

	if chk_ne(ws, x-1, y+1, "MAS") && chk_nw(ws, x+1, y+1, "MAS") {
		return true
	}

	if chk_se(ws, x-1, y-1, "MAS") && chk_sw(ws, x+1, y-1, "MAS") {
		return true
	}

	if chk_se(ws, x-1, y-1, "MAS") && chk_ne(ws, x-1, y+1, "MAS") {
		return true
	}

	if chk_sw(ws, x+1, y-1, "MAS") && chk_nw(ws, x+1, y+1, "MAS") {
		return true
	}

	return false
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

	mases := 0
	xmases := 0
	ws := make([][]string, 0)
	for line := range ch {
		ws = append(ws, strings.Split(strings.Trim(line, " "), ""))
	}

	for y, line := range ws {
		for x, char := range line {
			if char == "X" {
				xmases += common.Count([]bool{
					chk_n(ws, x, y, "XMAS"),
					chk_ne(ws, x, y, "XMAS"),
					chk_e(ws, x, y, "XMAS"),
					chk_se(ws, x, y, "XMAS"),
					chk_s(ws, x, y, "XMAS"),
					chk_sw(ws, x, y, "XMAS"),
					chk_w(ws, x, y, "XMAS"),
					chk_nw(ws, x, y, "XMAS"),
				}, func(x bool) bool { return x })
			} else if char == "A" {
				if chk_x_mas(ws, x, y) {
					mases++
				}
			}
		}
	}

	fmt.Println("Part 1 Result:", xmases)
	fmt.Println("Part 2 Result:", mases)
}
