package day07

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func splittedByCDCommands(logString string) []cdSplittedLine {
	logStringPerDirectory := strings.Split(logString, "$ cd ")
	if logStringPerDirectory[0] != "" {
		log.Fatalln("Expected \"$ cd DIR\" in first line")
	}
	return logStringPerDirectory[1:]
}

func splittedNameAndRest(cdSplittedLine cdSplittedLine) (fileName, cdFreeLine) {
	tmp := strings.SplitN(cdSplittedLine, "\n", 2)
	return tmp[0], tmp[1]
}

func builtTaskList(cdSplittedLineList []cdSplittedLine) *[]*task {
	taskList := []*task{}

	path := []fileName{}
	for _, cdSplittedLine := range cdSplittedLineList[0:] {
		dirname, cdFreeLine := splittedNameAndRest(cdSplittedLine)
		if dirname != ".." {
			path = append(path, dirname)
			// manual deepcopy since path is a slice, hence shallow copies otherwise
			newPath := make([]fileName, len(path))
			copy(newPath, path)
			taskList = append(taskList, &task{path: newPath, cdFreeLine: cdFreeLine})
		} else {
			path = path[:len(path)-1]
		}
	}

	return &taskList
}

func builtFileTree(taskList *[]*task) ([]*AbsFile, fileSize) {
	waitGroup := new(sync.WaitGroup)
	fileChannel := make(chan *[]*AbsFile)
	for _, t := range *taskList {
		waitGroup.Add(1)
		task := *t
		go convertTaskToAbsFileList(task, fileChannel)
	}

	fileList := []*AbsFile{}
	var totalSize fileSize = 0
	go func() {
		for {
			for _, f := range *<-fileChannel {
				fileList = append(fileList, f)
				totalSize += f.size
			}

			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
	// wrong: 49192532
	return fileList, totalSize
}

func convertTaskToAbsFileList(t task, fileChannel chan<- *[]*AbsFile) {
	fileList := []*AbsFile{}

	lineList := strings.Split(t.cdFreeLine, "\n")
	for _, line := range lineList {
		if len(line) == 0 {
			continue
		}
		if len(line) >= 4 && line[0:4] == "$ ls" {
			continue
		}

		name := new(fileName)
		var path []fileName
		var size fileSize
		if len(line) >= 4 && line[0:4] == "dir " {
			// is directory

			name = nil

			path = make([]fileName, len(t.path))
			copy(path, t.path)
			path = append(path, line[4:])

			size = 0
		} else {
			// is file

			tmp := strings.Split(line, " ")

			*name = tmp[1]

			size, _ = strconv.ParseUint(tmp[0], 10, 0)
		}

		fileList = append(fileList, &AbsFile{
			name: name,
			path: path,
			size: size,
		})
	}

	fileChannel <- &fileList
}

func Run(args []string) error {
	var err error

	config, err := parseCmdline(args)
	if err != nil {
		return err
	}

	logByteList, err := os.ReadFile(config.logFilepath)
	if err != nil {
		return err
	}

	taskList := builtTaskList(splittedByCDCommands(string(logByteList)))
	_, size := builtFileTree(taskList)
	fmt.Println("Directory size:", size)

	return err
}
