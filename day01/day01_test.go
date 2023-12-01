package day01

import "testing"

func TestPart1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`
	expected := 142
	actual := Part1([]byte(input))
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`
	expected := 281
	actual := Part2([]byte(input))
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
