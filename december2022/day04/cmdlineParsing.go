package day04

import (
	"flag"
)

type Config struct {
	assignmentPairsFilepath string
	strategy                Strategy
}

func parseCmdline(args []string) Config {
	flagSet := flag.NewFlagSet("Camp Cleanup", flag.ExitOnError)
	var helpMsg string

	helpMsg = "Assignments in pairs to cleanup camp section"
	assignmentPairsFilepath := flagSet.String("assignment-pairs-file", "", helpMsg)

	helpMsg = "Partial overlap vs. full overlap"
	strategy := flagSet.String("overlap", "full", helpMsg)

	flagSet.Parse(args)

	return Config{
		assignmentPairsFilepath: *assignmentPairsFilepath,
		strategy:                *strategy,
	}
}
