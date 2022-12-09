package day06

import (
	"fmt"
	"os"
)

func Run(args []string) error {
	config := parseCmdline(args)
	var err error
	streamByteList, err := os.ReadFile(config.streamFilepath)
	if err != nil {
		return err
	}
	streamData := string(streamByteList)

	runeList := []rune(streamData)
	patternLen := config.patternLen

	i := 0
	for i < len(runeList)-(patternLen-1) {
		oldI := i
		for offset := (patternLen - 1); offset > 0; offset-- {
			// assuming p==r and offset==3 (patternLen == 4)
			// ..p|r -> skip 3
			// .p.|r -> skip 2
			// p..|r -> skip 1
			r := runeList[i+offset]
			for j := i + offset - 1; j >= i; j-- {
				p := runeList[j]
				if p == r {
					i = j + 1
					// break
					j = -1
					offset = -1
				}
			}
		}
		if i == oldI {
			break
		}
	}

	fmt.Println("Processed characters until first marker is processed:", i+patternLen)

	return err
}
