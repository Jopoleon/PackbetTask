package main

import (
	"testing"
)

var OKpositionTests = []struct {
	in       string
	expected []string
}{
	{"a8", []string{"c7", "b6"}},
	{"h8", []string{"f7", "g6"}},
	{"h1", []string{"f2", "g3"}},
	{"a1", []string{"c2", "b3"}},
	{"b7", []string{"d8", "d6", "c5", "a5"}},
	{"g2", []string{"e1", "e3", "f4", "h4"}},
	{"b2", []string{"a4", "c4", "d3", "d1"}},
	{"b2", []string{"a4", "c4", "d3", "d1"}},
	{"g7", []string{"e6", "e8", "h5", "f5"}},
	{"e7", []string{"c6", "c8", "g8", "g6", "f5", "d5"}},
	{"d7", []string{"b8", "b6", "c5", "e5", "f6", "f8"}},
	{"g4", []string{"h6", "f6", "e5", "e3", "f2", "h2"}},
	{"d2", []string{"b1", "b3", "c4", "e4", "f3", "f1"}},
	{"b5", []string{"a7", "c7", "d6", "d4", "a3", "c3"}},
	{"d4", []string{"b5", "b3", "c6", "e6", "f5", "f3", "c2", "e2"}},
	{"c5", []string{"a4", "a6", "b7", "d7", "e6", "e4", "d3", "b3"}},
}

var ErrpositionTests = []struct {
	in       string
	expected []string
}{
	{"d3", []string{"c7", "b6"}},
	{"a8", []string{"f7", "g6"}},
	{"c1", []string{"f2", "g3"}},
	{"b1", []string{"c2", "b3"}},
	{"c7", []string{"d8", "d6", "c5", "a5"}},
	{"d2", []string{"e1", "e3", "f4", "h4"}},
	{"f2", []string{"a4", "c4", "d3", "d1"}},
	{"g2", []string{"a4", "c4", "d3", "d1"}},
	{"e7", []string{"e6", "e8", "h5", "f5"}},
	{"a7", []string{"c6", "c8", "g8", "g6", "f5", "d5"}},
	{"e7", []string{"b8", "b6", "c5", "e5", "f6", "f8"}},
	{"e4", []string{"h6", "f6", "e5", "e3", "f2", "h2"}},
	{"e2", []string{"b1", "b3", "c4", "e4", "f3", "f1"}},
	{"e5", []string{"a7", "c7", "d6", "d4", "a3", "c3"}},
	{"e4", []string{"b5", "b3", "c6", "e6", "f5", "f3", "c2", "e2"}},
	{"e5", []string{"a4", "a6", "b7", "d7", "e6", "e4", "d3", "b3"}},
}

func TestKnightMoves(t *testing.T) {
	for _, pt := range OKpositionTests {
		result := knightMoves(pt.in)
		if !CompareSlices(result, pt.expected) {
			t.Errorf("Wrong positions for knight, expected %v, got %v", pt.expected, result)
		}
	}
	for _, pt := range ErrpositionTests {
		result := knightMoves(pt.in)
		if CompareSlices(result, pt.expected) {
			t.Errorf("Error values passed, expected %v, got %v", pt.expected, result)
		}
	}
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func compareElements(a, b []string) bool {
	for _, elm1 := range a {
		if contains(b, elm1) {
			for _, elm2 := range b {
				if contains(a, elm2) {
					return true
				}
				return false
			}
		}
		return false
	}
	return false
}
func CompareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	return compareElements(a, b)
}
