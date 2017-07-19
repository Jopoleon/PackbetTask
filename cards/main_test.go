package main

import (
	"log"
	"testing"
)

var AceOfSpades = Card{Name: "Ace", Value: 8, Suit: Suit{IsTrump: true}}

func TestGenerateCards(t *testing.T) {
	trump := "Hearts"
	PC := GenerateCards(trump)
	expectedLength := 38
	if len(PC.Cards) != expectedLength {
		t.Fatalf("Bad number of cards: expected %s, got %s", expectedLength, len(PC.Cards))
	}
	for i, card := range PC.Cards {
		if PC.Cards[i].Suit.Name == trump && !PC.Cards[i].Suit.IsTrump {
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
		if PC.Cards[i].Suit.Name == trump && !PC.Cards[i].Suit.IsTrump {
			t.Fatalf("Some cards of trump suit are not trumps! : %+v", card)
		}
	}
	var expcetedAce = false
	for i, card := range PC.Cards {
		if PC.Cards[i].Name == "Ace" && PC.Cards[i].Value == 9 && PC.Cards[i].Suit.IsTrump {
			log.Println("Good, we got the AceOfSpades \\o/ , ", card)
			expcetedAce = true
		}
	}
	if !expcetedAce {
		t.Fatalf("There is no AceOfSpade inside new deck!")
	}
}
