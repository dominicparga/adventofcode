package day03

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
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

func createRucksackListList(rawRucksackList string, splitFactor int) ([][]Rucksack, error) {
	var originalRucksackList []Rucksack = strings.Split(rawRucksackList, "\n")
	var rucksackListList [][]Rucksack

	// rearrange packed items
	switch splitFactor {
	case -2:
		for _, rucksack := range originalRucksackList {
			k := len(rucksack) / -splitFactor
			rucksackListList = append(rucksackListList, []Rucksack{rucksack[:k], rucksack[k:]})
		}
	case 3:
		i := 0
		for ; i < len(originalRucksackList); i += 3 {
			rucksackListList = append(rucksackListList, originalRucksackList[i:i+3])
		}
		if i < len(originalRucksackList) {
			rucksackListList = append(rucksackListList, originalRucksackList[i:])
		}
	default:
		return nil, errors.New(fmt.Sprint("Unsupported split factor ", splitFactor))
	}

	return rucksackListList, nil
}

func findCommonItemTypes(prioritySumChannel chan<- int, rucksackList ...Rucksack) {
	// fmt.Println("packed items:", packedItemsList)

	var priorityListList [][]Priority
	var priorityListIndexList []int
	for _, rucksack := range rucksackList {
		priorityListList = append(priorityListList, priorityListFromCompartment(compartmentFromRucksack(rucksack)))
		priorityListIndexList = append(priorityListIndexList, 0)
	}

	// fmt.Println("prioListList:   ", priorityListList)
	// fmt.Println("prioListIdxList:", priorityListIndexList)

	isEachIndexInRange := func() bool {
		for i, priorityListIndex := range priorityListIndexList {
			if priorityListIndex >= len(priorityListList[i]) {
				return false
			}
		}
		return true
	}
	equalizePointedToPriorities := func() bool {
		// comparing and adapting pairwise is enough, if
		// - equal pairs are jumped over
		// - unequal pairs are corrected and loop starts again
		for i := 0; i < len(priorityListList)-1; i++ {
			j := i + 1
			var priorityIndexI *int = &priorityListIndexList[i]
			var priorityIndexJ *int = &priorityListIndexList[j]
			var priorityI Priority = priorityListList[i][*priorityIndexI]
			var priorityJ Priority = priorityListList[j][*priorityIndexJ]
			if priorityI != priorityJ {
				// for this to work, a sorting in ascending order is required
				if priorityI < priorityJ {
					(*priorityIndexI)++
				} else if priorityJ < priorityI {
					(*priorityIndexJ)++
				}
				// priorities are not equal
				return false
			}
		}
		// priorities are all equal
		return true
	}
	incPriorityIndices := func() {
		for i := 0; i < len(priorityListList); i++ {
			priorityListIndexList[i]++
		}
	}

	prioritySum := 0
	for isEachIndexInRange() {
		var arePointedToPrioritiesEqual bool = equalizePointedToPriorities()
		if arePointedToPrioritiesEqual {
			// since all are equal and in range, any index can be chosen
			i := 0
			prioritySum += priorityListList[i][priorityListIndexList[i]]
			incPriorityIndices()
		}
	}

	// log.Fatalln("priosum", prioritySum)

	prioritySumChannel <- prioritySum
}

type Config struct {
	rucksacksFilepath string
	splitFactor       int
}

func parseCmdline(args []string) Config {
	flagSet := flag.NewFlagSet("Rucksack Reorganization", flag.ExitOnError)
	var helpMsg string

	helpMsg = "The Elves' rucksacks' packing"
	rucksacksFilepath := flagSet.String("rucksacks-file", "", helpMsg)

	helpMsg = fmt.Sprint(
		"Splitting strategy; for simplicity, only supported values are -2 and 3.",
		"If split factor is negative, it is interpreted as the number of compartments.",
		"If split factor is positive, it is interpreted as the number of rucksacks.",
	)
	splitFactor := flagSet.Int("split-factor", -2, helpMsg)

	flagSet.Parse(args)

	return Config{
		rucksacksFilepath: *rucksacksFilepath,
		splitFactor:       *splitFactor,
	}
}

func Run(args []string) error {
	config := parseCmdline(args)
	var err error
	rucksackByteList, err := os.ReadFile(config.rucksacksFilepath)
	if err != nil {
		return err
	}
	rucksackListList, err := createRucksackListList(string(rucksackByteList), config.splitFactor)
	if err != nil {
		return err
	}

	waitGroup := new(sync.WaitGroup)
	prioritySumChannel := make(chan int)
	for _, rucksackList := range rucksackListList {
		waitGroup.Add(1)
		go findCommonItemTypes(prioritySumChannel, rucksackList...)
	}

	prioritySum := 0
	go func() {
		for {
			prioritySum += <-prioritySumChannel
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
	fmt.Println("Total priority sum of mixed item types:", prioritySum)

	return err
}
