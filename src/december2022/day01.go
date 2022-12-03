package december2022

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func SetupCmdlineParser() {
}

func Day01(mainfileDirpath string) {
	var inputFilepath string

	help_msg := "Input"
	flag.StringVar(&inputFilepath, "input", mainfileDirpath+"/../res/input/20221201.0.txt", help_msg)

	flag.Parse()

	content, err := os.ReadFile(inputFilepath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}
