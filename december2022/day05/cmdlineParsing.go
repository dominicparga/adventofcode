package day05

import (
	"flag"
)

type Config struct {
	stackAndMovesFilepath string
}

func parseCmdline(args []string) Config {
	flagSet := flag.NewFlagSet("Camp Cleanup", flag.ExitOnError)
	var helpMsg string

	helpMsg = "File containing stack and stack moves"
	stackAndMovesFilepath := flagSet.String("stack-and-moves-file", "", helpMsg)

	flagSet.Parse(args)

	return Config{
		stackAndMovesFilepath: *stackAndMovesFilepath,
	}
}
