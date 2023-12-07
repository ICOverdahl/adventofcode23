package day06

import (
	"log"
	"os"
)

func Solve() {
	input, err := os.ReadFile("./inputs/006.txt")
	if err != nil {
		log.Fatalf("File read failed: %v", err)
	}
	uno := part1(input)
	dos := part2(input)
	log.Printf("Day 6\nPart 1: %d\nPart 2: %d\n\n", uno, dos)
}

func part1(input []byte) int {
	start := 0
	nums := []int{}
	for end := range input {
		if isDigit(input[start]) == isDigit(input[end]) {
			continue
		}
		if !isDigit(input[start]) && isDigit(input[end]) {
			start = end
			continue
		}
		nums = append(nums, readNum(input, start, end))
		start = end
	}
	times := nums[:len(nums)/2]
	distances := nums[len(nums)/2:]
	product := 1
	for i := range times {
		wins := countNumWins(times[i], distances[i])
		if 0 < wins {
			product *= wins
		}
	}
	return product
}

func part2(input []byte) int {
	var time, record int
	n := 0
	for i := range input {
		if isDigit(input[i]) {
			n = (10 * n) + int(input[i]-'0')
		}
		if input[i] == '\n' && time == 0 {
			time = n
			n = 0
		}
	}
	record = n
	return countNumWins(time, record)
}

func countNumWins(duration, record int) int {
	boatCharge := 0
	for ; boatCharge*(duration-boatCharge) <= record; boatCharge++ {
	}
	// log.Printf("%d %d %d", boatCharge, boatCharge*(duration-boatCharge), record)
	// log.Printf("(%d - %d) * 2 + %d", (duration+1)/2, boatCharge, (duration&1)^1)
	return ((((duration + 1) / 2) - boatCharge) * 2) + (duration & 1) ^ 1
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

func isLetter(c byte) bool {
	return !(c < 'a' || 'z' < c)
}
