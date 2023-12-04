package day03

import (
	"log"
	"os"
)

func Solve() {
	input, err := os.ReadFile("./inputs/003.txt")
	if err != nil {
		log.Fatalf("File read failed: %v", err)
	}
	uno := part1(input)
	dos := part2(input)
	log.Printf("Day 2\nPart 1: %d\nPart 2: %d\n\n", uno, dos)
}

func part1(input []byte) int {
	sum := 0
	start := 0
	rowLen := getRowLen(input)
	for end := range input {
		if isDigit(input[start]) == isDigit(input[end]) {
			continue
		}
		if !isDigit(input[start]) && isDigit(input[end]) {
			start = end
			continue
		}
		include := isSymbol(input[end]) || (0 < start && isSymbol(input[start-1]))
		for i := start - rowLen - 1; !include && -1 < i && i <= end-rowLen; i++ {
			include = isSymbol(input[i])
		}
		for i := start + rowLen - 1; !include && i < len(input) && i <= end+rowLen; i++ {
			include = isSymbol(input[i])
		}
		if include {
			sum += readNum(input, start, end)
		}
		start = end
	}
	return sum
}

func part2(input []byte) int {
	rowLen := getRowLen(input)
	stars := map[int][]int{}
	start := 0
	for end := range input {
		if isDigit(input[start]) == isDigit(input[end]) {
			continue
		}
		if !isDigit(input[start]) && isDigit(input[end]) {
			start = end
			continue
		}
		num := readNum(input, start, end)
		if input[end] == '*' {
			stars[end] = append(stars[end], num)
		}
		if 0 < start && input[start-1] == '*' {
			stars[start-1] = append(stars[start-1], num)
		}
		for i := start - rowLen - 1; -1 < i && i <= end-rowLen; i++ {
			if input[i] == '*' {
				stars[i] = append(stars[i], num)
			}
		}
		for i := start + rowLen - 1; i < len(input) && i <= end+rowLen; i++ {
			if input[i] == '*' {
				stars[i] = append(stars[i], num)
			}
		}
		start = end
	}
	sum := 0
	for _, adj := range stars {
		if len(adj) == 2 {
			sum += adj[0] * adj[1]
		}
	}
	return sum
}

func isDigit(c byte) bool {
	return !(c < '0' || '9' < c)
}

var SYMBOLS = []byte{'*', '%', '#', '@', '+', '-', '=', '/', '&', '$'}

func isSymbol(c byte) bool {
	for _, s := range SYMBOLS {
		if c == s {
			return true
		}
	}
	return false
}

func readNum(input []byte, start, end int) int {
	num := 0
	for _, c := range input[start:end] {
		num = (10 * num) + int(c-'0')
	}
	return num
}

func getRowLen(input []byte) int {
	for i := range input {
		if input[i] == '\n' {
			return i + 1
		}
	}
	return len(input)
}
