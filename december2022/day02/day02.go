package day02

import (
	"flag"
	"fmt"
	"os"
)

var encryptionTable = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
}

type Config struct {
	inputFilepath string
}

func parseCmdline(args []string) Config {
	flagSet := flag.NewFlagSet("rock-paper-scissors", flag.ExitOnError)
	var helpMsg string

	helpMsg = "The Elves' encrypted strategy guide for Rock Paper Scissors"
	inputFilepath := flagSet.String("encryptedStrategyGuide.txt", "", helpMsg)

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

	fmt.Println(content)

	return nil
}
