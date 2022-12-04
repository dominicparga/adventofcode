package day03

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

var priorityByItemType = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func findCommonItemTypes(packedItems string, prioritySumChannel chan<- int) {
	k := len(packedItems) / 2
	firstCompartment := packedItems[:k]
	secondCompartment := packedItems[k:]
	// fmt.Println("packed items:", packedItems)
	// fmt.Println("1st compartment:", firstCompartment)
	// fmt.Println("2nd compartment:", secondCompartment)

	priorityListFromCompartment := func(compartment string) []int {
		itemTypeList := strings.Split(compartment, "")

		sort.Slice(itemTypeList, func(i int, j int) bool {
			// ascending order is expected below
			return priorityByItemType[itemTypeList[i]] < priorityByItemType[itemTypeList[j]]
		})
		var priorityList []int
		for _, itemType := range itemTypeList {
			itemPriority := priorityByItemType[itemType]
			// ignore duplicates since we would like to sum up over types, not items
			if len(priorityList) == 0 || priorityList[len(priorityList)-1] != itemPriority {
				priorityList = append(priorityList, itemPriority)
			}
		}
		return priorityList
	}
	firstPriorityList := priorityListFromCompartment(firstCompartment)
	secondPriorityList := priorityListFromCompartment(secondCompartment)

	// fmt.Println("1st priolist:", firstPriorityList)
	// fmt.Println("2nd priolist:", secondPriorityList)

	firstI := 0
	secondI := 0
	prioritySum := 0
	for firstI < len(firstPriorityList) && secondI < len(secondPriorityList) {
		firstPriority := firstPriorityList[firstI]
		secondPriority := secondPriorityList[secondI]
		// for this to work, a sorting in ascending order is required
		if firstPriority < secondPriority {
			firstI++
		} else if secondPriority < firstPriority {
			secondI++
		} else {
			prioritySum += firstPriority
			firstI++
			secondI++
		}
	}

	// log.Fatalln("priosum", prioritySum)

	prioritySumChannel <- prioritySum
}

type Config struct {
	packedRucksacksFilepath string
}

func parseCmdline(args []string) Config {
	flagSet := flag.NewFlagSet("Rucksack Reorganization", flag.ExitOnError)
	var helpMsg string

	helpMsg = "The Elves' rucksacks' packing"
	packedRucksacksFilepath := flagSet.String("packed-rucksacks-file", "", helpMsg)

	flagSet.Parse(args)

	return Config{
		packedRucksacksFilepath: *packedRucksacksFilepath,
	}
}

func Run(args []string) error {
	config := parseCmdline(args)

	var err error
	packedRucksacks, err := os.ReadFile(config.packedRucksacksFilepath)
	if err != nil {
		return err
	}

	waitGroup := new(sync.WaitGroup)
	prioritySumChannel := make(chan int)
	packedItemsList := strings.Split(string(packedRucksacks), "\n")
	for _, packedItems := range packedItemsList {
		waitGroup.Add(1)
		go findCommonItemTypes(packedItems, prioritySumChannel)
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

	return nil
}
