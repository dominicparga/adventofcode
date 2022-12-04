package day01

import (
	"adventofcode/december2022/day01/calorieCounting"
	"errors"
	"fmt"
	"log"
)

func usage() string {
	return fmt.Sprintln(
		"Usage:\n",
		"    adventofcode 2022 1 PUZZLE\n",
		"\n",
		"    PUZZLE\n",
		"    Choose from [calorie-counting].",
	)
}

type Config struct {
	puzzle string
	args   []string
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
		puzzle: args[0],
		args:   args[1:],
	}
}

func Run(args []string) error {
	config := parseCmdline(args)

	var err error
	switch config.puzzle {
	case "calorie-counting":
		err = calorieCounting.Run(config.args)
	default:
		err = errors.New(
			fmt.Sprintln("[ERROR] Puzzle", config.puzzle, "is not supported.\n",
				"\n",
				usage(),
			))
	}

	return err
}
