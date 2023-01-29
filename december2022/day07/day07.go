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

func builtTaskList(cdSplittedLineList []cdSplittedLine) []task {
	taskList := []task{}

	path := []fileName{}
	for _, cdSplittedLine := range cdSplittedLineList {
		dirname, cdFreeLine := splittedNameAndRest(cdSplittedLine)
		if dirname != ".." {
			path = append(path, dirname)
			// manual deepcopy since path is a slice, hence shallow copies otherwise
			newPath := make([]fileName, len(path))
			copy(newPath, path)
			taskList = append(taskList, task{path: newPath, cdFreeLine: cdFreeLine})
		} else {
			path = path[:len(path)-1]
		}
	}

	return taskList
}

func builtFileTree(taskList []task) FileNode {
	waitGroup := new(sync.WaitGroup)
	fileChannel := make(chan []AbsFile)
	for _, t := range taskList {
		waitGroup.Add(1)
		go convertTaskToAbsFileList(t, fileChannel)
	}

	fileNode := new(FileNode).Init()
	go func() {
		for {
			receiveFileList(fileNode, fileChannel)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
	return *fileNode
}

func convertTaskToAbsFileList(t task, fileChannel chan<- []AbsFile) {
	fileList := []AbsFile{}

	lineList := strings.Split(t.cdFreeLine, "\n")
	for _, line := range lineList {
		if len(line) == 0 {
			continue
		}
		if len(line) >= 4 && line[0:4] == "$ ls" {
			continue
		}

		var name fileName
		var size *fileSize
		if len(line) >= 4 && line[0:4] == "dir " {
			name = line[4:]
			size = nil
		} else {
			tmp := strings.Split(line, " ")

			name = tmp[1]

			val, _ := strconv.ParseUint(tmp[0], 10, 0)
			*size = val
		}

		path := make([]fileName, len(t.path))
		copy(path, t.path)
		path = append(path, name)
		fileList = append(fileList, AbsFile{
			path: path,
			size: size,
		})
	}

	fileChannel <- fileList
}

func receiveFileList(fileNode *FileNode, fileChannel <-chan []AbsFile) {
	in := <-fileChannel
	for _, f := range in {
		if fileNode.name == nil {
			fileNode.name = &f.path[0]
		}
		fileNode.Add(f)
	}
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
	fileNode := builtFileTree(taskList)
	// wrong: 49192532
	fmt.Println("Directory size:", fileNode.size)

	return err
}
