package day05

import (
	"flag"
)

type Config struct {
	stackAndMovesFilepath string
	multimoveStrategy     string
}

func parseCmdline(args []string) Config {
	flagSet := flag.NewFlagSet("Supply Stacks", flag.ExitOnError)
	var helpMsg string

	helpMsg = "File containing stack and stack moves"
	stackAndMovesFilepath := flagSet.String("stack-and-moves-file", "", helpMsg)

	helpMsg = "Strategy (LIFO vs FIFO) when moving multiple crates"
	multimoveStrategy := flagSet.String("multimove-strategy", "lifo", helpMsg)

	flagSet.Parse(args)

	return Config{
		stackAndMovesFilepath: *stackAndMovesFilepath,
		multimoveStrategy:     *multimoveStrategy,
	}
}
