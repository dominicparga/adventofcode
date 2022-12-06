package december2022

import (
	"adventofcode/december2022/day01"
	"adventofcode/december2022/day02"
	"adventofcode/december2022/day03"
	"adventofcode/december2022/day04"
	"errors"
	"fmt"
	"log"
)

func usage() string {
	return fmt.Sprintln(
		"Usage:\n",
		"    adventofcode 2022 DAY ARGS\n",
		"\n",
		"    DAY\n",
		"    Choose from [1,...,24].",
	)
}

type Config struct {
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
		day:  args[0],
		args: args[1:],
	}
}

func Run(args []string) error {
	config := parseCmdline(args)

	var err error
	switch config.day {
	case "1", "01", "calorie-counting":
		err = day01.Run(config.args)
	case "2", "02", "rock-paper-scissors":
		err = day02.Run(config.args)
	case "3", "03", "rucksack-reorganization":
		err = day03.Run(config.args)
	case "4", "04", "camp-cleanup":
		err = day04.Run(config.args)
	default:
		err = errors.New(
			fmt.Sprintln("[ERROR] Day", config.day, "is not supported.\n",
				"\n",
				usage(),
			))
	}

	return err
}
