package calorieCounting

import (
	"flag"
	"fmt"
	"os"
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

func selectMax(maxCalorieSum *int, calorieSumChannel <-chan int) {
	calorieSum := <-calorieSumChannel
	if *maxCalorieSum < calorieSum {
		*maxCalorieSum = calorieSum
	}
}

type Config struct {
	inputFilepath string
}

func parseCmdline(args []string) Config {
	flagSet := flag.NewFlagSet("calorie-counting", flag.ExitOnError)

	help_msg := "The Elves' calorie list (for Santa's reindeers)"
	inputFilepath := flagSet.String("input-file", "", help_msg)

	flagSet.Parse(args)

	return Config{
		inputFilepath: *inputFilepath,
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
	maxCalorieSum := 0
	go func() {
		for {
			selectMax(&maxCalorieSum, calorieSumChannel)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
	fmt.Println("Maximum sum of calories:", maxCalorieSum)

	return nil
}
