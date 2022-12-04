package day03

import (
	"flag"
	"fmt"
)

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
