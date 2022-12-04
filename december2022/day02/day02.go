package day02

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

const rockVs = "A"
const paperVs = "B"
const scissorsVs = "C"

// default strategy
const defaultStrategyName = "default"
const rockMe = "X"
const paperMe = "Y"
const scissorsMe = "Z"

// ultra top secret strategy
const ultraTopSecretStrategyName = "ultraTopSecret"
const meLosing = "X"
const meDrawing = "Y"
const meWinning = "Z"

const lossScore = 0
const drawScore = 3
const winScore = 6
const rockMeScore = 1
const paperMeScore = 2
const scissorsMeScore = 3

var scoreTable = map[string]map[string]int{
	defaultStrategyName: {
		rockVs + " " + rockMe:         rockMeScore + drawScore,
		rockVs + " " + paperMe:        paperMeScore + winScore,
		rockVs + " " + scissorsMe:     scissorsMeScore + lossScore,
		paperVs + " " + rockMe:        rockMeScore + lossScore,
		paperVs + " " + paperMe:       paperMeScore + drawScore,
		paperVs + " " + scissorsMe:    scissorsMeScore + winScore,
		scissorsVs + " " + rockMe:     rockMeScore + winScore,
		scissorsVs + " " + paperMe:    paperMeScore + lossScore,
		scissorsVs + " " + scissorsMe: scissorsMeScore + drawScore,
	},
	ultraTopSecretStrategyName: {
		rockVs + " " + meLosing:      scissorsMeScore + lossScore,
		rockVs + " " + meDrawing:     rockMeScore + drawScore,
		rockVs + " " + meWinning:     paperMeScore + winScore,
		paperVs + " " + meLosing:     rockMeScore + lossScore,
		paperVs + " " + meDrawing:    paperMeScore + drawScore,
		paperVs + " " + meWinning:    scissorsMeScore + winScore,
		scissorsVs + " " + meLosing:  paperMeScore + lossScore,
		scissorsVs + " " + meDrawing: scissorsMeScore + drawScore,
		scissorsVs + " " + meWinning: rockMeScore + winScore,
	},
}

func computeRoundScore(round string, strategy string, roundScoreChannel chan<- int) {
	choiceList := strings.Fields(round)
	roundScoreChannel <- scoreTable[strategy][choiceList[0]+" "+choiceList[1]]
}

type Config struct {
	inputFilepath string
	strategy      string
}

func parseCmdline(args []string) Config {
	flagSet := flag.NewFlagSet("Rock Paper Scissors", flag.ExitOnError)
	var helpMsg string

	helpMsg = "The Elves' encrypted strategy guide for Rock Paper Scissors"
	inputFilepath := flagSet.String("calorie-group-list-file", "", helpMsg)

	helpMsg = "Default strategy or ultra top secret strategy"
	isUsingUltraTopSecretStrategy := flagSet.Bool("ultra-top-secret-strategy", false, helpMsg)

	flagSet.Parse(args)

	strategy := defaultStrategyName
	if *isUsingUltraTopSecretStrategy {
		strategy = ultraTopSecretStrategyName
	}
	return Config{
		inputFilepath: *inputFilepath,
		strategy:      strategy,
	}
}

func Run(args []string) error {
	config := parseCmdline(args)

	var err error
	strategyGuide, err := os.ReadFile(config.inputFilepath)
	if err != nil {
		return err
	}

	waitGroup := new(sync.WaitGroup)
	roundScoreChannel := make(chan int)
	roundList := strings.Split(string(strategyGuide), "\n")
	for _, round := range roundList {
		waitGroup.Add(1)
		go computeRoundScore(round, config.strategy, roundScoreChannel)
	}
	totalScore := 0
	go func() {
		for {
			totalScore += <-roundScoreChannel
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
	fmt.Println("Total score:", totalScore)

	return nil
}
