package day07

import (
	"fmt"
	"log"
	"os"
	"slices"
)

func Solve() {
	input, err := os.ReadFile("./inputs/007.txt")
	if err != nil {
		log.Fatalf("File read failed: %v", err)
	}
	uno := Part1(input)
	dos := Part2(input)
	log.Printf("Day 7\nPart 1: %d\nPart 2: %d\n\n", uno, dos)
}

const (
	HIGH_CARD       int = iota // 1 1 1 1 1
	ONE_PAIR                   // 2 1 1 1
	TWO_PAIR                   // 2 2 1
	THREE_OF_A_KIND            // 3 1 1
	FULL_HOUSE                 // 3 2
	FOUR_OF_A_KIND             // 4 1
	FIVE_OF_A_KIND             // 5
)

type Hand struct {
	Cards, Bid int
}

func (h *Hand) Type() int {
	counts := make([]int, 13)
	mask := 0xFF
	for i := 0; i < 5; i++ {
		card := (h.Cards & mask) >> (i * 8)
		counts[card]++
		mask <<= 8
	}
	slices.Sort(counts)
	switch counts[12] {
	case 5:
		return FIVE_OF_A_KIND
	case 4:
		return FOUR_OF_A_KIND
	case 3:
		if counts[11] == 2 {
			return FULL_HOUSE
		}
		return THREE_OF_A_KIND
	case 2:
		if counts[11] == 2 {
			return TWO_PAIR
		}
		return ONE_PAIR
	}
	return HIGH_CARD
}

func cmpHands(h1, h2 *Hand) int {
	if h1.Type()-h2.Type() == 0 {
		return h1.Cards - h2.Cards
	}
	return h1.Type() - h2.Type()
}

func Part1(input []byte) int {
	hands := readHands(input)
	slices.SortFunc(hands, cmpHands)
	sum := 0
	for i := range hands {
		rank := i + 1
		sum += rank * hands[i].Bid
	}
	return sum
}

func readHands(input []byte) []*Hand {
	hands := []*Hand{}
	start := 0
	var cards, bid int
	for end := range input {
		if input[end] == ' ' {
			cards = readCards(input[start:end])
			start = end + 1
		}
		if input[end] == '\n' {
			bid = readNum(input[start:end])
			hands = append(hands, &Hand{Cards: cards, Bid: bid})
			start = end + 1
		}
	}
	return hands
}

func readCards(input []byte) int {
	cards := 0
	for i, c := range input {
		switch c {
		case '2', '3', '4', '5', '6', '7', '8', '9':
			c -= '2'
		case 'T':
			c = 8
		case 'J':
			c = 9
		case 'Q':
			c = 10
		case 'K':
			c = 11
		case 'A':
			c = 12
		}
		cards |= int(c) << ((4 - i) * 8)
	}
	return cards
}

func readNum(input []byte) int {
	num := 0
	for _, c := range input {
		num = (10 * num) + int(c-'0')
	}
	return num
}

// ----------------------------------------

func Part2(input []byte) int {
	hands := readHands2(input)
	slices.SortFunc(hands, cmpHands2)
	sum := 0
	for i := range hands {
		rank := i + 1
		sum += rank * hands[i].Bid
	}
	return sum
}

func (h *Hand) Type2() int {
	counts := make([]int, 13)
	mask := 0xFF
	for i := 0; i < 5; i++ {
		card := (h.Cards & mask) >> (i * 8)
		counts[card]++
		mask <<= 8
	}
	jokers := counts[0]
	fmt.Printf("%v\n", counts)
	slices.Sort(counts[1:])
	fmt.Printf("%v\n", counts)
	counts[12] += jokers
	fmt.Printf("%v\n\n", counts)
	switch counts[12] {
	case 5:
		return FIVE_OF_A_KIND
	case 4:
		return FOUR_OF_A_KIND
	case 3:
		if counts[11] == 2 {
			return FULL_HOUSE
		}
		return THREE_OF_A_KIND
	case 2:
		if counts[11] == 2 {
			return TWO_PAIR
		}
		return ONE_PAIR
	}
	return HIGH_CARD
}

func cmpHands2(h1, h2 *Hand) int {
	if h1.Type2()-h2.Type2() == 0 {
		return h1.Cards - h2.Cards
	}
	return h1.Type2() - h2.Type2()
}

func readHands2(input []byte) []*Hand {
	hands := []*Hand{}
	start := 0
	var cards, bid int
	for end := range input {
		if input[end] == ' ' {
			cards = readCards2(input[start:end])
			start = end + 1
		}
		if input[end] == '\n' {
			bid = readNum(input[start:end])
			hands = append(hands, &Hand{Cards: cards, Bid: bid})
			start = end + 1
		}
	}
	return hands
}

func readCards2(input []byte) int {
	cards := 0
	for i, c := range input {
		switch c {
		case 'J':
			c = 0
		case '2', '3', '4', '5', '6', '7', '8', '9':
			c -= '1'
		case 'T':
			c = 9
		case 'Q':
			c = 10
		case 'K':
			c = 11
		case 'A':
			c = 12
		}
		cards |= int(c) << ((4 - i) * 8)
	}
	return cards
}
