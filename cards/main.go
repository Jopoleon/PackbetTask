package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sort"
)

type Card struct {
	Name    string
	Value   int
	Suit    string
	IsTrump bool
}
type PlayingCards struct {
	Cards []Card
}

type lessFunc func(p1, p2 *Card) bool

// multiSorter implements the Sort interface, sorting the changes within.
type multiSorter struct {
	changes []Card
	less    []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(changes *[]Card) {
	ms.changes = *changes
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.changes)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that is either Less or
// !Less. Note that it can call the less functions twice per call. We
// could change the functions to return -1, 0, 1 and reduce the
// number of calls for greater efficiency: an exercise for the reader.
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

func GenerateCards(trump string) PlayingCards {
	var pc []Card
	var CardNames = []string{"6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	var CardSuits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	//log.Printf("Suits: %+v", CardSuits)

	//log.Printf("Suits: %+v", CardSuits)
	for j := 0; j < len(CardSuits); j++ {
		for i := 0; i < len(CardNames); i++ {
			if CardSuits[j] == trump {
				pc = append(pc, Card{Name: CardNames[i], Value: i + 1, Suit: CardSuits[j], IsTrump: true})
			} else {
				pc = append(pc, Card{Name: CardNames[i], Value: i + 1, Suit: CardSuits[j]})
			}

		}
	}
	pc = append(pc, Card{Name: "JokerRed", Value: len(CardNames) + 1, IsTrump: true})
	pc = append(pc, Card{Name: "JokerBlack", Value: len(CardNames) + 1, IsTrump: true})
	return PlayingCards{Cards: pc}
}

func Shuffle1(slc []Card) {
	for i := 1; i < len(slc); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slc[r], slc[i] = slc[i], slc[r]
		}
	}
}

// ExampleMultiKeys demonstrates a technique for sorting a struct type using different
// sets of multiple fields in the comparison. We chain together "Less" functions, each of
// which compares a single field.
func main() {
	var trump string
	flag.StringVar(&trump, "trump", "", "Specify what suit do you want to be the trump (Hearts, Diamonds, Clubs or Spades)")
	flag.Parse()
	if trump != "Hearts" && trump != "Diamonds" && trump != "Clubs" && trump != "Spades" && trump != "" {
		log.Fatal("No such trump: ", trump, ". Use one of these (Hearts, Diamonds, Clubs or Spades)")
	}
	cards2 := GenerateCards(trump)
	cards1 := cards2.Cards
	Shuffle1(cards1)
	// Closures that order the PlayingCards structure.
	byName := func(c1, c2 *Card) bool {
		return c1.Name < c2.Name
	}
	byValue := func(c1, c2 *Card) bool {
		return c1.Value < c2.Value
	}
	byTrump := func(c1, c2 *Card) bool {

		if c1.IsTrump && !c2.IsTrump {
			return false

		}
		if !c1.IsTrump && c2.IsTrump {
			return true

		}
		return false
	}

	OrderedBy(byName).Sort(&cards1)
	fmt.Printf("By byName: %+v\n", cards1)

	Shuffle1(cards1)
	OrderedBy(byValue).Sort(&cards1)
	fmt.Printf("By byValue: %+v\n", cards1)

	Shuffle1(cards1)
	OrderedBy(byTrump).Sort(&cards1)
	fmt.Printf("By byTrump: %+v\n", cards1)

	Shuffle1(cards1)
	OrderedBy(byTrump, byValue).Sort(&cards1)
	fmt.Printf("By byTrump>byValue: %+v\n", cards1)
}
