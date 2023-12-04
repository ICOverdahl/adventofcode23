package day04

import (
	"log"
	"os"
)

func Solve() {
	input, err := os.ReadFile("./inputs/004.txt")
	if err != nil {
		log.Fatalf("File read failed: %v", err)
	}
	uno := part1(input)
	dos := part2(input)
	log.Printf("Day 4\nPart 1: %d\nPart 2: %d\n\n", uno, dos)
}

func part1(input []byte) int {
	totalPoints := 0
	var (
		big, lil uint64

		bigWin, lilWin uint64
	)
	start := 0
	for end := range input {
		if input[end] == '|' {
			bigWin, lilWin = big, lil
			big, lil = 0, 0
			continue
		}
		if input[end] == ':' || (!isDigit(input[start]) && isDigit(input[end])) {
			start = end
			continue
		}
		if isDigit(input[start]) == isDigit(input[end]) {
			continue
		}
		num := readNum(input, start, end)
		if num < 64 {
			lil |= (1 << num)
		} else {
			big |= (1 << (num - 63))
		}
		if input[end] == '\n' {
			matchingNums := countSetBits(bigWin&big) + countSetBits(lilWin&lil)
			if 0 < matchingNums {
				totalPoints += (1 << (matchingNums - 1))
			}
			big = 0
			lil = 0
		}
		start = end
	}
	return totalPoints
}

func part2(input []byte) int {
	cardCounts := make([]int, 208) // just looked at input file, hopefully that's not cheating
	var (
		big, lil uint64

		bigWin, lilWin uint64

		card int
	)
	start := 0
	for end := range input {
		if !isDigit(input[start]) && isDigit(input[end]) {
			start = end
			continue
		}
		if input[end] == '|' {
			bigWin, lilWin = big, lil
			big, lil = 0, 0
			continue
		}
		if isDigit(input[start]) == isDigit(input[end]) {
			continue
		}
		num := readNum(input, start, end)
		if input[end] == ':' {
			card = num - 1
			cardCounts[card]++
			start = end
			continue
		}
		if num < 64 {
			lil |= (1 << num)
		} else {
			big |= (1 << (num - 63))
		}
		if input[end] == '\n' {
			matchingNums := countSetBits(bigWin&big) + countSetBits(lilWin&lil)
			// log.Printf("Card %d, %d", card, matchingNums)
			for i := card + 1; i <= card+matchingNums && i < len(cardCounts); i++ {
				cardCounts[i] += cardCounts[card]
			}
			big = 0
			lil = 0
		}
		start = end
	}
	sum := 0
	for _, c := range cardCounts {
		sum += c
	}
	return sum
}

func countSetBits(n uint64) int {
	count := 0
	for n > 0 {
		n &= n - 1
		count++
	}
	return count
}

func readNum(input []byte, start, end int) int {
	num := 0
	for _, c := range input[start:end] {
		num = (10 * num) + int(c-'0')
	}
	return num
}

func isDigit(c byte) bool {
	return !(c < '0' || '9' < c)
}
