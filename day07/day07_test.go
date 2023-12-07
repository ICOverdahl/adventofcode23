package day07

import "testing"

const testInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

func TestReadCards(t *testing.T) {
	cases := []struct {
		Input    []byte
		Expected int
	}{
		{[]byte("23456"), 0x0001020304},
		{[]byte("TJQKA"), 0x08090A0B0C},
		{[]byte("AAAAA"), 0x0C0C0C0C0C},
	}
	for _, c := range cases {
		actual := readCards(c.Input)
		if actual != c.Expected {
			t.Fail()
			t.Logf("%s: expected %x, got %x", c.Input, c.Expected, actual)
		}
	}
}

func TestReadHands(t *testing.T) {
	expected := []Hand{
		{0x010008010B, 765},
		{0x0803030903, 684},
		{0x0B0B040505, 28},
		{0x0B08090908, 220},
		{0x0A0A0A090C, 483},
	}
	actual := readHands([]byte(testInput))
	if len(actual) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}
	for i := range expected {
		if *actual[i] != expected[i] {
			t.Fatalf("Expected %v, got %v", expected, actual)
		}
	}
}

func TestHandType(t *testing.T) {
	cases := []struct {
		Cards    string
		Expected int
	}{
		{"23456", HIGH_CARD},
		{"TJQKA", HIGH_CARD},
		{"22345", ONE_PAIR},
		{"A2234", ONE_PAIR},
		{"KA223", ONE_PAIR},
		{"QKA22", ONE_PAIR},
		{"22334", TWO_PAIR},
		{"A2233", TWO_PAIR},
		{"22234", THREE_OF_A_KIND},
		{"A2223", THREE_OF_A_KIND},
		{"KA222", THREE_OF_A_KIND},
		{"22233", FULL_HOUSE},
		{"A222A", FULL_HOUSE},
		{"AA222", FULL_HOUSE},
		{"22223", FOUR_OF_A_KIND},
		{"32222", FOUR_OF_A_KIND},
		{"23222", FOUR_OF_A_KIND},
		{"22422", FOUR_OF_A_KIND},
		{"222A2", FOUR_OF_A_KIND},
		{"22222", FIVE_OF_A_KIND},
	}
	for _, c := range cases {
		hand := &Hand{Cards: readCards([]byte(c.Cards))}
		actual := hand.Type()
		if actual != c.Expected {
			t.Fail()
			t.Logf("%v: expected %x, got %x", *hand, c.Expected, actual)
		}
	}
}

func TestHandType2(t *testing.T) {
	cases := []struct {
		Cards    string
		Expected int
	}{
		{"2J345", ONE_PAIR},
		{"2J245", THREE_OF_A_KIND},
		{"2J244", FULL_HOUSE},
		{"JJ234", THREE_OF_A_KIND},
		{"JJJ34", FOUR_OF_A_KIND},
		{"JJJ33", FIVE_OF_A_KIND},
		{"JJJJ4", FIVE_OF_A_KIND},
		{"JJJJJ", FIVE_OF_A_KIND},
	}
	for _, c := range cases {
		hand := &Hand{Cards: readCards2([]byte(c.Cards))}
		actual := hand.Type2()
		if actual != c.Expected {
			t.Fail()
			t.Logf("%v: expected %x, got %x", *hand, c.Expected, actual)
		}
	}
}

func TestPart1(t *testing.T) {
	expected := 6440
	actual := Part1([]byte(testInput))
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 5905
	actual := Part2([]byte(testInput))
	if actual != expected {
		t.Fatalf("Expected %d, got %d", expected, actual)
	}
}
