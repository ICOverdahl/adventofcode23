package day03

import (
	"os"
	"testing"
)

func TestDistinctSymbols(t *testing.T) {
	input, _ := os.ReadFile("../inputs/003.txt")
	chars := []byte{}
	for _, c := range input {
		new := true
		for _, s := range chars {
			if c == s {
				new = false
				break
			}
		}
		if new {
			chars = append(chars, c)
		}
	}
	t.Logf("%s", string(chars))
}
