package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// mainfileDirpath, err := os.Executable()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// mainfileDirpath = path.Dir(mainfileDirpath)

	// december2022.Day01(mainfileDirpath)

	chosenFunc := "help"
	if len(os.Args) >= 1 {
		chosenFunc = os.Args[0]
	}

	supportedYear
	var year *flag.FlagSet
	switch chosenFunc {
	case "2022":
		year = flag.NewFlagSet("2022", flag.ExitOnError)
	default:
		// usage := "
		// "
		fmt.Println("Usage: adventofcode <year> [<args>]")
		fmt.Println("Supported years are: ")
		fmt.Println(" ask   Ask questions")
		fmt.Println(" send  Send messages to your contacts")
		return
	}

	day := year2022.Uint("day", 1, "Day")

	year2022.Parse(os.Args[1:])

	fmt.Println(*day)

	// content, err := os.ReadFile(inputFilepath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(content))
}
