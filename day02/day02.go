package day02

import (
	"log"
	"os"
)

func Solve() {
	input, err := os.ReadFile("./inputs/002.txt")
	if err != nil {
		log.Fatalf("File read failed: %v", err)
	}
	uno := part1(input)
	dos := part2(input)
	log.Printf("Day 2\nPart 1: %d\nPart 2: %d\n\n", uno, dos)
}

func part1(input []byte) int {
	idSum := 0
	var (
		red, green, blue byte = 12, 13, 14

		num, id byte
	)
	possible := true
	for i := 0; i < len(input); i++ {
		c := input[i]
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num = (10 * num) + (c - '0')
		case ':':
			id = num
			num = 0
		case 'r':
			possible = num <= red
			num = 0
			i += 2
		case 'g':
			possible = num <= green
			num = 0
			i += 4
		case 'b':
			possible = num <= blue
			num = 0
			i += 3
		case '\n':
			idSum += int(id)
		}
		if !possible {
			for input[i] != '\n' {
				i++
			}
			possible = true
		}
	}
	return idSum
}

func part2(input []byte) int {
	powerSum := 0
	var (
		num, minRed, minBlue, minGreen byte
	)
	for i := 0; i < len(input); i++ {
		c := input[i]
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num = (10 * num) + (c - '0')
		case ':':
			num = 0
		case 'r':
			if minRed < num {
				minRed = num
			}
			num = 0
			i += 2
		case 'g':
			if minGreen < num {
				minGreen = num
			}
			num = 0
			i += 4
		case 'b':
			if minBlue < num {
				minBlue = num
			}
			num = 0
			i += 3
		case '\n':
			power := int(minRed) * int(minGreen) * int(minBlue)
			powerSum += power
			minRed, minGreen, minBlue = 0, 0, 0
		}
	}
	return powerSum
}
