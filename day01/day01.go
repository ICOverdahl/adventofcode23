package day01

import (
	"log"
	"os"
)

func Solve() {
	input, err := os.ReadFile("./inputs/001.txt")
	if err != nil {
		log.Fatalf("File read failed: %v", err)
	}
	uno := Part1(input)
	dos := Part2(input)
	log.Printf("Day 1\nPart 1: %d\nPart 2: %d\n\n", uno, dos)
}

func Part1(input []byte) int {
	var last byte = '\n'
	sum := 0
	for _, c := range input {
		switch c {
		case '\n':
			if last < 10 {
				sum += int(last)
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			c -= '0'
			if last == '\n' {
				sum += 10 * int(c)
			}
		default:
			continue
		}
		last = c
	}
	return sum
}

func Part2(input []byte) int {
	var last byte = '\n'
	sum := 0
	for i, c := range input {
		if c == '\n' && last < 10 {
			sum += int(last)
			last = c
			continue
		}
		var val byte = 0xFF
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val = c - '0'
		case 'o': // one
			if i+3 < len(input) && string(input[i:i+3]) == "one" {
				val = 1
			}
		case 'e': // eight
			if i+5 < len(input) && string(input[i:i+5]) == "eight" {
				val = 8
			}
		case 'n': // nine
			if i+4 < len(input) && string(input[i:i+4]) == "nine" {
				val = 9
			}
		case 't': // two three
			if i+3 < len(input) && string(input[i:i+3]) == "two" {
				val = 2
			} else if i+5 < len(input) && string(input[i:i+5]) == "three" {
				val = 3
			}
		case 'f': // four five
			if i+4 < len(input) && string(input[i:i+4]) == "four" {
				val = 4
			} else if i+4 < len(input) && string(input[i:i+4]) == "five" {
				val = 5
			}
		case 's': // six seven
			if i+3 < len(input) && string(input[i:i+3]) == "six" {
				val = 6
			} else if i+5 < len(input) && string(input[i:i+5]) == "seven" {
				val = 7
			}
		}
		if val == 0xFF {
			continue
		}
		if last == '\n' {
			sum += 10 * int(val)
		}
		last = val
	}
	return sum
}
