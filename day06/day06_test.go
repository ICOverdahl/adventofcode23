package day06

import "testing"

func TestCountNumWins(t *testing.T) {
	cases := []struct {
		D, R, Expected int
	}{
		{7, 9, 4},
		{15, 40, 8},
		{30, 200, 9},
	}
	for _, c := range cases {
		actual := countNumWins(c.D, c.R)
		if actual != c.Expected {
			t.Fatalf("%d, %d: expected %d, got %d", c.D, c.R, c.Expected, actual)
		}
	}
}

func TestPart1(t *testing.T) {
	blob := `Time:      7  15   30
Distance:  9  40  200
`
	expected := 288
	actual := part1([]byte(blob))
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	blob := `Time:      7  15   30
Distance:  9  40  200
`
	expected := 71503
	actual := part2([]byte(blob))
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
