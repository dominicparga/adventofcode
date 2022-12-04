package day03

import (
	"sort"
	"strings"
)

// NOTE
// Rucksackand Compartment are defined logically, not intuitively.
// Hence during these calculations, Rucksacks can be interpreted as one or multiple Compartments.
// Intuitively, a Rucksack is probably seen as []Compartment, which is not the case here.
type Item = string
type ItemType = Item
type Compartment = []Item
type Rucksack = string
type Priority = int

var priorityByItemType = map[ItemType]Priority{
	"a": 1, "b": 2, "c": 3, "d": 4, "e": 5,
	"f": 6, "g": 7, "h": 8, "i": 9, "j": 10,
	"k": 11, "l": 12, "m": 13, "n": 14, "o": 15,
	"p": 16, "q": 17, "r": 18, "s": 19, "t": 20,
	"u": 21, "v": 22, "w": 23, "x": 24, "y": 25,
	"z": 26, "A": 27, "B": 28, "C": 29, "D": 30,
	"E": 31, "F": 32, "G": 33, "H": 34, "I": 35,
	"J": 36, "K": 37, "L": 38, "M": 39, "N": 40,
	"O": 41, "P": 42, "Q": 43, "R": 44, "S": 45,
	"T": 46, "U": 47, "V": 48, "W": 49, "X": 50,
	"Y": 51, "Z": 52,
}

func compartmentFromRucksack(s string) Compartment {
	return strings.Split(s, "")
}

func priorityListFromCompartment(compartment Compartment) []Priority {
	sort.Slice(compartment, func(i int, j int) bool {
		// ascending order is expected below
		return priorityByItemType[compartment[i]] < priorityByItemType[compartment[j]]
	})
	var priorityList []Priority
	for _, item := range compartment {
		var itemType ItemType = item
		var itemPriority Priority = priorityByItemType[itemType]
		// ignore duplicates since we would like to sum up over types, not items
		if len(priorityList) == 0 || priorityList[len(priorityList)-1] != itemPriority {
			priorityList = append(priorityList, itemPriority)
		}
	}
	return priorityList
}
