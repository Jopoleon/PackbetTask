package main

import (
	"log"
	"sort"
	"testing"
)

var AceOfSpades = Card{Name: "Ace", Value: 8, Suit: "Spades", IsTrump: true}

func TestGenerateCards(t *testing.T) {
	trump := "Hearts"
	PC := GenerateCards(trump)
	expectedLength := 38
	if len(PC.Cards) != expectedLength {
		t.Fatalf("Bad number of cards: expected %s, got %s", expectedLength, len(PC.Cards))
	}
	for i, card := range PC.Cards {
		if PC.Cards[i].Suit == trump && !PC.Cards[i].IsTrump {
			t.Fatalf("Some cards of trump suit are not trumps! : %+v", card)
		}
	}

	trump = "Spades"
	PC = GenerateCards(trump)
	//log.Println()
	expectedLength = 38
	if len(PC.Cards) != expectedLength {
		t.Fatalf("Bad number of cards: expected %s, got %s", expectedLength, len(PC.Cards))
	}
	for i, card := range PC.Cards {
		if PC.Cards[i].Name == trump && !PC.Cards[i].IsTrump {
			t.Fatalf("Some cards of trump suit are not trumps! : %+v", card)
		}
	}
	var expcetedAce = false
	for i, card := range PC.Cards {
		if PC.Cards[i].Name == "Ace" && PC.Cards[i].Value == 9 && PC.Cards[i].IsTrump {
			log.Println("Good, we got the AceOfSpades \\o/ , ", card)
			expcetedAce = true
		}
	}
	if !expcetedAce {
		t.Fatalf("There is no AceOfSpade inside new deck!")
	}
}
func TestShuffle1(t *testing.T) {
	trump := "Spades"
	PC := GenerateCards(trump)
	cards := PC.Cards
	Shuffle1(cards)

	byValue := func(c1, c2 *Card) bool {
		return c1.Value < c2.Value
	}
	//OrderedBy(byValue).Sort(&cards)
	ms := multiSorter{changes: cards, less: OrderedBy(byValue).less}
	if sort.IsSorted(&ms) {
		t.Fatalf("cards are sorted but they do not must be s:\n %+v", cards)
	}
}
func TestOrderedBy(t *testing.T) {
	trump := "Spades"
	PC := GenerateCards(trump)
	cards := PC.Cards
	Shuffle1(cards)
	//byName := func(c1, c2 *Card) bool {
	//	return c1.Name < c2.Name
	//}
	byValue := func(c1, c2 *Card) bool {
		return c1.Value < c2.Value
	}
	byTrump := func(c1, c2 *Card) bool {
		//if c1.IsTrump && c2.IsTrump {
		//	return c1.Value < c2.Value
		//}
		if c1.IsTrump && !c2.IsTrump {
			return false

		}
		if !c1.IsTrump && c2.IsTrump {
			return true

		}
		return false
	}
	Shuffle1(cards)
	OrderedBy(byValue).Sort(&cards)
	ms := multiSorter{changes: cards, less: OrderedBy(byValue).less}
	if !sort.IsSorted(&ms) {
		t.Fatalf("OrderedBy(byValue) is not sorting cards:\n %+v", cards)
	}
	Shuffle1(cards)
	OrderedBy(byTrump).Sort(&cards)
	ms2 := multiSorter{changes: cards, less: OrderedBy(byTrump).less}
	if !sort.IsSorted(&ms2) {
		t.Fatalf("OrderedBy(byTrump) is not sorting cards:\n %+v", cards)
	}
	Shuffle1(cards)
	OrderedBy(byTrump, byValue).Sort(&cards)
	ms3 := multiSorter{changes: cards, less: OrderedBy(byTrump, byValue).less}
	if !sort.IsSorted(&ms3) {
		t.Fatalf("OrderedBy(byTrump,byValue) is not sorting cards: \n %+v", cards)
	}
}
