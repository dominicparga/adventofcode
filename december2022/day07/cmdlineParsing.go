package day07

import (
	"flag"
)

type Config struct {
	logFilepath string
}

func parseCmdline(args []string) (Config, error) {
	flagSet := flag.NewFlagSet("no space left on device", flag.ExitOnError)
	var helpMsg string

	helpMsg = "Message stream"
	logFilepath := flagSet.String("log-file", "", helpMsg)

	if err := flagSet.Parse(args); err != nil {
		return Config{}, err
	}

	return Config{
		logFilepath: *logFilepath,
	}, nil
}
