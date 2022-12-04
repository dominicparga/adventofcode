package day03

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

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
