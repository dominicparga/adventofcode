package december2022

import (
	"adventofcode/december2022/day01"
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
	case "1":
		err = day01.Run(config.args)
	default:
		err = errors.New(
			fmt.Sprintln("[ERROR] Day", config.day, "is not supported.\n",
				"\n",
				usage(),
			))
	}

	return err
}
