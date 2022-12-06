package day04

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func checkAssignmentsForOverlap(rawAssignmentPair RawAssignmentPair, strategy Strategy, resultChannel chan<- bool) {
	var assignmentIntervalPair AssignmentIntervalPair = assignmentIntervalPairFromRawAssignmentPair(rawAssignmentPair)
	switch strategy {
	case "partial":
		resultChannel <- assignmentIntervalPair.checkForPartialOverlap()
	case "full":
		resultChannel <- assignmentIntervalPair.checkForFullOverlap()
	default:
		log.Fatalln("Unknown strategy", strategy)
	}
}

func Run(args []string) error {
	config := parseCmdline(args)
	var err error
	assignmentPairByteList, err := os.ReadFile(config.assignmentPairsFilepath)
	if err != nil {
		return err
	}

	assignmentPairList := strings.Split(string(assignmentPairByteList), "\n")

	waitGroup := new(sync.WaitGroup)
	resultChannel := make(chan bool)
	for _, assignmentPair := range assignmentPairList {
		waitGroup.Add(1)
		go checkAssignmentsForOverlap(assignmentPair, config.strategy, resultChannel)
	}

	overlapCount := 0
	go func() {
		for {
			if <-resultChannel {
				overlapCount++
			}
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
	fmt.Println("Number of completely overlapping assignment pairs:", overlapCount)

	return err
}
