package main

import (
	"adventofcode/december2022"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func usage() string {
	return fmt.Sprintln(
		"Usage:\n",
		"    adventofcode YEAR ARGS\n",
		"\n",
		"    YEAR\n",
		"    Choose between [2022].",
	)
}

type Config struct {
	year string
	day  string
	args []string
}

func parseCmdline(args []string) Config {
	const expectedLength int = 1
	if len(args) < expectedLength {
		log.Fatalln(
			"[ERROR] Too few arguments\n",
			"\n",
			usage(),
		)
	}

	return Config{
		year: args[0],
		args: args[expectedLength:],
	}
}

func createStopwatch() func() time.Duration {
	startMoment := time.Now()
	return func() time.Duration {
		return time.Since(startMoment)
	}
}

func main() {
	defer fmt.Println("Time elapsed: ", createStopwatch()())

	config := parseCmdline(os.Args[1:])

	var err error
	switch config.year {
	case "2022":
		err = december2022.Run(config.args)
	default:
		err = errors.New(fmt.Sprintln(
			"[ERROR] Year", config.year, "is not supported.\n",
			"\n",
			usage(),
		))
	}

	if err != nil {
		log.Fatalln(err)
	}
}
