package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sort"
)

type Card struct {
	Name  string
	Value int
	Suit  Suit
}
type Suit struct {
	Name    string
	IsTrump bool
}
type PlayingCards struct {
	Cards []Card
}

func main() {
	var trump string
	flag.StringVar(&trump, "trump", "Spades", "Specify what suit do you want to be the trump (Hearts, Diamonds, Clubs or Spades)")
	flag.Parse()
	if trump != "Hearts" || trump != "Diamonds" || trump != "Clubs" || trump != "Spades" {
		log.Fatal("No such trump: ", trump, ". Use one of these (Hearts, Diamonds, Clubs or Spades)")
	}

	PC := GenerateCards(trump)
	fmt.Println("Ampout of card:", len(PC.Cards))
	for _, c := range PC.Cards {
		fmt.Printf("before shuffle: %+v\n", c)
	}

	Shuffle(PC)
	for _, c := range PC.Cards {
		fmt.Printf("after shuffle: %+v\n", c)
	}
	sort.Sort(PC)
	for i, c := range PC.Cards {
		fmt.Printf("after SORT: %d %+v\n", i, c)
	}
}

//GenerateCards returns default deck of playing cards
func GenerateCards(trump string) PlayingCards {
	var pc []Card
	var CardNames = []string{"6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	var CardSuits = []*Suit{
		{Name: "Hearts"},
		{Name: "Diamonds"},
		{Name: "Clubs"},
		{Name: "Spades"},
	}
	//log.Printf("Suits: %+v", CardSuits)
	if trump != "" {
		CardSuits = withTrump(CardSuits, trump)
	}
	log.Printf("Suits: %+v", CardSuits)
	for j := 0; j < len(CardSuits); j++ {
		for i := 0; i < len(CardNames); i++ {
			pc = append(pc, Card{Name: CardNames[i], Value: i + 1, Suit: *CardSuits[j]})
		}
	}
	pc = append(pc, Card{Name: "JokerRed", Value: 8, Suit: Suit{IsTrump: true}})
	pc = append(pc, Card{Name: "JokerBlack", Value: 8, Suit: Suit{IsTrump: true}})
	return PlayingCards{Cards: pc}
}

// withTrump makes given suit trump
func withTrump(CardSuits []*Suit, trump string) []*Suit {
	for _, c := range CardSuits {
		if c.Name == trump {
			log.Println("True", c)
			c.IsTrump = true
		}
	}
	return CardSuits
}

//Shuffle shuffles cards in random order
func Shuffle(data Interface) {
	for i := data.Len() - 1; i > 0; i-- {
		if j := rand.Intn(i + 1); i != j {
			data.Swap(i, j)
		}
	}
}
func (c PlayingCards) Len() int {
	return len(c.Cards)
}

func (c PlayingCards) Less(i, j int) bool {

	//check for Jokers
	if c.Cards[i].Name == "JokerRed" || c.Cards[i].Name == "JokerBlack" {
		return false
	}
	if c.Cards[j].Name == "JokerRed" || c.Cards[j].Name == "JokerBlack" {
		return true
	}
	//if card is the trump
	if c.Cards[i].Suit.IsTrump && c.Cards[j].Suit.IsTrump {
		return c.Cards[i].Value < c.Cards[j].Value

	}
	if c.Cards[i].Suit.IsTrump && !c.Cards[j].Suit.IsTrump {
		return false

	}
	if !c.Cards[i].Suit.IsTrump && c.Cards[j].Suit.IsTrump {
		return true

	}
	//if card is not the trump
	if !c.Cards[i].Suit.IsTrump && !c.Cards[j].Suit.IsTrump {
		return c.Cards[i].Value < c.Cards[j].Value

	}
	return c.Cards[i].Value < c.Cards[j].Value

}

func (c PlayingCards) Swap(i, j int) {
	c.Cards[i], c.Cards[j] = c.Cards[j], c.Cards[i]
}

//interface fo shuffling
type Interface interface {
	Len() int
	Swap(i, j int)
}
