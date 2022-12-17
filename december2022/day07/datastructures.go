package day07

type cdSplittedLine = string
type fileName = string
type cdFreeLine = string
type fileSize = uint64

type task struct {
	path       []fileName
	cdFreeLine cdFreeLine
}

type AbsFile struct {
	name *fileName
	path []fileName
	size fileSize
}

func (absFile AbsFile) IsDirectory() bool {
	return absFile.name == nil
}
