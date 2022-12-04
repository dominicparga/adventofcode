package day01

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sumUpCalorieGroup(calorieGroup string, calorieSumChannel chan<- int) {
	calorieSum := 0
	for _, calorieStr := range strings.Fields(calorieGroup) {
		calorie, _ := strconv.Atoi(calorieStr)
		calorieSum += calorie
	}
	calorieSumChannel <- calorieSum
}

type Config struct {
	inputFilepath string
	kSum          uint
}

func parseCmdline(args []string) Config {
	flagSet := flag.NewFlagSet("Calorie Counting", flag.ExitOnError)
	var helpMsg string

	helpMsg = "The Elves' Calorie list (for Santa's reindeers)"
	inputFilepath := flagSet.String("calorie-group-list-file", "", helpMsg)

	helpMsg = "Sum up the Calories of this number of Elves' with maximum Calorie carriage"
	kSum := flagSet.Uint("sum-k", 1, helpMsg)

	flagSet.Parse(args)

	return Config{
		inputFilepath: *inputFilepath,
		kSum:          *kSum,
	}
}

func Run(args []string) error {
	config := parseCmdline(args)

	var err error
	content, err := os.ReadFile(config.inputFilepath)
	if err != nil {
		return err
	}

	calorieGroupList := strings.Split(string(content), "\n\n")
	waitGroup := new(sync.WaitGroup)
	calorieSumChannel := make(chan int)
	for _, calorieGroup := range calorieGroupList {
		waitGroup.Add(1)
		go sumUpCalorieGroup(calorieGroup, calorieSumChannel)
	}
	calorieSumList := []int{}
	go func() {
		for {
			calorieSumList = append(calorieSumList, <-calorieSumChannel)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
	sort.Slice(calorieSumList, func(i int, j int) bool {
		return calorieSumList[i] > calorieSumList[j]
	})
	maxCalorieSum := 0
	for i := 0; i < int(config.kSum); i++ {
		maxCalorieSum += calorieSumList[i]
	}
	fmt.Println("Sum of", config.kSum, "maximum sums of calories:", maxCalorieSum)

	return nil
}
