package day06

import (
	"flag"
)

type Config struct {
	streamFilepath string
	patternLen     int
}

func parseCmdline(args []string) Config {
	flagSet := flag.NewFlagSet("Tuning Trouble", flag.ExitOnError)
	var helpMsg string

	helpMsg = "Message stream"
	streamFilepath := flagSet.String("stream-file", "", helpMsg)

	helpMsg = "Pattern length"
	patternLen := flagSet.Int("pattern-length", 4, helpMsg)

	flagSet.Parse(args)

	return Config{
		streamFilepath: *streamFilepath,
		patternLen:     *patternLen,
	}
}
